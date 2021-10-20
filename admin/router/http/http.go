package http

import (
	"net/http"
	"rin-echo/admin/app"
	"rin-echo/admin/inject"
	"rin-echo/admin/router/middleware"
	echox "rin-echo/common/echo"
	"rin-echo/common/utils"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}

func (h *HttpServer) RegisterRouter(g *echo.Group) {
	middleware.AddDefaultMiddleware(g)

	test(g)

	h.RegisterAuthRouter(g)
	h.RegisterResourceRouter(g)
	h.RegisterMenuRouter(g)
	h.RegisterUserRouter(g)
}

func test(g *echo.Group) {
	g.GET("/test_language", func(c echo.Context) error {
		ll, _ := echox.MustContext(c).Localizer()

		username, _ := ll.Localize(&i18n.LocalizeConfig{
			MessageID: "token_malformed",
		})

		return c.JSON(http.StatusOK, username)
	})

	g.GET("/test_token", func(c echo.Context) error {
		token, _ := inject.GetAuther().Token(c.Request().Context(), map[string]interface{}{
			"ID":       uint(1),
			"UUID":     utils.MustUUID(),
			"Username": "administrator",
		})

		return c.JSON(http.StatusOK, token)
	})

	ag := g.Group("/auth")
	{
		middleware.AddJWTMiddleware(ag)
		middleware.AddCasbinMiddleware(ag)
		ag.GET("/info", func(c echo.Context) error {
			echox.MustContext(c).MustSession()

			return c.JSON(http.StatusOK, echox.MustContext(c).MustSession())
		})

		ag.POST("/valid", func(c echo.Context) error {
			type Account struct {
				Username string `json:"username" validate:"required,min=5"`
				Password string `json:"password" validate:"required,min=5"`
			}

			var cmd Account

			if err := c.Bind(&cmd); err != nil {
				return err
			}
			if err := c.Validate(cmd); err != nil {
				return err
			}

			return c.JSON(http.StatusOK, cmd)
		})
	}
}
