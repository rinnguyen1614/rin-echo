package echo

import (
	"rin-echo/common/setting"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type (
	Service struct {
		Uow             iuow.UnitOfWork
		SettingProvider setting.Provider
		Localizer       *i18n.Localizer
		Logger          *zap.Logger
		ctx             Context
	}
)

func NewService(uow iuow.UnitOfWork, settingProvider setting.Provider, logger *zap.Logger) *Service {
	return &Service{
		Logger:          logger,
		Uow:             uow,
		SettingProvider: settingProvider,
	}
}

func NewServiceWithContext(ctx Context, uow iuow.UnitOfWork, settingProvider setting.Provider, logger *zap.Logger) Service {
	return Service{
		ctx:             ctx,
		Uow:             uow.WithContext(ctx.RequestContext()),
		SettingProvider: settingProvider,
		Logger:          logger,
	}
}

func (s *Service) Context() Context {
	return s.ctx
}

func (s *Service) WithContext(ctx Context) *Service {
	if ctx == nil {
		panic("nil context")
	}
	s2 := new(Service)
	*s2 = *s
	s2.ctx = ctx
	s2.Uow = s.Uow.WithContext(ctx.RequestContext())
	s2.SettingProvider = s.SettingProvider.WithContext(ctx.RequestContext())

	localizer, _ := ctx.Localizer()
	s2.Localizer = localizer
	return s2
}

func (s Service) Translate(msg string) string {
	return s.TranslateWithDefaultMsg(msg, msg)
}

func (s Service) TranslateWithDefaultMsg(msg, defaultMsg string) string {
	return utils.Translate(s.Localizer, msg, defaultMsg)
}
