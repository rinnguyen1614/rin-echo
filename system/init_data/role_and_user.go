package initdata

import (
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/adapters/repository"
	"rin-echo/system/domain"
	"strings"
)

func initRoleAndUser(uow iuow.UnitOfWork, permissionManager domain.PermissionManager) error {
	var (
		repoRole = repository.NewRoleRepository(uow.DB())
	)

	// create administrator and user role.
	rAdmin, err := repoRole.FirstBySlug(domain.RoleAdministrator, nil)
	if rAdmin == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.RoleAdministrator), domain.RoleAdministrator, true, false)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
		rAdmin = rNew
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
