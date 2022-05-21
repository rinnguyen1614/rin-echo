package handler

import (
	echox "rin-echo/common/echo"
	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/setting"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
	"rin-echo/common/validation"
	"rin-echo/system/app/model/request"
	"strconv"

	"go.uber.org/zap"
)

type SettingHandler struct {
	echox.Handler
}

func NewSettingHandler(uow iuow.UnitOfWork,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) SettingHandler {
	return SettingHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
	}
}

func (h SettingHandler) Set(c echox.Context) error {
	var cmd request.Setting
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}

	provider := h.SettingProvider.WithContext(c.RequestContext())

	err := provider.Set("system.user.require_email_confirmation_for_login", utils.ToString(cmd.Security.RequireEmailConfirmationForLogin))

	if err != nil {
		return err
	}

	echox.OKWithData(c, cmd)
	return nil
}

func (h SettingHandler) Get(c echox.Context) error {
	provider := h.SettingProvider.WithContext(c.RequestContext())
	result := request.Setting{}

	requireEmail, _ := provider.Get("system.user.require_email_confirmation_for_login")
	result.Security.RequireEmailConfirmationForLogin, _ = strconv.ParseBool(requireEmail)

	echox.OKWithData(c, result)
	return nil
}
