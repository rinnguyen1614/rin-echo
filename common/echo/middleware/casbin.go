package middleware

import (
	"rin-echo/common/auth"
	casbinx "rin-echo/common/casbin"
	echox "rin-echo/common/echo"
	"rin-echo/common/utils"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	CasbinConfig struct {
		Skipper             middleware.Skipper
		SubObjectActionFunc SubObjectActionFunc
		Enforcer            *casbin.Enforcer
	}

	SubObjectActionFunc func(c echo.Context) (sub, obj, act interface{}, err error)
)

var (
	DefaultCasbinConfig = CasbinConfig{
		Skipper:             middleware.DefaultSkipper,
		SubObjectActionFunc: casbinDefaultSubObjectActionFunc,
	}
)

func Casbin(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	c := DefaultCasbinConfig
	c.Enforcer = enforcer
	return CasbinWithConfig(c)
}

func CasbinWithConfig(config CasbinConfig) echo.MiddlewareFunc {
	if config.Enforcer == nil {
		panic("casbin requires Enforcer")
	}
	if config.Skipper == nil {
		config.Skipper = DefaultCasbinConfig.Skipper
	}
	if config.SubObjectActionFunc == nil {
		config.SubObjectActionFunc = DefaultCasbinConfig.SubObjectActionFunc
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			sub, obj, act, err := casbinDefaultSubObjectActionFunc(c)

			allowed, err := config.Enforcer.Enforce(sub, obj, act)
			if err != nil {
				return casbinx.ERR_NOT_PERMISSION
			}

			if !allowed {
				return casbinx.ERR_NOT_PERMISSION
			}
			return next(c)
		}
	}
}

func casbinDefaultSubObjectActionFunc(c echo.Context) (sub, obj, act interface{}, err error) {
	cc := echox.MustContext(c)
	session, err := cc.Session()
	if err != nil {
		return nil, nil, nil, auth.NewAuthenticationErrorWithInner(err, "not_found_current_session", "Can get current session to enforce")
	}

	sub = utils.ToString(session.UserID())
	obj = c.Request().URL.RequestURI()
	act = c.Request().Method
	return
}
