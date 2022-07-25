package scope

import (
	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting/adapter"
)

type GlobalScopeProvider struct {
	*scopeProvider
}

func NewGlobalScopeProvider(adapter adapter.Adapter) *GlobalScopeProvider {
	return &GlobalScopeProvider{
		scopeProvider: newScopeProvider(GlobalSettingProviderName, adapter),
	}
}

func (s *GlobalScopeProvider) WithContext(ctx core.Context) ScopeProvider {
	if ctx == nil {
		panic("nil context")
	}
	s2 := new(GlobalScopeProvider)
	*s2 = *s
	s2.adapter = s.adapter.WithContext(ctx)
	return s2
}

func (s GlobalScopeProvider) GetOrInit(name string) string {
	return s.scopeProvider.adapter.GetOrInit(name, s.name, "")
}

func (s GlobalScopeProvider) GetMulti(names []string) map[string]string {
	return s.scopeProvider.adapter.GetMulti(names, s.name, "")
}

func (s GlobalScopeProvider) GetAll() map[string]string {
	return s.scopeProvider.adapter.GetAll(s.name, "")
}

func (s GlobalScopeProvider) Set(name, value string) error {
	return s.scopeProvider.adapter.Set(name, value, s.name, "")
}

func (s GlobalScopeProvider) Delete(name string) error {
	return s.scopeProvider.adapter.Delete(name, s.name, "")
}
