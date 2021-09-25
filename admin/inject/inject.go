package inject

import (
	"rin-echo/common/auth/jwt"
	"rin-echo/common/config"
	"rin-echo/common/echo/models/query/rest_query"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var service *Service

func init() {
	if service == nil {
		service = &Service{}
	}
}

type Service struct {
	cfg      *config.Config
	db       *gorm.DB
	logger   *zap.Logger
	auther   *jwt.JWT
	echo     *echo.Echo
	enforcer *casbin.Enforcer
	i18n     *i18n.Bundle
	query    rest_query.RestQuery
}

func GetConfig() *config.Config {
	if service.cfg == nil {
		c := config.ReadFromYaml("conf/config.yaml")
		service.cfg = &c
	}
	return service.cfg
}
