package setting

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
)

func Get[T any](provider Provider, name string) (T, error) {
	var value T
	if provider == nil {

	}
	v, err := provider.Get(name)
	if err != nil {
		return value, nil
	}

	return utils.Parse[T](v)
}

func MustGet[T any](provider Provider, name string) T {
	v, err := Get[T](provider, name)
	if err != nil {
		panic(err)
	}
	return v
}

func Set[T any](provider Provider, value T, name string) error {
	return provider.Set(utils.ToString(value), name)
}

func MustSet[T any](provider Provider, value T, name string) {
	err := Set(provider, value, name)
	if err != nil {
		panic(err)
	}
}
