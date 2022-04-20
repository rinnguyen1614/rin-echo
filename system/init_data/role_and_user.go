package initdata

import (
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/adapters/repository"
	"rin-echo/system/domain"
	"strings"
)

func initRoleAndUser(uow iuow.UnitOfWork, permissionManager domain.PermissionManager) error {
	var (
		repoRole = repository.NewRoleRepository(uow.DB())
	)

	// create admini and user role.
	rAdmin, err := repoRole.FirstBySlug(domain.RoleAdmin, nil)
	if rAdmin == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.RoleAdmin), domain.RoleAdmin, true, false)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
		rAdmin = rNew
		assignPermissionsToRoleAdmin(uow, permissionManager, rAdmin)
	}

	rUser, err := repoRole.FirstBySlug(domain.RoleUser, nil)
	if rUser == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.RoleUser), domain.RoleUser, false, true)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
	}

	return createUserAdmin(uow, permissionManager, *rAdmin)
}

// create user is an admin
func createUserAdmin(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, rAdmin domain.Role) error {
	var repoUser = repository.NewUserRepository(uow.DB())
	user, err := repoUser.FirstByUsernameOrEmail("admin", nil)
	if user == nil {
		user, _ := domain.NewUser("admin", "admin@0809", "Admin", "admin@rin-echo.com", []uint{rAdmin.ID})
		if err = repoUser.Create(&user); err != nil {
			return err
		}

		for _, uR := range user.UserRoles {
			if _, err := permissionManager.AddRoleForUser(user.ID, uR.RoleID); err != nil {
				return err
			}

		}
	}

	return nil
}

// assignPermissions to admin role
func assignPermissionsToRoleAdmin(ux iuow.UnitOfWork, permissionManager domain.PermissionManager, rAdmin *domain.Role) error {
	var (
		repo    = repository.NewMenuRepository(ux.DB())
		menuIDs []uint
	)
	if err := uow.Find(repo.Query(nil, nil).Select("id"), &menuIDs); err != nil {
		return err
	}
	return assignPermissions(ux, permissionManager, rAdmin, menuIDs)
}

func assignPermissions(ux iuow.UnitOfWork, permissionManager domain.PermissionManager, role *domain.Role, menuIDs []uint) (err error) {
	if len(menuIDs) == 0 {
		return
	}

	var (
		repoResource      = repository.NewResourceRepository(ux.DB())
		newPermissions, _ = domain.NewPermissionsForRole(role.ID, menuIDs)
		resources         domain.Resources
		resourcesForPer   domain.Resources
	)
	if err = uow.Find(repoResource.QueryByMenus(menuIDs, nil).Select("resources.path, resources.method"), &resources); err != nil {
		return err
	}

	if err = ux.Association(role, "Permissions").Append(newPermissions); err != nil {
		return err
	}

	if len(resources) != 0 {
		for _, re := range resources {
			if !permissionManager.HasPermissionForRole(role.ID, *re) {
				resourcesForPer = append(resourcesForPer, re)
			}
		}

		if _, err = permissionManager.AddPermissionsForRole(role.ID, resourcesForPer); err != nil {
			return err
		}
	}
	return nil
}
