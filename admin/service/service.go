package service

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/app"
	"rin-echo/admin/app/command"
	"rin-echo/admin/app/query"
	"rin-echo/admin/inject"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

func NewApplication(db *gorm.DB, casbin *casbin.SyncedEnforcer) app.Application {

	uow := adapters.NewUnitOfWork(db)
	rbac := adapters.NewRBACCasbin(inject.GetCasbin())

	return app.Application{
		Commands: app.Commands{
			Login:          command.NewLoginHandler(uow),
			Register:       command.NewRegisterHandler(uow, &rbac),
			CreateResource: command.NewCreateResourceHandler(uow),
			CreateMenu:     command.NewCreateMenuHandler(uow),
			CreateRole:     command.NewCreateRoleHandler(uow, &rbac),
			CreateUser:     command.NewCreateUserHandler(uow, &rbac),
		},
		Queries: app.Queries{
			TokenUser: query.NewTokenUserHandler(uow),
			FindUsers: query.NewFindUsersHandler(uow),
		},
	}
}
