package app

import (
	"rin-echo/common/auth/jwt"
	"rin-echo/common/echo/models/query/rest_query"
	"rin-echo/system/adapters"
	"rin-echo/system/adapters/manager"
	"rin-echo/system/app/handler"

	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	AccountHandler  handler.AccountHandler
	ResourceHandler handler.ResourceHandler
	MenuHandler     handler.MenuHandler
	RoleHandler     handler.RoleHandler
	UserHandler     handler.UserHandler
}

func NewApplication(db *gorm.DB, casbin *casbin.SyncedEnforcer, logger *zap.Logger, restQuery rest_query.RestQuery, auther *jwt.JWT) Application {

	uow := adapters.NewUnitOfWork(db)
	permissionManager := manager.NewPermissionManager(casbin)

	return Application{
		AccountHandler:  handler.NewAccountHandler(uow, permissionManager, logger, restQuery, auther),
		ResourceHandler: handler.NewResourceHandler(uow, permissionManager, logger, restQuery),
		MenuHandler:     handler.NewMenuHandler(uow, permissionManager, logger, restQuery),
		RoleHandler:     handler.NewRoleHandler(uow, permissionManager, logger, restQuery),
		UserHandler:     handler.NewUserHandler(uow, permissionManager, logger, restQuery),
	}
}
