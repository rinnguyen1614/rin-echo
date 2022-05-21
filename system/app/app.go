package app

import (
	"rin-echo/common/auth/jwt"
	"rin-echo/common/cache"
	"rin-echo/common/config"
	"rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/setting"
	setting_adapter "rin-echo/common/setting/adapter"
	setting_scope "rin-echo/common/setting/scope"
	"rin-echo/common/utils"
	"rin-echo/common/utils/file"
	"rin-echo/common/validation"
	"rin-echo/system/adapters"
	"rin-echo/system/adapters/manager"
	"rin-echo/system/app/handler"
	"rin-echo/system/inject"

	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	settingDefinitions []setting.SettingDefinition
)

type Application struct {
	SettingHandler  handler.SettingHandler
	AccountHandler  handler.AccountHandler
	ResourceHandler handler.ResourceHandler
	MenuHandler     handler.MenuHandler
	RoleHandler     handler.RoleHandler
	UserHandler     handler.UserHandler

	FileHandler handler.FileHandler
}

func NewApplication(
	config *config.Config,
	db *gorm.DB,
	cacheManager *cache.CacheManager,
	casbin *casbin.SyncedEnforcer,
	logger *zap.Logger,
	restQuery rest_query.RestQuery,
	auther *jwt.JWT,
	validator *validation.Validator,
) Application {

	uow := adapters.NewUnitOfWork(db)
	permissionManager := manager.NewPermissionManager(casbin)

	settingAdapter := setting_adapter.NewDefaultAdapter(uow, cacheManager.Get(inject.SettingCacheName))
	settingProvider := setting.New(
		settingDefinitions,
		[]setting_scope.ScopeProvider{
			setting_scope.NewUserScopeProvider(settingAdapter),
			setting_scope.NewGlobalScopeProvider(settingAdapter),
		},
		utils.NewEncryption(config.App.Encryption),
	)

	return Application{
		SettingHandler:  handler.NewSettingHandler(uow, logger, restQuery, settingProvider, validator),
		AccountHandler:  handler.NewAccountHandler(uow, permissionManager, logger, restQuery, settingProvider, validator, auther),
		ResourceHandler: handler.NewResourceHandler(uow, permissionManager, logger, restQuery, settingProvider, validator),
		MenuHandler:     handler.NewMenuHandler(uow, permissionManager, logger, restQuery, settingProvider, validator),
		RoleHandler:     handler.NewRoleHandler(uow, permissionManager, logger, restQuery, settingProvider, validator),
		UserHandler:     handler.NewUserHandler(uow, permissionManager, logger, restQuery, settingProvider, validator),

		FileHandler: handler.NewFileHandler(uow, permissionManager, logger, restQuery, settingProvider, validator),
	}
}

func init() {
	settingDefinitions = append(settingDefinitions, definationsForApplication()...)
	settingDefinitions = append(settingDefinitions, definationsForUser()...)
}

func definationsForApplication() []setting.SettingDefinition {
	return []setting.SettingDefinition{
		{
			Name:             "app.name",
			DefaultValue:     "Rin Application",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "app.logo.path",
			DefaultValue:     "",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.require_email_confirmation_for_login",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.password.min_length",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "6",
		},
		{
			Name:             "system.user.password.require_digit",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "true",
		},
		{
			Name:             "system.user.password.require_lower_case",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "true",
		},
		{
			Name:             "system.user.password.require_upper_case",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "true",
		},
		{
			Name:             "system.user.password.require_special_character",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "true",
		},
		{
			Name:             "system.user.lockout.enable",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.lockout.time_locked",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.lockout.max_failed",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "files.public.path",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			DefaultValue:     "static/public",
			VisibleToClients: true,
		},
		{
			Name:             "files.upload.path",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			DefaultValue:     "static/upload",
			VisibleToClients: true,
		},
		{
			Name:             "files.upload.max_size",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			DefaultValue:     utils.ToString(file.MB * 10),
			VisibleToClients: true,
		},
		{
			Name:             "files.avatar_path",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			DefaultValue:     "static/public/avatar",
			VisibleToClients: true,
		},
	}
}

func definationsForUser() []setting.SettingDefinition {
	return []setting.SettingDefinition{
		{
			Name:             "app.language",
			DefaultValue:     "en",
			AllowedProviders: []string{setting_scope.UserSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "app.theme.background",
			DefaultValue:     "#fff",
			AllowedProviders: []string{setting_scope.UserSettingProviderName},
			VisibleToClients: true,
		},
	}
}
