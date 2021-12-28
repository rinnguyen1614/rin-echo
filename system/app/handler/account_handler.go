package handler

import (
	"rin-echo/common/auth/jwt"
	echox "rin-echo/common/echo"
	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/setting"
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

func NewAccountHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	auther *jwt.JWT) AccountHandler {

	return AccountHandler{
		Handler:     echox.NewHandler(logger, restQuery, settingProvider),
		service:     service.NewAccountService(uow, settingProvider, logger, auther),
		userService: service.NewUserService(uow, permissionManager, settingProvider, logger),
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

	// validator, ok := c.Echo().Validator.(*validation.Validator)
	// if ok {
	// 	var (
	// 		minLength                                                                  int
	// 		requireDigit, requireLowercase, requireUppercase, requiredSpecialCharacter bool
	// 	)

	// 	v, err := h.SettingProvider.Get("system.user.password.min_length")
	// 	if err == nil {
	// 		minLength, _ = strconv.Atoi(v)
	// 	}

	// 	v, err = h.SettingProvider.Get("system.user.password.require_digit")
	// 	if err == nil && v == "true" {
	// 		requireDigit = true
	// 	}

	// 	v, err = h.SettingProvider.Get("system.user.password.require_lower_case")
	// 	if err == nil && v == "true" {
	// 		requireLowercase = true
	// 	}

	// 	v, err = h.SettingProvider.Get("system.user.password.require_upper_case")
	// 	if err == nil && v == "true" {
	// 		requireUppercase = true
	// 	}

	// 	v, err = h.SettingProvider.Get("system.user.password.require_special_character")
	// 	if err == nil && v == "true" {
	// 		requiredSpecialCharacter = true
	// 	}

	// 	//validator.RegisterValidationForPassword(minLength, requireDigit, requireLowercase, requireUppercase, requiredSpecialCharacter)
	// }

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

func (h AccountHandler) Setting(c echox.Context) error {
	return nil
}
