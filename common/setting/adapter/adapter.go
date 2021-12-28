package adapter

import "rin-echo/common"

type Adapter interface {
	WithContext(ctx common.Context) Adapter

	GetOrInit(name, providerName, providerKey string) string

	GetMulti(names []string, providerName, providerKey string) map[string]string

	GetAll(providerName, providerKey string) map[string]string

	//If it's not stored in database, then create it. Otherwise, update it
	Set(name, value, providerName, providerKey string) error
	// Delete value by name, providerName and providerKey in store.
	Delete(name, providerName, providerKey string) error
}
