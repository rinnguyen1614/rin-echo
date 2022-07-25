package initdata

import (
	"strings"

	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"

	"github.com/rinnguyen1614/rin-echo/internal/core/uow"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

func initRoleAndUser(uow iuow.UnitOfWork, permissionManager domain.PermissionManager) error {
	var (
		repoRole = repository.NewRoleRepository(uow.DB())
	)

	// create admin and user role.
	rAdmin, err := repoRole.FirstBySlug(domain.RoleAdmin, nil)
	if rAdmin == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.RoleAdmin), domain.RoleAdmin, true, false)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
		rAdmin = rNew
		if err := assignPermissionsToRoleAdmin(uow, permissionManager, rAdmin); err != nil {
			return err
		}

		if err := assignMenusToRoleAdmin(uow, rAdmin); err != nil {
			return err
		}
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
		user, _ := domain.NewUser("admin", "Admin@0809", "Admin", "admin@rin-echo.com", []uint{rAdmin.ID})
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
		repo      = repository.NewResourceRepository(ux.DB())
		resources domain.Resources
	)
	if err := uow.Find(repo.Query(nil, nil), &resources); err != nil {
		return err
	}

	return assignPermissions(ux, permissionManager, rAdmin, resources)
}

func assignPermissions(ux iuow.UnitOfWork, permissionManager domain.PermissionManager, role *domain.Role, resources domain.Resources) (err error) {
	if len(resources) == 0 {
		return
	}

	var (
		newPermissions, _ = domain.NewPermissionsForRole(role.ID, resources.IDs())
		resourcesForPer   domain.Resources
	)

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

// assignMenus to admin role
func assignMenusToRoleAdmin(ux iuow.UnitOfWork, rAdmin *domain.Role) error {
	var (
		repo  = repository.NewMenuRepository(ux.DB())
		menus domain.Menus
	)
	if err := uow.Find(repo.Query(nil, nil), &menus); err != nil {
		return err
	}

	return assignMenus(ux, rAdmin, menus)
}

func assignMenus(ux iuow.UnitOfWork, role *domain.Role, menusNews domain.Menus) (err error) {
	if len(menusNews) == 0 {
		return
	}

	if err = ux.Association(role, "Menus").Append(menusNews); err != nil {
		return err
	}
	return nil
}
