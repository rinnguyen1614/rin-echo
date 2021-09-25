package middleware

import (
	echox "rin-echo/common/echo"
	"rin-echo/common/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	RequestIDConfig struct {
		// request id from client request
		HeaderClient string
		// request id from server response
		Header string

		middleware.RequestIDConfig
	}
)

var (
	DefaultRequestIDConfig = RequestIDConfig{
		HeaderClient: echox.HeaderClientRequestID,
		Header:       echox.HeaderRequestID,
		RequestIDConfig: middleware.RequestIDConfig{
			Skipper:   middleware.DefaultSkipper,
			Generator: generator,
		},
	}
)

func RequestID() echo.MiddlewareFunc {
	return RequestIDWithConfig(DefaultRequestIDConfig)
}

func RequestIDWithConfig(config RequestIDConfig) echo.MiddlewareFunc {
	if config.Header == "" {
		config.Header = DefaultRequestIDConfig.Header
	}
	if config.HeaderClient == "" {
		config.HeaderClient = DefaultRequestIDConfig.HeaderClient
	}
	if config.Skipper == nil {
		config.Skipper = DefaultRequestIDConfig.Skipper
	}
	if config.Generator == nil {
		config.Generator = generator
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			client_rid := req.Header.Get(config.HeaderClient)
			rid := config.Generator()
			if client_rid == "" {
				client_rid = rid
			}

			res.Header().Set(config.HeaderClient, client_rid)
			res.Header().Set(config.Header, rid)
			if config.RequestIDHandler != nil {
				config.RequestIDHandler(c, rid)
			}

			return next(c)
		}
	}
}

func generator() string {
	return utils.MustUUID().String()
}
