package service

import (
	echox "rin-echo/common/echo"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/adapters/repository"
	"rin-echo/system/app/model/request"
	"rin-echo/system/domain"
	"rin-echo/system/errors"

	"go.uber.org/zap"
)

type (
	RoleService interface {
		WithContext(echox.Context) RoleService

		Create(request.CreateRole) (uint, error)

		Update(id uint, cmd request.UpdateRole) (err error)

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

func NewRoleService(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, logger *zap.Logger) RoleService {
	return &roleService{
		Service: echox.NewService(uow, logger),

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

		role, err = domain.NewRole(cmd.Name, cmd.Slug, cmd.IsStatic, cmd.IsDefault)
		if err != nil {
			return err
		}

		if err = repo.Create(role); err != nil {
			return err
		}

		if len(cmd.MenuIDs) != 0 {
			newPermissions, _ := domain.NewPermissionsForRole(role.ID, cmd.MenuIDs)
			return s.assignPermissionsToRole(ux, role, newPermissions)
		}

		return nil

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

		if err = repo.Update(&role); err != nil {
			return nil
		}

		return s.SetMenus(&role, cmd.MenuIDs)
	})
}

func (s roleService) SetMenus(role *domain.Role, menuIDs []uint) (err error) {

	if role == nil {
		panic("requires role")
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

		newPermissions, _ = domain.NewPermissionsForRole(role.ID, menuIDs)

		permissionNews, permissionDels := role.ComparePermissions(newPermissions)

		// remove from removed menus
		if len(permissionDels) != 0 {
			if err = s.removePermissionsFromRole(ux, role, permissionDels); err != nil {
				return nil
			}
		}

		// add to added menus
		if len(permissionNews) != 0 {
			if err = s.assignPermissionsToRole(ux, role, permissionNews); err != nil {
				return nil
			}
		}
		return nil
	})
}

func (s roleService) removePermissionsFromRole(ux iuow.UnitOfWork, role *domain.Role, permissionDels domain.Permissions) (err error) {
	var (
		repoResource = repository.NewResourceRepository(ux.DB())
		resources    domain.Resources
	)
	if err = uow.Find(repoResource.QueryByMenus(permissionDels.MenuIDs(), nil).Select("resources.path, resources.method"), &resources); err != nil {
		return err
	}

	if err = ux.Association(role, "Permissions").Delete(permissionDels); err != nil {
		return err
	}

	if len(resources) != 0 {
		if _, err = s.permissionManager.RemovePermissionsForRole(role.ID, resources); err != nil {
			return err
		}
	}
	return nil
}

func (s roleService) assignPermissionsToRole(ux iuow.UnitOfWork, role *domain.Role, permissionNews domain.Permissions) (err error) {
	var (
		repoResource = repository.NewResourceRepository(ux.DB())
		resources    domain.Resources
	)

	if err = uow.Find(repoResource.QueryByMenus(permissionNews.MenuIDs(), nil).Select("resources.path, resources.method"), &resources); err != nil {
		return err
	}

	if err = ux.Association(role, "Permissions").Append(permissionNews); err != nil {
		return err
	}
	if len(resources) != 0 {
		if _, err = s.permissionManager.AddPermissionsForRole(role.ID, resources); err != nil {
			return err
		}
	}
	return nil
}

func (s roleService) CheckExistBySlug(slug string) error {
	if ok := s.repo.Contains(map[string][]interface{}{"slug": {slug}}); ok {
		return errors.ErrRoleSlugExists
	}

	return nil
}
