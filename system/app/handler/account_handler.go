package handler

import (
	"rin-echo/common"
	"rin-echo/common/auth/jwt"
	echox "rin-echo/common/echo"
	_ "rin-echo/common/echo/models"
	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/model"
	"rin-echo/common/setting"
	iuow "rin-echo/common/uow/interfaces"
	fileutil "rin-echo/common/utils/file"
	"rin-echo/common/validation"
	"rin-echo/system/app/model/request"
	_ "rin-echo/system/app/model/response"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"
	"rin-echo/system/inject"
	"strings"

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
	validator *validation.Validator,
	auther *jwt.JWT) AccountHandler {

	return AccountHandler{
		Handler:     echox.NewHandler(logger, restQuery, settingProvider, validator),
		service:     service.NewAccountService(uow, settingProvider, logger, auther),
		userService: service.NewUserService(uow, permissionManager, settingProvider, logger),
	}
}

// @Summary      Login
// @Description  Login
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.Login true "Login"
// @Success      200  {object}  models.Response{data=jwt.Token} "{"data": {}}"
// @Router       /account/login [post]
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

// @Summary      Change password
// @Description  Change password
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.ChangePassword true "Change password"
// @Success      200  {object}  models.Response{data=jwt.Token} "{"data": {}}"
// @Router       /account/password [put]
// @Security Bearer
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

// @Summary      Profile
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  models.Response{data=response.Profile} "{"data": {}}"
// @Router       /account/profile [get]
// @Security Bearer
func (h AccountHandler) Profile(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}

// @Summary      Update Profile
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.UpdateProfile true "update profile"
// @Success      200  {object}  models.Response{data=response.Profile} "{"data": {}}"
// @Router       /account/profile [put]
// @Security Bearer
func (h AccountHandler) UpdateProfile(c echox.Context) error {
	session := c.MustSession()
	var cmd request.UpdateProfile
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := c.Validate(cmd); err != nil {
		return err
	}

	err := h.service.WithContext(c).UpdateProfile(session.UserID(), cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(session.UserID()))
	return nil
}

// @Summary      Change avatar
// @Tags         account
// @Accept       multipart/form-data
// @Produce      application/json
// @Param 		 file formData file true "File type is image type"
// @Success      200  {object}  models.Response{data=response.Profile} "{"data": {}}"
// @Router       /account/avatar [put]
// @Security Bearer
func (h AccountHandler) ChangeAvatar(c echox.Context) error {
	session := c.MustSession()
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// check file type is an image type
	mimeType, err := fileutil.GetMimeTypeFromFileHeader(file)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(mimeType, "image") {
		return common.NewRinError("avatar_invalid", "The profile picture must be an image.")
	}

	f, err := h.service.WithContext(c).ChangeAvatar(session.UserID(), file)
	if err != nil {
		return err
	}
	echox.OKWithData(c, f)
	return nil
}

// @Summary      Change Email
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.ChangeEmail true "change email"
// @Success      200  {object}
// @Router       /account/profile [put]
// @Security Bearer
func (h AccountHandler) ChangeEmail(c echox.Context) error {
	session := c.MustSession()
	var cmd request.ChangeEmail
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := c.Validate(cmd); err != nil {
		return err
	}

	err := h.service.WithContext(c).ChangeEmail(session.UserID(), cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(session.UserID()))
	return nil
}

// @Summary      Verify Email
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.VerifyEmail true "change email"
// @Success      200  {object}
// @Router       /account/profile [put]
// @Security Bearer
func (h AccountHandler) VerifyEmail(c echox.Context) error {
	session := c.MustSession()
	var cmd request.VerifyEmail
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := c.Validate(cmd); err != nil {
		return err
	}

	err := h.service.WithContext(c).VerifyEmail(session.UserID(), cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(session.UserID()))
	return nil
}

// @Summary      Change Phone
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.ChangePhone true "change email"
// @Success      200  {object}
// @Router       /account/profile [put]
// @Security Bearer
func (h AccountHandler) ChangePhone(c echox.Context) error {
	session := c.MustSession()
	var cmd request.ChangePhone
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := c.Validate(cmd); err != nil {
		return err
	}

	err := h.service.WithContext(c).ChangePhone(session.UserID(), cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(session.UserID()))
	return nil
}

// @Summary      Verify Phone
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Param        data body request.VerifyPhone true "change email"
// @Success      200  {object}
// @Router       /account/profile [put]
// @Security Bearer
func (h AccountHandler) VerifyPhone(c echox.Context) error {
	session := c.MustSession()
	var cmd request.VerifyPhone
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := c.Validate(cmd); err != nil {
		return err
	}

	err := h.service.WithContext(c).VerifyPhone(session.UserID(), cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(session.UserID()))
	return nil
}

// @Summary      Get menus
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  models.Response{data=response.UserMenus} "{"data": {}}"
// @Router       /account/menus [get]
// @Security Bearer
func (h AccountHandler) Menus(c echox.Context) error {
	session := c.MustSession()
	result, err := h.service.WithContext(c).FindMenuTrees(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// @Summary      Get permissions
// @Tags         account
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  models.Response{data=response.UserPermissions} "{"data": {}}"
// @Router       /account/permissions [get]
// @Security Bearer
func (h AccountHandler) Permissions(c echox.Context) error {
	session := c.MustSession()
	result, err := h.service.WithContext(c).FindPermissions(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

func (h AccountHandler) Setting(c echox.Context) error {
	return nil
}
