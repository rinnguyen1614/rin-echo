package middleware

import (
	cdm "rin-echo/common/domain"
	echox "rin-echo/common/echo"
	mwx "rin-echo/common/echo/middleware"
	"rin-echo/system/domain"
	"rin-echo/system/inject"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
