package middleware

import (
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	LocalizerConfig struct {
		Header    string
		Skipper   middleware.Skipper
		Bundle    *i18n.Bundle
		FormValue string
	}
)

var (
	DefaultLocalizerConfig = LocalizerConfig{
		Header:    "Accept-Language",
		FormValue: "lang",
		Skipper:   middleware.DefaultSkipper,
	}
)

func Localizer(bundle *i18n.Bundle) echo.MiddlewareFunc {
	c := DefaultLocalizerConfig
	c.Bundle = bundle
	return LocalizerWithConfig(c)
}

func LocalizerWithConfig(config LocalizerConfig) echo.MiddlewareFunc {

	if config.Bundle == nil {
		panic("localizer requies Bundle")
	}

	if config.Header == "" {
		config.Header = DefaultLocalizerConfig.Header
	}

	if config.FormValue == "" {
		config.FormValue = DefaultLocalizerConfig.FormValue
	}

	if config.Skipper == nil {
		config.Skipper = DefaultLocalizerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			lang := c.FormValue(config.FormValue)
			accept := c.Request().Header.Get(config.Header)
			localizer := i18n.NewLocalizer(config.Bundle, lang, accept)

			cc := echox.MustContext(c)
			cc.SetLocalizer(localizer)

			return next(c)
		}
	}
}
