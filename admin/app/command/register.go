package command

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	gormx "rin-echo/common/gorm"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type Register struct {
	Username string `json:"username" validate:"required,min=5"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r Register) ToCreateUser() CreateUser {
	return CreateUser{
		Username: r.Username,
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
	}
}

type RegisterHandler struct {
	uow  iuow.UnitOfWork
	rbac *adapters.RBACCasbin
}

func NewRegisterHandler(uow iuow.UnitOfWork, rbac *adapters.RBACCasbin) RegisterHandler {
	if uow == nil {
		panic("NewRegisterHandler requires uow")
	}

	return RegisterHandler{uow, rbac}
}

func (h RegisterHandler) Handle(ctx echox.Context, cmd Register) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow      = h.uow.WithContext(ctx.RequestContext())
		repoRole = repository.NewRoleRepository(uow.DB())
		roleIDs  []uint
	)

	if err := gormx.FindWrapError(repoRole.Query(map[string][]interface{}{"is_default": {true}}, nil), &roleIDs); err != nil {
		return err
	}

	cmdUser := cmd.ToCreateUser()
	cmdUser.RoleIDs = roleIDs
	if err := cmdUser.Check(uow); err != nil {
		return err
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		return cmdUser.Create(ux, h.rbac)
	}); err != nil {
		return err
	}

	return nil
}
