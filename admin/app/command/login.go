package command

import (
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type (
	Login struct {
		Username string `validate:"required,min=5"`
		Password string `validate:"required,min=6"`
	}

	LoginHandler struct {
		uow iuow.UnitOfWork
	}
)

func NewLoginHandler(uow iuow.UnitOfWork) LoginHandler {
	if uow == nil {
		panic("newloginhandler requires uow")
	}

	return LoginHandler{uow}
}

func (h LoginHandler) Handle(ctx echox.Context, cmd Login) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow      = h.uow.WithContext(ctx.RequestContext())
		repoUser = repository.NewUserRepository(uow.DB())
	)

	u, err := repoUser.FindByUsernameOrEmail(cmd.Username, nil)
	if err != nil {
		return err
	}

	if !u.CheckPassword(cmd.Password) {
		return errors.ERR_USERNAME_PASSWORD_NOT_MATCH
	}

	return nil
}
