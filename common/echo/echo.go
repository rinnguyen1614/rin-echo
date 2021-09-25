package echo

import (
	"net/http"
	"rin-echo/common/echo/models"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	HeaderRequestID       = "request-id"
	HeaderClientRequestID = "client-request-id"
)

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

func OKWithData(c echo.Context, data interface{}) {
	c.JSON(http.StatusOK, models.NewResponseWithData(data))
}

type HandlerFunc func(c Context) error

func WrapHandler(handler HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := MustContext(c)
		return handler(cc)
	}
}
