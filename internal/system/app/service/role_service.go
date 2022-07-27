package service

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/response"
	querybuilder "github.com/rinnguyen1614/rin-echo/internal/system/domain/query_builder"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/core/query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/system/errors"

	"github.com/jinzhu/copier"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"go.uber.org/zap"
)

type (
	RoleService interface {
		WithContext(echox.Context) RoleService

		Create(request.CreateRole) (uint, error)

		Update(id uint, cmd request.UpdateRole) (err error)

		Delete(id uint) (err error)

		Query(q *query.Query) (*model.QueryResult, error)

		Get(id uint) (response.Role, error)

		// 	// Update()

		// 	// Disable()

		// 	// Enable()

		// 	// ChangePassword()

		// 	// AssignToRoles(id uint, roleIDs []uint)

		// 	// AssignToRole(id, roleID uint)

		// 	// FindByRolenameOrEmail(rolenameOrEmail string) error

	}

	roleService struct {
		*echox.Service

		permissionManager domain.PermissionManager
		repo              domain.RoleRepository
		repoMenu          domain.MenuRepository
		repoPermission    domain.PermissionRepository
	}
)

func NewRoleService(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, settingProvider setting.Provider, logger *zap.Logger) RoleService {
	return &roleService{
		Service: echox.NewService(uow, settingProvider, logger),

		permissionManager: permissionManager,
		repo:              repository.NewRoleRepository(uow.DB()),
		repoMenu:          repository.NewMenuRepository(uow.DB()),
		repoPermission:    repository.NewPermissionRepository(uow.DB()),
	}
}

func (s *roleService) WithContext(ctx echox.Context) RoleService {
	return &roleService{
		Service: s.Service.WithContext(ctx),

		permissionManager: s.permissionManager,
		repo:              s.repo.WithTransaction(s.Service.Uow.DB()),
		repoMenu:          s.repoMenu.WithTransaction(s.Service.Uow.DB()),
		repoPermission:    s.repoPermission.WithTransaction(s.Service.Uow.DB()),
	}
}

func (s roleService) Create(cmd request.CreateRole) (uint, error) {
	var (
		role *domain.Role
		err  error
	)
	if err = s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var repo = s.repo.WithTransaction(ux.DB())

		if err = s.CheckExistBySlug(cmd.Slug); err != nil {
			return err
		}

		role, err = domain.NewNotRoleStatic(cmd.Name, cmd.Slug, cmd.IsDefault)
		if err != nil {
			return err
		}

		if err = repo.Create(role); err != nil {
			return err
		}

		if len(cmd.ResourceIDs) != 0 {
			newPermissions, _ := domain.NewPermissionsForRole(role.ID, cmd.ResourceIDs)
			return s.assignPermissionsToRole(ux, role, newPermissions)
		}

		return s.SetMenus(role, cmd.MenuIDs)

	}); err != nil {
		return 0, err
	}

	return role.ID, nil
}

func (s roleService) Update(id uint, cmd request.UpdateRole) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo = s.repo.WithTransaction(ux.DB())
			role domain.Role
		)
		err := repo.GetID(&role, id, nil)
		if err != nil {
			return err
		}

		if cmd.Slug != role.Slug {
			if err = s.CheckExistBySlug(cmd.Slug); err != nil {
				return err
			}
		}
		cmd.Populate(&role)
		if err = repo.Update(&role); err != nil {
			return nil
		}

		if err = s.SetResources(&role, cmd.ResourceIDs); err != nil {
			return err
		}

		return s.SetMenus(&role, cmd.MenuIDs)
	})
}

func (s roleService) SetResources(role *domain.Role, resourceIDs []uint) (err error) {

	if role == nil {
		panic("requires role")
	}

	if len(resourceIDs) == 0 {
		return nil
	}

	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo           = s.repo.WithTransaction(ux.DB())
			newPermissions domain.Permissions
		)

		err := repo.Find(role, nil, map[string][]interface{}{"Permissions": nil})
		if err != nil {
			return err
		}

		newPermissions, _ = domain.NewPermissionsForRole(role.ID, resourceIDs)

		permissionNews, permissionGranteds, permissionUngranteds := role.ComparePermissions(newPermissions)

		// ungrant permission from removed resources
		if len(permissionUngranteds) != 0 {
			if err = s.ungrantPermissionsFromRole(ux, role, permissionUngranteds); err != nil {
				return err
			}
		}

		// if permission is existed, grant permission
		if len(permissionGranteds) != 0 {
			if err = s.grantPermissionsFromRole(ux, role, permissionGranteds); err != nil {
				return err
			}
		}

		// add new permission.
		if len(permissionNews) != 0 {
			if err = s.assignPermissionsToRole(ux, role, permissionNews); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s roleService) SetMenus(role *domain.Role, menuIDs []uint) error {
	if role == nil {
		panic("requires role")
	}

	if len(menuIDs) == 0 {
		return nil
	}

	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repoMenu = s.repoMenu.WithTransaction(ux.DB())
			repo     = s.repo.WithTransaction(ux.DB())
			newMenus domain.Menus
		)

		err := repo.Find(role, nil, map[string][]interface{}{"Menus": nil})
		if err != nil {
			return err
		}

		if err = repoMenu.FindID(&newMenus, menuIDs, nil); err != nil {
			return err
		}

		menusNews, menusDels := role.CompareMenus(newMenus)

		// remove from removed menus
		if len(menusDels) != 0 {
			if err = ux.Association(role, "Menus").Delete(menusDels); err != nil {
				return err
			}
		}

		// add to added menus
		if len(menusNews) != 0 {
			if err = ux.Association(role, "Menus").Append(menusNews); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s roleService) Delete(id uint) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo           = s.repo.WithTransaction(ux.DB())
			hasResource, _ = uow.Contains(ux.DB().Table("permissions").Where("role_id", id))
			hasMenu, _     = uow.Contains(ux.DB().Table("menu_roles").Where("role_id", id))
			hasUser, _     = uow.Contains(ux.DB().Table("user_roles").Where("role_id", id))
		)

		if hasResource {
			return errors.ErrRoleReferencedResource
		}

		if hasMenu {
			return errors.ErrRoleReferencedMenu
		}

		if hasUser {
			return errors.ErrRoleReferencedUser
		}

		if err := repo.Delete(id); err != nil {
			return err
		}

		return nil
	})
}

func (s roleService) CheckExistBySlug(slug string) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"slug": {slug}}); ok {
		return errors.ErrRoleSlugExists
	}

	return nil
}

func (s roleService) Query(q *query.Query) (*model.QueryResult, error) {
	var (
		queryBuilder    = querybuilder.NewRoleQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"UserRoles":   querybuilder.NewUserRoleQueryBuilder(),
			"Permissions": querybuilder.NewPermissionQueryBuilder(),
			"Menus":       querybuilder.NewMenuQueryBuilder(),
		}
	)

	return q.QueryResult(s.repo, queryBuilder, preloadBuilders, domain.Role{}, response.Role{})
}

func (s roleService) Get(id uint) (response.Role, error) {
	var (
		role domain.Role
		res  response.Role
	)
	if err := s.repo.GetID(&role, id, map[string][]interface{}{"Menus": nil, "Permissions": {"is_granted = ?", true}, "Permissions.Resource": nil}); err != nil {
		return response.Role{}, err
	}

	if err := copier.CopyWithOption(&res, role, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return response.Role{}, err
	}
	return res, nil
}

func (s roleService) ungrantPermissionsFromRole(ux iuow.UnitOfWork, role *domain.Role, permissionUngranteds domain.Permissions) (err error) {
	var (
		repoPermission = repository.NewPermissionRepository(ux.DB())
	)

	if err = repoPermission.UpdateValues(
		map[string][]interface{}{"id": {permissionUngranteds.IDs()}},
		map[string]interface{}{
			"is_granted": false,
		}); err != nil {
		return err
	}

	return s.removePermissionFromCasbin(ux, role, permissionUngranteds)
}

func (s roleService) grantPermissionsFromRole(ux iuow.UnitOfWork, role *domain.Role, permissionGranteds domain.Permissions) (err error) {
	var (
		repoPermission = repository.NewPermissionRepository(ux.DB())
	)

	if err = repoPermission.UpdateValues(
		map[string][]interface{}{"id": {permissionGranteds.IDs()}},
		map[string]interface{}{
			"is_granted": true,
		}); err != nil {
		return err
	}

	return s.addPermissionToCasbin(ux, role, permissionGranteds)
}

func (s roleService) assignPermissionsToRole(ux iuow.UnitOfWork, role *domain.Role, permissionNews domain.Permissions) (err error) {

	if err = ux.Association(role, "Permissions").Append(permissionNews); err != nil {
		return err
	}

	return s.addPermissionToCasbin(ux, role, permissionNews)
}

func (s roleService) addPermissionToCasbin(ux iuow.UnitOfWork, role *domain.Role, permissionNews domain.Permissions) (err error) {
	var (
		repoResource    = repository.NewResourceRepository(ux.DB())
		resources       domain.Resources
		resourcesForPer domain.Resources
	)

	if err = uow.Find(repoResource.
		Query(map[string][]interface{}{"id": {permissionNews.ResourceIDs()}}, nil).
		Select("resources.object, resources.action"), &resources); err != nil {
		return err
	}

	if len(resources) != 0 {
		for _, re := range resources {
			if !s.permissionManager.HasPermissionForRole(role.ID, *re) {
				resourcesForPer = append(resourcesForPer, re)
			}
		}

		if _, err = s.permissionManager.AddPermissionsForRole(role.ID, resourcesForPer); err != nil {
			return err
		}
	}

	return nil
}

func (s roleService) removePermissionFromCasbin(ux iuow.UnitOfWork, role *domain.Role, permissionDels domain.Permissions) (err error) {
	var (
		repoResource = repository.NewResourceRepository(ux.DB())
		resources    domain.Resources
	)
	if err = uow.Find(repoResource.
		Query(map[string][]interface{}{"id": {permissionDels.ResourceIDs()}}, nil).
		Select("resources.object, resources.action"), &resources); err != nil {
		return err
	}

	if len(resources) != 0 {
		if _, err = s.permissionManager.RemovePermissionsForRole(role.ID, resources); err != nil {
			return err
		}
	}
	return nil
}
