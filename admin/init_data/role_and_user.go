package initdata

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/app/command"
	"rin-echo/admin/domain"
	"rin-echo/common/cqrs"
	iuow "rin-echo/common/uow/interfaces"
	"strings"
)

func initRoleAndUser(uow iuow.UnitOfWork, rbac *adapters.RBACCasbin) error {
	var (
		repoRole = repository.NewRoleRepository(uow.DB())
		cmdUser  = command.CreateUser{
			CreateCommand: cqrs.CreateCommand{ID: 1},
			Username:      "administrator",
			FullName:      "Administrator",
			Email:         "administrator@rin-echo.com",
			Password:      "administrator@0809",
			RoleIDs:       []uint{},
		}
	)

	r, err := repoRole.FirstBySlug(domain.Role_Administrator, nil)
	if r == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.Role_Administrator), domain.Role_Administrator, true, false)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
		r = &rNew
	}

	cmdUser.RoleIDs = []uint{r.ID}
	if err = cmdUser.Create(uow, rbac); err != nil {
		return err
	}

	r, err = repoRole.FirstBySlug(domain.Role_User, nil)
	if r == nil {
		rNew, _ := domain.NewRole(strings.ToTitle(domain.Role_User), domain.Role_User, false, true)
		if err = repoRole.Create(&rNew); err != nil {
			return err
		}
	}

	return nil
}
