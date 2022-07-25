package handler

import (
	"strings"

	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	_ "github.com/rinnguyen1614/rin-echo/internal/system/app/model/response"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/service"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"github.com/rinnguyen1614/rin-echo/internal/system/inject"

	core "github.com/rinnguyen1614/rin-echo/internal/core"

	"github.com/rinnguyen1614/rin-echo/internal/core/auth/jwt"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	_ "github.com/rinnguyen1614/rin-echo/internal/core/echo/models"
	rquery "github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	fileutil "github.com/rinnguyen1614/rin-echo/internal/core/utils/file"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

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
		return core.NewRinError("avatar_invalid", "The profile picture must be an image.")
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
// @Success      200  {object}  models.Response "{"data": {}}"
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
// @Param        data body request.VerifyEmail true "verify email"
// @Success      200  {object}  models.Response "{"data": {}}"
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
// @Param        data body request.ChangePhone true "change phone"
// @Success      200  {object}  models.Response "{"data": {}}"
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
// @Param        data body request.VerifyPhone true "verify phone"
// @Success      200  {object}  models.Response "{"data": {}}"
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
