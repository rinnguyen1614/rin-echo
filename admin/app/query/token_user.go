package query

import (
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/inject"
	iuow "rin-echo/common/uow/interfaces"

	echox "rin-echo/common/echo"
)

type TokenUserHandler struct {
	uow iuow.UnitOfWork
}

func NewTokenUserHandler(uow iuow.UnitOfWork) TokenUserHandler {
	if uow == nil {
		panic("NewTokenUserHandler requires uow")
	}

	return TokenUserHandler{uow}
}

func (h TokenUserHandler) Handle(ctx echox.Context, username string) (interface{}, error) {
	var (
		uow      = h.uow.WithContext(ctx.RequestContext())
		repoUser = repository.NewUserRepository(uow.DB())
	)

	u, err := repoUser.FindByUsernameOrEmail(username, nil)
	if err != nil {
		return nil, err
	}

	token, err := inject.GetAuther().Token(ctx.Request().Context(), map[string]interface{}{
		"FullName": u.FullName,
		"Email":    u.Email,
		"ID":       u.ID,
		"UUID":     u.UUID,
		"Username": u.Username,
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
