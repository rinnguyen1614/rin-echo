package service

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/response"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain/query_builder"
	querybuilder "github.com/rinnguyen1614/rin-echo/internal/system/domain/query_builder"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/core/query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"github.com/rinnguyen1614/rin-echo/internal/system/errors"
)

type (
	ResourceService interface {
		WithContext(echox.Context) ResourceService

		Create(request.CreateResource) (uint, error)

		Update(id uint, cmd request.UpdateResource) (err error)

		Delete(id uint) (err error)

		Get(id uint) (response.Resource, error)

		Find(q *query.Query) (*model.QueryResult, error)

		FindTrees(q *query.Query) (*model.QueryResult, error)
	}

	resourceService struct {
		*echox.Service
		permissionManager domain.PermissionManager
		repo              domain.ResourceRepository
	}
)

func NewResourceService(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, settingProvider setting.Provider, logger *zap.Logger) ResourceService {
	return &resourceService{
		Service:           echox.NewService(uow, settingProvider, logger),
		permissionManager: permissionManager,
		repo:              repository.NewResourceRepository(uow.DB()),
	}
}

func (s *resourceService) WithContext(ctx echox.Context) ResourceService {
	return &resourceService{
		Service:           s.Service.WithContext(ctx),
		permissionManager: s.permissionManager,
		repo:              s.repo.WithTransaction(s.Service.Uow.DB()),
	}
}

func (s resourceService) Create(cmd request.CreateResource) (uint, error) {
	var id uint
	if err := s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		parentsByIndex, err := s.Check(request.CreateResources{&cmd}, true)
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

func (s resourceService) createRecursive(cmd request.CreateResource, parent *domain.Resource) (uint, error) {
	resource, err := domain.NewResource(cmd.Name, cmd.Slug, cmd.Object, cmd.Action, cmd.Description, parent)
	if err != nil {
		return 0, err
	}

	if err = s.repo.Create(resource); err != nil {
		return 0, err
	}

	for _, mc := range cmd.Children {
		if _, err := s.createRecursive(*mc, resource); err != nil {
			return 0, err
		}
	}

	return resource.ID, nil
}

func (s resourceService) Update(id uint, cmd request.UpdateResource) (err error) {
	if err = s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {

		var (
			repo             = s.repo.WithTransaction(ux.DB())
			updatePermission bool
			resource         domain.Resource
			beforeUpdate     domain.Resource
		)

		err := repo.GetID(&resource, id, nil)
		if err != nil {
			return err
		}

		beforeUpdate = resource
		if cmd.Slug != resource.Slug {
			if err = s.CheckExistBySlug(cmd.Slug); err != nil {
				return err
			}
		}
		if !cmd.IsEqualObjectAndAction(resource.Object, resource.Action) && !cmd.IsEmptyObjectAndAction() {
			if err = s.CheckExistObjectAndAction(cmd.Object, cmd.Action); err != nil {
				return err
			}
		}

		var parentID = utils.DefaultValue(cmd.ParentID, 0).(uint)
		if utils.DefaultValue(beforeUpdate.ParentID, 0).(uint) != parentID {
			var parent *domain.Resource
			if parentID != 0 {
				if err := s.repo.GetID(&parent, parentID, nil); err != nil {
					return err
				}
			}

			resource.SetParent(parent)
		}

		updatePermission = cmd.Object != resource.Object || cmd.Action != resource.Action

		cmd.Populate(&resource)
		if err := repo.Update(&resource); err != nil {
			return err
		}

		if updatePermission {
			return s.UpdatePermission(beforeUpdate, resource)
		}

		return nil

	}); err != nil {
		return err
	}

	return nil
}

func (s resourceService) UpdatePermission(oldResource domain.Resource, newResource domain.Resource) error {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repoPermission = repository.NewPermissionRepository(ux.DB())
			roleIDs        []uint
		)

		if err := uow.Find(repoPermission.Query(map[string][]interface{}{"resource_id": {oldResource.ID}}, nil).Select("permissions.role_id"), &roleIDs); err != nil {
			return err
		}

		if len(roleIDs) > 0 {
			if _, err := s.permissionManager.UpdatePermissionForRoles(roleIDs, oldResource, newResource); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s resourceService) Delete(id uint) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo       = s.repo.WithTransaction(ux.DB())
			hasRole, _ = uow.Contains(ux.DB().Table("permissions").Where("resource_id", id))
		)

		if hasRole {
			return errors.ErrResourceReferencedRole
		}

		if err := repo.Delete(id); err != nil {
			return err
		}

		return nil
	})
}

func (s resourceService) CheckExistBySlug(slug string) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"slug": {slug}}); ok {
		return errors.ErrResourceSlugExists
	}

	return nil
}

func (s resourceService) CheckExistParent(parentID uint) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"id": {parentID}}); !ok {
		return errors.ErrResourceParentNotFound
	}

	return nil
}

func (s resourceService) CheckExistObjectAndAction(object, action string) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"object": {object}, "action": {action}}); ok {
		return errors.ErrResourceObjectAndActionExists
	}

	return nil
}

func (s resourceService) Check(cmds request.CreateResources, checkParent bool) (parentsByIndex map[int]*domain.Resource, err error) {
	var (
		errorsByIndex           = make(map[int][]error)
		childrenByIndex         = make(map[int]request.CreateResources)
		indexsBySlug            = make(map[string]int)
		indexsByObjectAndAction = make(map[string]int)
		indexsByParentID        = make(map[uint][]int)
		cmdSlugs                []string
		cmdPaths                []string
		cmdParentIDs            []uint
	)

	parentsByIndex = make(map[int]*domain.Resource)

	for i, cmd := range cmds {
		//slug
		cmdSlugs = append(cmdSlugs, cmd.Slug)
		if _, ok := indexsBySlug[cmd.Slug]; ok {
			errorsByIndex[i] = append(errorsByIndex[i], errors.ErrResourceSlugExists)
		} else {
			indexsBySlug[cmd.Slug] = i
		}

		// object_action
		cmdPaths = append(cmdPaths, cmd.Object)
		if !cmd.IsEmptyObjectAndAction() {
			key := joinObjectAndAction(cmd.Object, cmd.Action)
			if _, ok := indexsByObjectAndAction[key]; ok {
				errorsByIndex[i] = append(errorsByIndex[i], errors.ErrResourceObjectAndActionExists)
			} else {
				indexsByObjectAndAction[key] = i
			}
		}

		// parent
		if parentID := cmd.ParentID; parentID > 0 {
			cmdParentIDs = append(cmdParentIDs, parentID)
			idxs := indexsByParentID[cmd.ParentID]
			idxs = append(idxs, i)
			indexsByParentID[cmd.ParentID] = idxs
		}

		if len(cmd.Children) != 0 {
			childrenByIndex[i] = cmd.Children
		}
	}

	if len(cmdSlugs) != 0 {
		if err := s.CheckExistBySlugs(cmdSlugs, indexsBySlug, errorsByIndex); err != nil {
			return nil, err
		}
	}

	if len(cmdPaths) != 0 {
		if err := s.CheckExistByObjectAndActions(cmdPaths, indexsByObjectAndAction, errorsByIndex); err != nil {
			return nil, err
		}
	}

	if len(cmdParentIDs) != 0 && checkParent {
		if parentsByIndex, err = s.CheckExistParents(cmdParentIDs, indexsByParentID, errorsByIndex); err != nil {
			return nil, err
		}
	}

	for i, children := range childrenByIndex {
		if _, err = s.Check(children, false); err != nil {
			if rErr, ok := err.(*core.RinErrors); ok {
				err := core.NewRinErrors(rErr.Errors(), "children_error", "You have some errors for create resources")
				errorsByIndex[i] = append(errorsByIndex[i], err)
			} else {
				return nil, err
			}
		}
	}

	if len(errorsByIndex) != 0 {
		return nil, core.NewRinErrors(errorsByIndex, "create_resource_error", "You have some errors for create resources")
	}

	return parentsByIndex, nil
}

func (s resourceService) CheckExistBySlugs(slugs []string, indexsBySlug map[string]int, errorsByIndex map[int][]error) error {
	var (
		qb         = query_builder.NewResourceQueryBuilder()
		slugsFound []string
	)
	qb.SetCondition(gormx.InOperator.Condition("slug"), slugs)
	qb.SetSelect("slug")

	if err := s.repo.QueryBuilderFind(&slugsFound, qb); err != nil {
		return err
	}

	for _, slug := range slugsFound {
		if i, ok := indexsBySlug[slug]; ok {
			errorsByIndex[i] = append(errorsByIndex[i], errors.ErrResourceSlugExists)
		}
	}
	return nil
}

func (s resourceService) CheckExistByObjectAndActions(objects []string, indexsByObjectAndAction map[string]int, errorsByIndex map[int][]error) error {
	var (
		qb           = query_builder.NewResourceQueryBuilder()
		objectsFound []map[string]interface{}
	)
	qb.SetCondition(gormx.InOperator.Condition("object"), objects)
	qb.SetSelect("object", "action")

	if err := s.repo.QueryBuilderFind(&objectsFound, qb); err != nil {
		return err
	}

	for _, pF := range objectsFound {
		object, action := pF["object"].(string), pF["action"].(string)
		key := joinObjectAndAction(object, action)
		if i, ok := indexsByObjectAndAction[key]; ok {
			errorsByIndex[i] = append(errorsByIndex[i], errors.ErrResourceObjectAndActionExists)
		}
	}
	return nil
}

func (s resourceService) CheckExistParents(parentIDs []uint, indexsByParentID map[uint][]int, errorsByIndex map[int][]error) (map[int]*domain.Resource, error) {
	var (
		parents        domain.Resources
		parentsByIndex = make(map[int]*domain.Resource)
	)
	if err := s.repo.FindID(&parents, parentIDs, nil); err != nil {
		return nil, err
	}

	mapParents := parents.ToMap()
	for parentID, idxs := range indexsByParentID {
		if p, ok := mapParents[parentID]; !ok {
			for i := range idxs {
				errorsByIndex[i] = append(errorsByIndex[i], errors.ErrResourceParentNotFound)
			}
		} else {
			for i := range idxs {
				parentsByIndex[i] = p
			}
		}
	}

	return parentsByIndex, nil
}

func (s resourceService) Get(id uint) (response.Resource, error) {
	var resource domain.Resource
	if err := s.repo.GetID(&resource, id, nil); err != nil {
		return response.Resource{}, err
	}
	return response.NewResource(resource), nil
}

func (s resourceService) Find(q *query.Query) (*model.QueryResult, error) {
	var (
		queryBuilder    = querybuilder.NewResourceQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"Menus": querybuilder.NewMenuQueryBuilder(),
		}
	)

	return q.QueryResult(s.repo, queryBuilder, preloadBuilders, domain.Resource{}, response.Resource{})
}

func (s resourceService) FindTree(id uint) (resource *domain.Resource, err error) {
	if err = s.repo.GetID(&resource, id, nil); err != nil {
		return nil, err
	}
	return resource, nil
}

func (s resourceService) FindTrees(q *query.Query) (*model.QueryResult, error) {
	var (
		repo            = s.repo
		queryBuilder    = querybuilder.NewResourceQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"Permissions": querybuilder.NewPermissionQueryBuilder(),
			"Roles":       querybuilder.NewRoleQueryBuilder(),
		}
		srcModels = new(domain.Resources)
		desModel  = response.ResourceTree{}
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

func joinObjectAndAction(object, action string) string {
	return object + "_join_" + action
}
