package command

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/domain"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type CreateUser struct {
	Username string `json:"username" validate:"required,min=5"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateUserHandler struct {
	uow  iuow.UnitOfWork
	repo *adapters.UserRepository
}

func NewCreateUserHandler(uow iuow.UnitOfWork) CreateUserHandler {
	if uow == nil {
		panic("NewCreateUserHandler requires uow")
	}

	return CreateUserHandler{uow, uow.GetRepository("UserRepository").(*adapters.UserRepository)}
}

func (h CreateUserHandler) Handle(ctx echox.Context, cmd CreateUser) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	if _, err = h.repo.FindByUsernameOrEmail(ctx.RequestContext(), cmd.Username, nil); err == nil {
		return errors.ERR_USERNAME_EXISTS
	}

	r, err := domain.NewUser(cmd.Username, cmd.Password, cmd.FullName, cmd.Email)
	if err != nil {
		return err
	}

	if err := h.repo.Create(ctx.RequestContext(), &r); err != nil {
		return err
	}
	return nil
}
