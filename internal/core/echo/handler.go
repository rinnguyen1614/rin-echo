package echo

import (
	"strconv"

	rquery "github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"go.uber.org/zap"
)

var (
	ErrRequestIDRequired = core.NewRinError("request_id_required", "request_id_required")
	ErrRequestIDInvalid  = core.NewRinError("request_id_invalid", "request_id_invalid")
)

type Handler struct {
	Logger *zap.Logger

	RestQuery       rquery.RestQuery
	SettingProvider setting.Provider

	Validator *validation.Validator
}

func NewHandler(logger *zap.Logger, restQuery rquery.RestQuery, settingProvider setting.Provider, validator *validation.Validator) Handler {
	return Handler{
		Logger:          logger,
		RestQuery:       restQuery,
		SettingProvider: settingProvider,
		Validator:       validator,
	}
}

func CheckRequestIDParam(src string) (uint, error) {
	if len(src) == 0 {
		return 0, ErrRequestIDRequired
	}
	id, err := strconv.Atoi(src)
	if err != nil {
		return 0, ErrRequestIDInvalid
	}
	return uint(id), nil
}
