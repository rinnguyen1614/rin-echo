package middleware

import (
	cdm "github.com/rinnguyen1614/rin-echo/internal/core/domain"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	mwx "github.com/rinnguyen1614/rin-echo/internal/core/echo/middleware"
	"github.com/rinnguyen1614/rin-echo/internal/system/inject"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

func AddDefaultMiddleware(g *echo.Group) {
	g.Use(mwx.Logger())
	g.Use(mwx.Recover())
	g.Use(middleware.RemoveTrailingSlash())

	g.Use(mwx.Localizer(inject.GetI18n()))
}

func AddJWTMiddleware(g *echo.Group) {
	g.Use(mwx.JWTWithConfig(mwx.JWTConfig{
		Auther: inject.GetAuther(),
		WrapSessionContext: func(c echo.Context, claims jwt.Claims) {
			cc := echox.MustContext(c)
			cc.SetSession(claims.(*inject.Claims))
		},
	}))
}

func AddCasbinMiddleware(g *echo.Group) {
	g.Use(mwx.Casbin(inject.GetCasbin()))
}

func AddRequestLoggerMiddleware(g *echo.Group) {
	var logFunc = func(c echo.Context, auditLog cdm.AuditLog) error {
		cc := echox.MustContext(c)
		db := inject.GetDB().WithContext(cc.RequestContext())
		return db.Create(&domain.AuditLog{AuditLog: auditLog}).Error
	}

	g.Use(mwx.RequestLogger(inject.GetConfig().App.AppName, logFunc))
}
