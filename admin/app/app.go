package app

import (
	"rin-echo/admin/app/command"
	"rin-echo/admin/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Login    command.LoginHandler
	Register command.RegisterHandler

	// resource
	CreateResource command.CreateResourceHandler

	// menu
	CreateMenu command.CreateMenuHandler

	// role
	CreateRole command.CreateRoleHandler

	// user
	CreateUser command.CreateUserHandler
}

type Queries struct {
	TokenUser query.TokenUserHandler
	FindUsers query.FindUsersHandler
}
