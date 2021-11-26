package echo

import (
	"net/http"
	"rin-echo/common/echo/models"

	"github.com/labstack/echo/v4"
)

var (
	HeaderRequestID       = "request-id"
	HeaderClientRequestID = "client-request-id"
	HeaderDeviceID        = "device-id"
	HeaderDeviceName      = "device-name"
)

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

func WrapHandlerWithOperation(handler HandlerFunc, operationName, operationMethod string) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := MustContext(c)
		cc.SetOperation(operationName, operationMethod)
		return handler(cc)
	}
}
