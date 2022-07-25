package scope

import (
	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting/adapter"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
)

//
type UserScopeProvider struct {
	ctx core.Context

	*scopeProvider
}

func NewUserScopeProvider(adapter adapter.Adapter) *UserScopeProvider {
	return &UserScopeProvider{
		scopeProvider: newScopeProvider(UserSettingProviderName, adapter),
	}
}

func NewUserScopeProviderWithContext(adapter adapter.Adapter, ctx core.Context) *UserScopeProvider {
	return &UserScopeProvider{
		scopeProvider: newScopeProvider(UserSettingProviderName, adapter),
		ctx:           ctx,
	}
}
func (s *UserScopeProvider) WithContext(ctx core.Context) ScopeProvider {
	if ctx == nil {
		panic("nil context")
	}
	s2 := new(UserScopeProvider)
	*s2 = *s
	s2.ctx = ctx
	s2.adapter = s.adapter.WithContext(ctx)
	return s2
}

func (s UserScopeProvider) GetOrInit(name string) string {
	return s.scopeProvider.adapter.GetOrInit(name, s.name, utils.ToString(s.ctx.MustSession().UserID()))
}

func (s UserScopeProvider) GetMulti(names []string) map[string]string {
	return s.scopeProvider.adapter.GetMulti(names, s.name, utils.ToString(s.ctx.MustSession().UserID()))
}

func (s UserScopeProvider) GetAll() map[string]string {
	return s.scopeProvider.adapter.GetAll(s.name, utils.ToString(s.ctx.MustSession().UserID()))
}

func (s UserScopeProvider) Set(name, value string) error {
	return s.scopeProvider.adapter.Set(name, value, s.name, utils.ToString(s.ctx.MustSession().UserID()))
}

func (s UserScopeProvider) Delete(name string) error {
	return s.scopeProvider.adapter.Delete(name, s.name, utils.ToString(s.ctx.MustSession().UserID()))
}
