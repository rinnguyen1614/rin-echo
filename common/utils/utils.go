package utils

import (
	"reflect"
	"runtime"
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}

func Translate(localizer *i18n.Localizer, msgID, defaultMsg string) string {
	if localizer == nil {
		return defaultMsg
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: msgID,
		DefaultMessage: &i18n.Message{
			ID:    msgID,
			Other: defaultMsg,
		},
	})

	if err != nil {
		return err.Error()
	}

	return msg
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
