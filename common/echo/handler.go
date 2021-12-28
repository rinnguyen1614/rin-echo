package echo

import (
	"rin-echo/common"
	"rin-echo/system/errors"
	"strconv"

	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/setting"

	"go.uber.org/zap"
)

var (
	ErrRequestIDRequired = common.NewRinError("request_id_required", "request_id_required")
	ErrRequestIDInvalid  = common.NewRinError("request_id_invalid", "request_id_invalid")
)

type Handler struct {
	Logger *zap.Logger

	RestQuery       rquery.RestQuery
	SettingProvider setting.Provider
}

func NewHandler(logger *zap.Logger, restQuery rquery.RestQuery, settingProvider setting.Provider) Handler {
	return Handler{
		Logger:          logger,
		RestQuery:       restQuery,
		SettingProvider: settingProvider,
	}
}

func CheckRequestIDParam(src string) (uint, error) {
	if len(src) == 0 {
		return 0, errors.ErrRequestIDRequired
	}
	id, err := strconv.Atoi(src)
	if err != nil {
		return 0, errors.ErrRequestIDInvalid
	}
	return uint(id), nil
}
