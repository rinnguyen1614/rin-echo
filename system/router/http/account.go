package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterAccountRouter(g *echo.Group) {

	router := g.Group("/account")
	{
		// without authentication
		router.POST("/login", echox.WrapHandler(h.app.AccountHandler.Login))
		router.POST("/register", echox.WrapHandler(h.app.AccountHandler.Register))

		// with authentication
		middleware.AddJWTMiddleware(router)
		router.PUT("/logout", echox.WrapHandler(h.app.AccountHandler.Logout))
		router.PUT("/password", echox.WrapHandler(h.app.AccountHandler.ChangePassword))
		router.GET("/token_info", echox.WrapHandler(h.app.AccountHandler.TokenInfo))
		router.GET("/profile", echox.WrapHandler(h.app.AccountHandler.Profile))
		router.PUT("/profile", echox.WrapHandler(h.app.AccountHandler.UpdateProfile))
		router.PUT("/email", echox.WrapHandler(h.app.AccountHandler.ChangeEmail))
		router.POST("/verify_email", echox.WrapHandler(h.app.AccountHandler.VerifyEmail))
		router.PUT("/phone", echox.WrapHandler(h.app.AccountHandler.ChangePhone))
		router.POST("/verify_phone", echox.WrapHandler(h.app.AccountHandler.VerifyPhone))
		router.PUT("/avatar", echox.WrapHandler(h.app.AccountHandler.ChangeAvatar))
		router.GET("/menus", echox.WrapHandler(h.app.AccountHandler.Menus))
		router.GET("/permissions", echox.WrapHandler(h.app.AccountHandler.Permissions))
	}
}
