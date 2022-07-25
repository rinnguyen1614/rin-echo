package handler

import (
	"strconv"

	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	rquery "github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

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
