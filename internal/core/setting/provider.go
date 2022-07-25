package setting

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/setting/scope"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/thoas/go-funk"
)

// Provider provides the configuration settings
type Provider interface {
	WithContext(ctx core.Context) Provider

	Get(name string) (string, error)
	GetMulti(names []string) (map[string]string, error)
	Set(name, value string) error

	AddDefination(SettingDefinition) error
	RemoveDefination(name string) error

	AddScopeProvider(provider scope.ScopeProvider) error
	RemoveScopeProvider(name string) error
}

type provider struct {
	definitionManager    *SettingDefinitionManager
	scopeProviderManager *scope.ScopeProviderManager
	encryption           *utils.Encryption
	ctx                  core.Context
}

func New(definitions []SettingDefinition, scopeProviders []scope.ScopeProvider, encryption *utils.Encryption) Provider {
	return &provider{
		definitionManager:    NewSettingDefinitionManager(definitions),
		scopeProviderManager: scope.NewScopeProviderManager(scopeProviders),
		encryption:           encryption,
	}

}

func (p *provider) Context() core.Context {
	return p.ctx
}

func (p *provider) WithContext(ctx core.Context) Provider {
	if ctx == nil {
		panic("nil context")
	}
	p2 := new(provider)
	*p2 = *p
	p2.ctx = ctx
	p2.scopeProviderManager = p.scopeProviderManager.WithContext(ctx)

	return p2
}

func (p *provider) Get(name string) (string, error) {
	var (
		definition, err = p.definitionManager.Get(name)
		providers       []scope.ScopeProvider
	)

	if err != nil {
		return "", err
	}

	if len(definition.AllowedProviders) > 0 {
		providers = p.scopeProviderManager.Get(definition.AllowedProviders...)
	} else {
		providers = p.scopeProviderManager.Providers()
	}

	value := p.getFromProviders(providers, definition)
	if value != "" && definition.IsEncrypted {
		value = p.encryption.Decrypt(value)
	}

	return value, nil
}

func (p *provider) GetMulti(names []string) (map[string]string, error) {
	var (
		providers      = p.scopeProviderManager.Providers()
		definitions    = p.definitionManager.GetMulti(names)
		values         = make(map[string]string)
		mapDefinitions = make(map[string]SettingDefinition)
	)

	// init value
	for _, definition := range definitions {
		values[definition.Name] = definition.DefaultValue
		mapDefinitions[definition.Name] = definition
	}

	for _, provider := range providers {
		var (
			definitionFounds []SettingDefinition
			definitionNames  []string
			valueFounds      map[string]string
		)

		for _, definition := range mapDefinitions {
			if funk.Contains(definition.AllowedProviders, provider.Name()) {
				definitionFounds = append(definitionFounds, definition)
				definitionNames = append(definitionNames, definition.Name)
			}
		}

		valueFounds = provider.GetMulti(definitionNames)
		for _, definition := range definitionFounds {
			var (
				name  = definition.Name
				value = valueFounds[name]
			)
			if definition.IsEncrypted {
				value = p.encryption.Decrypt(value)
			}
			values[name] = value
			// delete found definition
			delete(mapDefinitions, name)
		}
	}

	return values, nil
}

func (p provider) getFromProviders(providers []scope.ScopeProvider, definition SettingDefinition) string {
	for _, provider := range providers {
		value := provider.GetOrInit(definition.Name)
		if value != "" {
			return value
		}
	}
	return definition.DefaultValue
}

func (p *provider) Set(name, value string) error {
	var (
		definition, err = p.definitionManager.Get(name)
		providers       []scope.ScopeProvider
		defaultValue    string
	)

	if err != nil {
		return err
	}

	defaultValue = definition.DefaultValue

	if len(definition.AllowedProviders) > 0 {
		providers = p.scopeProviderManager.Get(definition.AllowedProviders...)
	} else {
		providers = p.scopeProviderManager.Providers()
	}

	if len(providers) == 0 {
		return ErrProviderNotFound
	}

	//No need to store on database if the value is the default value
	if defaultValue == value {
		for _, provider := range providers {
			if err = provider.Delete(name); err != nil {
				return err
			}
		}
	} else {
		if definition.IsEncrypted && value != "" {
			value = p.encryption.Encrypt(value)
		}

		//If it's not default value and not stored in database, then create it. Otherwise, update it
		for _, provider := range providers {
			if err = provider.Set(name, value); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *provider) AddDefination(definition SettingDefinition) error {
	return p.definitionManager.Add(definition)
}

func (p *provider) RemoveDefination(name string) error {
	return p.definitionManager.Remove(name)
}

func (p *provider) AddScopeProvider(provider scope.ScopeProvider) error {
	return p.scopeProviderManager.Add(provider)
}

func (p *provider) RemoveScopeProvider(name string) error {
	return p.scopeProviderManager.Remove(name)
}
