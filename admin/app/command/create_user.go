package command

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/domain"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type CreateUser struct {
	cqrs.CreateCommand

	Username string `json:"username" validate:"required,min=5"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleIDs  []uint `json:"role_ids"`
}

func (cmd CreateUser) Check(uow iuow.UnitOfWork) error {
	var (
		repoUser = repository.NewUserRepository(uow.DB())
	)

	if ok := repoUser.Contains(map[string][]interface{}{"user_name": {cmd.Username}}); ok {
		return errors.ERR_USERNAME_EXISTS
	}

	return nil
}

func (cmd CreateUser) Create(uow iuow.UnitOfWork, rbac *adapters.RBACCasbin) error {
	var (
		repoUser     = repository.NewUserRepository(uow.DB())
		repoRole     = repository.NewRoleRepository(uow.DB())
		repoUserRole = repository.NewUserRoleRepository(uow.DB())
		user         domain.User
	)

	user, err := domain.NewUser(cmd.Username, cmd.Password, cmd.FullName, cmd.Email)
	if err != nil {
		return err
	}

	if err = repoUser.Create(&user); err != nil {
		return err
	}

	if len(cmd.RoleIDs) != 0 {
		var (
			userRoles  []*domain.UserRole
			roleIDStrs []string
			roles      []*domain.Role
		)

		if err := repoRole.FindID(&roles, cmd.RoleIDs, nil); err != nil {
			return err
		}

		for _, r := range roles {
			uR, _ := domain.NewUserRole(user.ID, r.ID)
			userRoles = append(userRoles, &uR)
			roleIDStrs = append(roleIDStrs, utils.ToString(r.ID))
		}

		if err = repoUserRole.Create(&userRoles); err != nil {
			return err
		}

		if _, err = rbac.AddRolesForUser(utils.ToString(user.ID), roleIDStrs); err != nil {
			return err
		}
	}

	return nil
}

type CreateUserHandler struct {
	uow  iuow.UnitOfWork
	rbac *adapters.RBACCasbin
}

func NewCreateUserHandler(uow iuow.UnitOfWork, rbac *adapters.RBACCasbin) CreateUserHandler {
	if uow == nil {
		panic("NewCreateUserHandler requires uow")
	}

	return CreateUserHandler{uow, rbac}
}

func (h CreateUserHandler) Handle(ctx echox.Context, cmd CreateUser) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow = h.uow.WithContext(ctx.RequestContext())
	)

	if err := cmd.Check(uow); err != nil {
		return err
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		return cmd.Create(ux, h.rbac)
	}); err != nil {
		return err
	}

	return nil

}
