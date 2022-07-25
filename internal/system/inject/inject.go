package inject

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/auth/jwt"
	"github.com/rinnguyen1614/rin-echo/internal/core/cache"
	"github.com/rinnguyen1614/rin-echo/internal/core/config"
	"github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var service *Service

func init() {
	if service == nil {
		service = &Service{}
	}
}

type Service struct {
	cfg       *config.Config
	db        *gorm.DB
	cache     *cache.CacheManager
	logger    *zap.Logger
	auther    *jwt.JWT
	echo      *echo.Echo
	enforcer  *casbin.SyncedEnforcer
	i18n      *i18n.Bundle
	query     rest_query.RestQuery
	validator *validation.Validator
	mail      *gomail.Dialer
}

func GetConfig() *config.Config {
	if service.cfg == nil {
		c := config.ReadFromYaml("conf/config.yaml")
		service.cfg = &c
	}
	return service.cfg
}
