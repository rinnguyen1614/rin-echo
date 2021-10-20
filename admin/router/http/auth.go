package http

import (
	"rin-echo/admin/app/command"
	echox "rin-echo/common/echo"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) RegisterAuthRouter(g *echo.Group) {
	g.POST("/login", echox.WrapHandler(h.Login))
	g.POST("/register", echox.WrapHandler(h.Register))
}

func (h HttpServer) Login(c echox.Context) error {
	var cmd command.Login
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	err := h.app.Commands.Login.Handle(c, cmd)

	if err != nil {
		return err
	}

	token, err := h.app.Queries.TokenUser.Handle(c, cmd.Username)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}

func (h HttpServer) Register(c echox.Context) error {
	var cmd command.Register
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	err := h.app.Commands.Register.Handle(c, cmd)

	if err != nil {
		return err
	}

	token, err := h.app.Queries.TokenUser.Handle(c, cmd.Username)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}
