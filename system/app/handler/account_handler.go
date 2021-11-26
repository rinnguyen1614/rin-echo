package handler

import (
	"rin-echo/common/auth/jwt"
	echox "rin-echo/common/echo"
	rquery "rin-echo/common/echo/models/query/rest_query"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/app/model/request"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"
	"rin-echo/system/inject"

	"go.uber.org/zap"
)

type AccountHandler struct {
	echox.Handler
	service     service.AccountService
	userService service.UserService
}

func NewAccountHandler(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, logger *zap.Logger, restQuery rquery.RestQuery, auther *jwt.JWT) AccountHandler {
	return AccountHandler{
		Handler:     echox.NewHandler(logger, restQuery),
		service:     service.NewAccountService(uow, logger, auther),
		userService: service.NewUserService(uow, permissionManager, logger),
	}
}

func (h AccountHandler) Login(c echox.Context) error {
	var cmd request.Login
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}

	defer func() {

	}()

	token, err := h.service.WithContext(c).Login(cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}

func (h AccountHandler) Logout(c echox.Context) error {
	var cmd request.Login
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}

	token, err := h.service.WithContext(c).Login(cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}

func (h AccountHandler) Register(c echox.Context) error {
	var cmd request.Register
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}

	id, err := h.userService.WithContext(c).CreateDefault(cmd.ToCreateUser())
	if err != nil {
		return err
	}

	token, err := h.service.WithContext(c).Token(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}

func (h AccountHandler) ChangePassword(c echox.Context) error {
	var cmd request.ChangePassword
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	session := c.MustSession().(*inject.Claims)
	cmd.Username = session.Username

	if err := c.Validate(cmd); err != nil {
		return err
	}

	token, err := h.service.WithContext(c).ChangePassword(cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}

func (h AccountHandler) TokenInfo(c echox.Context) error {
	session := c.MustSession()
	echox.OKWithData(c, session)
	return nil
}

func (h AccountHandler) Profile(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}

func (h AccountHandler) UpdateProfile(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}

func (h AccountHandler) Menus(c echox.Context) error {
	session := c.MustSession()
	result, err := h.service.WithContext(c).FindMenuTrees(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}
