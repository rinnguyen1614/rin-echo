package query

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/inject"
	iuow "rin-echo/common/uow/interfaces"

	echox "rin-echo/common/echo"
)

type TokenUserHandler struct {
	uow  iuow.UnitOfWork
	repo *adapters.UserRepository
}

func NewTokenUserHandler(uow iuow.UnitOfWork) TokenUserHandler {
	if uow == nil {
		panic("NewTokenUserHandler requires uow")
	}

	return TokenUserHandler{uow, uow.GetRepository("UserRepository").(*adapters.UserRepository)}
}

func (h TokenUserHandler) Handle(ctx echox.Context, username string) (interface{}, error) {
	u, err := h.repo.FindByUsernameOrEmail(ctx.RequestContext(), username, nil)
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
