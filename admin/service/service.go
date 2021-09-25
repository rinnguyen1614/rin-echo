package service

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/app"
	"rin-echo/admin/app/command"
	"rin-echo/admin/app/query"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

func NewApplication(db *gorm.DB, casbin *casbin.Enforcer) app.Application {

	uow := adapters.NewUnitOfWork(db)

	return app.Application{
		Commands: app.Commands{
			Login:      command.NewLoginHandler(uow),
			CreateUser: command.NewCreateUserHandler(uow),
		},
		Queries: app.Queries{
			TokenUser: query.NewTokenUserHandler(uow),
			FindUsers: query.NewFindUsersHandler(uow),
		},
	}
}
