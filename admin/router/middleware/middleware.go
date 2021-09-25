package middleware

import (
	"rin-echo/admin/inject"
	echox "rin-echo/common/echo"
	mwx "rin-echo/common/echo/middleware"

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
