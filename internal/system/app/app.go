package app

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/manager"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/handler"

	"github.com/casbin/casbin/v2"
	"github.com/rinnguyen1614/rin-echo/internal/core/auth"
	"github.com/rinnguyen1614/rin-echo/internal/core/cache"
	"github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	setting_adapter "github.com/rinnguyen1614/rin-echo/internal/core/setting/adapter"
	setting_scope "github.com/rinnguyen1614/rin-echo/internal/core/setting/scope"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils/file"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters"
	"github.com/rinnguyen1614/rin-echo/internal/system/config"
	"github.com/rinnguyen1614/rin-echo/internal/system/inject"
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
	AuditLogHandler handler.AuditLogHandler

	FileHandler handler.FileHandler
}

func NewApplication(
	config *config.Config,
	db *gorm.DB,
	cacheManager *cache.CacheManager,
	casbin *casbin.SyncedEnforcer,
	logger *zap.Logger,
	restQuery rest_query.RestQuery,
	auther auth.Auther,
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
		AuditLogHandler: handler.NewAuditLogHandler(uow, logger, restQuery, settingProvider, validator),

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
			Name:             "system.user.verify_email.body",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.verify_email.subject",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
			DefaultValue:     "Rin verification code",
		},
		{
			Name:             "system.user.verify_phone.body",
			AllowedProviders: []string{setting_scope.GlobalSettingProviderName},
			VisibleToClients: true,
		},
		{
			Name:             "system.user.verify_phone.subject",
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
			DefaultValue:     utils.ToString(file.MB * 50),
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
