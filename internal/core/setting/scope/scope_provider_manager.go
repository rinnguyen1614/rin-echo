package scope

import (
	"errors"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
)

type ScopeProviderManager struct {
	providers       []ScopeProvider
	providersByName map[string]ScopeProvider
}

// Priority by index of providers
func NewScopeProviderManager(providers []ScopeProvider) *ScopeProviderManager {
	if providers == nil {
		panic("requires providers")
	}

	var mProviders = make(map[string]ScopeProvider)
	for _, provider := range providers {
		mProviders[provider.Name()] = provider
		if _, ok := mProviders[provider.Name()]; ok {

		}
	}

	return &ScopeProviderManager{
		providers:       providers,
		providersByName: mProviders,
	}
}

func (s *ScopeProviderManager) WithContext(ctx core.Context) *ScopeProviderManager {
	s2 := new(ScopeProviderManager)
	s2.providersByName = make(map[string]ScopeProvider)
	s2.providers = make([]ScopeProvider, len(s.providersByName))

	for i, pro := range s.providers {
		proCtx := pro.WithContext(ctx)
		s2.providers[i] = proCtx
		s2.providersByName[proCtx.Name()] = proCtx
	}

	return s2
}

func (s ScopeProviderManager) Get(names ...string) []ScopeProvider {
	var providers []ScopeProvider
	for _, name := range names {
		if provider, ok := s.providersByName[name]; ok {
			providers = append(providers, provider)
		}
	}

	return providers
}

func (s *ScopeProviderManager) Add(provider ScopeProvider) error {
	var name = provider.Name()
	if _, ok := s.providersByName[name]; ok {
		return errors.New("ScopeProvider has existed")
	}

	s.providers = append(s.providers, provider)
	s.providersByName[name] = provider

	return nil
}

func (s *ScopeProviderManager) Remove(name string) error {
	if _, ok := s.providersByName[name]; !ok {
		return errors.New("ScopeProvider doesn't exist")
	}

	delete(s.providersByName, name)
	// clone from map instead of spliptting slice.
	s.providers = make([]ScopeProvider, len(s.providersByName))
	i := 0
	for _, pro := range s.providersByName {
		s.providers[i] = pro
		i++
	}

	return nil
}

func (s ScopeProviderManager) Providers() []ScopeProvider {
	return s.providers
}
