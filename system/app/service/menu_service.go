package service

import (
	"rin-echo/system/adapters/repository"
	"rin-echo/system/app/model/request"
	"rin-echo/system/app/model/response"
	"rin-echo/system/domain"
	"rin-echo/system/domain/query_builder"
	querybuilder "rin-echo/system/domain/query_builder"
	"rin-echo/system/errors"

	core "github.com/rinnguyen1614/rin-echo-core"

	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	gormx "github.com/rinnguyen1614/rin-echo-core/gorm"
	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/query"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	"github.com/rinnguyen1614/rin-echo-core/uow"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/utils"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type (
	MenuService interface {
		WithContext(echox.Context) MenuService

		Create(request.CreateMenu) (uint, error)

		Update(id uint, cmd request.UpdateMenu) (err error)

		Delete(id uint) (err error)

		Get(id uint) (response.Menu, error)

		// Find(q *query.Query) (*model.QueryResult, error)

		FindTrees(q *query.Query) (*model.QueryResult, error)
	}

	menuService struct {
		*echox.Service
		permissionManager domain.PermissionManager
		repo              domain.MenuRepository
		repoResource      domain.ResourceRepository
	}
)

func NewMenuService(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, settingProvider setting.Provider, logger *zap.Logger) MenuService {
	return &menuService{
		Service:           echox.NewService(uow, settingProvider, logger),
		permissionManager: permissionManager,
		repo:              repository.NewMenuRepository(uow.DB()),
		repoResource:      repository.NewResourceRepository(uow.DB()),
	}
}

func (s *menuService) WithContext(ctx echox.Context) MenuService {
	return &menuService{
		Service:           s.Service.WithContext(ctx),
		permissionManager: s.permissionManager,
		repo:              s.repo.WithTransaction(s.Service.Uow.DB()),
		repoResource:      s.repoResource.WithTransaction(s.Service.Uow.DB()),
	}
}

func (s menuService) Create(cmd request.CreateMenu) (uint, error) {
	var id uint
	if err := s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		parentsByIndex, err := s.CheckForCreate(request.CreateMenus{&cmd}, true)
		if err != nil {
			return err
		}

		rid, err := s.createRecursive(cmd, parentsByIndex[0])
		if err != nil {
			return err
		}

		id = rid
		return nil
	}); err != nil {
		return id, err
	}
	return id, nil
}

func (s menuService) createRecursive(cmd request.CreateMenu, parent *domain.Menu) (uint, error) {
	menu, err := domain.NewMenu(cmd.Name, cmd.Slug, cmd.Path, cmd.Hidden, cmd.Component, cmd.Sort, cmd.Type, cmd.Icon, cmd.Title, parent)
	if err != nil {
		return 0, err
	}

	if err = s.repo.Create(menu); err != nil {
		return 0, err
	}

	for _, mc := range cmd.Children {
		if _, err := s.createRecursive(*mc, menu); err != nil {
			return 0, err
		}
	}

	return menu.ID, nil
}

func (s menuService) Update(id uint, cmd request.UpdateMenu) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {

		var (
			repo         = s.repo.WithTransaction(ux.DB())
			menu         domain.Menu
			beforeUpdate domain.Menu
		)

		err := repo.GetID(&menu, id, nil)
		if err != nil {
			return err
		}

		beforeUpdate = menu

		var parentID = utils.DefaultValue(cmd.ParentID, 0).(uint)
		if utils.DefaultValue(beforeUpdate.ParentID, 0).(uint) != parentID {
			var parent *domain.Menu
			if parentID != 0 {
				if err := s.repo.GetID(&parent, parentID, nil); err != nil {
					return err
				}
			}

			menu.SetParent(parent)
		}

		cmd.Populate(&menu)
		if err = repo.Update(&menu); err != nil {
			return err
		}

		return nil
	})
}

func (s menuService) Delete(id uint) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo       = s.repo.WithTransaction(ux.DB())
			hasRole, _ = uow.Contains(ux.DB().Table("menu_roles").Where("menu_id", id))
		)

		if hasRole {
			return errors.ErrMenuReferencedRole
		}

		if err := repo.Delete(id); err != nil {
			return err
		}

		return nil
	})
}

func (s menuService) CheckForCreate(cmds request.CreateMenus, checkParent bool) (parentsByIndex map[int]*domain.Menu, err error) {
	var (
		errorsByIndex    = make(map[int][]error)
		childrenByIndex  = make(map[int]request.CreateMenus)
		indexsBySlug     = make(map[string]int)
		indexsByParentID = make(map[uint][]int)
		cmdSlugs         []string
		cmdParentIDs     []uint
	)

	parentsByIndex = make(map[int]*domain.Menu)

	for i, m := range cmds {
		cmdSlugs = append(cmdSlugs, m.Slug)

		if _, ok := indexsBySlug[m.Slug]; ok {
			errorsByIndex[i] = append(errorsByIndex[i], errors.ErrMenuSlugExists)
		} else {
			indexsBySlug[m.Slug] = i
		}

		if parentID := m.ParentID; parentID > 0 {
			cmdParentIDs = append(cmdParentIDs, parentID)
			idxs := indexsByParentID[m.ParentID]
			idxs = append(idxs, i)
			indexsByParentID[m.ParentID] = idxs
		}

		if len(m.Children) != 0 {
			childrenByIndex[i] = m.Children
		}
	}

	if len(cmdSlugs) != 0 {
		if err = s.CheckExistBySlugs(cmdSlugs, indexsBySlug, errorsByIndex); err != nil {
			return nil, err
		}
	}

	if len(cmdParentIDs) != 0 && checkParent {
		if parentsByIndex, err = s.CheckExistParents(cmdParentIDs, indexsByParentID, errorsByIndex); err != nil {
			return nil, err
		}
	}

	for i, children := range childrenByIndex {
		if _, err = s.CheckForCreate(children, false); err != nil {
			if rErr, ok := err.(*core.RinErrors); ok {
				err := core.NewRinErrors(rErr.Errors(), "children_error", "You have some errors for create menu")
				errorsByIndex[i] = append(errorsByIndex[i], err)
			} else {
				return nil, err
			}
		}
	}

	if len(errorsByIndex) != 0 {
		return nil, core.NewRinErrors(errorsByIndex, "create_menu_error", "You have some errors for create menu")
	}

	return parentsByIndex, nil
}

// func (s menuService) CheckForUpdate(menu domain.Menu, cmd request.UpdateMenu) (parentsByIndex map[int]*domain.Menu, err error) {
// 	var (
// 		errorsByIndex      = make(map[int][]error)
// 		indexsBySlug       = map[string]int{cmd.Slug: 0}
// 		indexsByParentID   = map[uint][]int{cmd.ParentID: []int{0}}
// 		indexsByResourceID = make(map[uint][]int)
// 		cmdSlugs           = []string{cmd.Slug}
// 		cmdParentIDs       = []uint{cmd.ParentID}
// 		cmdResourceIDs     = cmd.ResourceIDs
// 	)

// 	if menu.Slug != cmd.Slug {
// 		if err = s.CheckExistBySlugs(cmdSlugs, indexsBySlug, errorsByIndex); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if cmd.ParentID != *menu.ParentID && cmd.ParentID != 0 {
// 		if parentsByIndex, err = s.CheckExistParents(cmdParentIDs, indexsByParentID, errorsByIndex); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if len(cmdResourceIDs) != 0 {
// 		for _, reID := range cmdResourceIDs {
// 			idxs := indexsByResourceID[reID]
// 			idxs = append(idxs, i)
// 			indexsByResourceID[reID] = idxs
// 		}
// 		if err = s.CheckExistResources(cmdResourceIDs, indexsByResourceID, errorsByIndex); err != nil {
// 			return nil, err
// 		}
// 	}
// }

func (s menuService) CheckExistBySlugs(slugs []string, indexsBySlug map[string]int, errorsByIndex map[int][]error) error {
	var (
		qb         = query_builder.NewMenuQueryBuilder()
		slugsFound []string
	)
	qb.SetCondition(gormx.InOperator.Condition("slug"), slugs)
	qb.SetSelect("slug")

	if err := s.repo.QueryBuilderFind(&slugsFound, qb); err != nil {
		return err
	}

	for _, slug := range slugsFound {
		if i, ok := indexsBySlug[slug]; ok {
			errorsByIndex[i] = append(errorsByIndex[i], errors.ErrMenuSlugExists)
		}
	}
	return nil
}

func (s menuService) CheckExistParents(parentIDs []uint, indexsByParentID map[uint][]int, errorsByIndex map[int][]error) (map[int]*domain.Menu, error) {
	var (
		parents        domain.Menus
		parentsByIndex = make(map[int]*domain.Menu)
	)
	if err := s.repo.FindID(&parents, parentIDs, nil); err != nil {
		return nil, err
	}

	mapParents := parents.ToMap()
	for parentID, idxs := range indexsByParentID {
		if p, ok := mapParents[parentID]; !ok {
			for i := range idxs {
				errorsByIndex[i] = append(errorsByIndex[i], errors.ErrMenuParentNotFound)
			}
		} else {
			for i := range idxs {
				parentsByIndex[i] = p
			}
		}
	}

	return parentsByIndex, nil
}

func (s menuService) Get(id uint) (response.Menu, error) {
	var menu domain.Menu
	if err := s.repo.GetID(&menu, id, nil); err != nil {
		return response.Menu{}, err
	}
	return response.NewMenu(menu), nil
}

func (s menuService) FindTrees(q *query.Query) (*model.QueryResult, error) {
	var (
		repo            = s.repo
		queryBuilder    = querybuilder.NewMenuQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"Resources": querybuilder.NewResourceQueryBuilder(),
		}
		srcModels = new(domain.Menus)
		desModel  = response.MenuTree{}
		nChild    = 4
	)

	fields, _, err := utils.GetFullFieldsByJsonTag(desModel)
	// remove children because "resource" table don't contains "children" field.
	delete(fields, "children")

	totalRecords, err := q.BindQueryBuilder(queryBuilder, preloadBuilders, repo.DB(), fields)
	if err != nil {
		return nil, err
	}

	if err = repo.QueryBuilderFind(srcModels, queryBuilder); err != nil {
		return nil, err
	}

	selects := q.FlatSelect()
	// nests n child, child has fields that includes by all fields of root
	for i, preChil, fields := 0, "", selects; i < nChild; i++ {
		preChil += "children."
		for _, f := range fields {
			selects = append(selects, preChil+f)
		}
	}

	// new slice of desModel with fields that get from query' selects
	prune, err := utils.NewSliceOfStructsByTag(desModel, selects, "json")
	if err != nil {
		return nil, err
	}

	err = copier.CopyWithOption(prune, srcModels.ToTree(), copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}

	return model.NewQueryResult(prune, totalRecords, q.Paging().Limit, q.Paging().Offset), nil
}
