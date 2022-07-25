package setting

import (
	"context"
	"fmt"
	"testing"

	"github.com/rinnguyen1614/rin-echo/internal/core/cache/memory"
	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting/adapter"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting/scope"
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/stretchr/testify/assert"
)

type session struct {
}

func (s session) UserID() uint {
	return currentUser
}

type setting struct {
	Name         string
	Value        string
	ProviderName string
	ProviderKey  string
}

var (
	ctx                      = core.NewContext(context.Background(), session{})
	currentUser         uint = 1
	gormAdapter         adapter.Adapter
	userScopeProvider   *scope.UserScopeProvider
	globalScopeProvider *scope.GlobalScopeProvider
	settingProvider     Provider
	encryption          = utils.NewEncryption("test")
	settingDefinitions  []SettingDefinition
)

func init() {
	initAdapter()
	initSettingDefinations()
	initScopeProviders()
	settingProvider = New(settingDefinitions, []scope.ScopeProvider{
		userScopeProvider,
		globalScopeProvider,
	}, encryption)
}

func initAdapter() {
	db, err := gormx.OpenWithConfig(gormx.Database{
		Driver: "postgresql",
		DNS:    "host=localhost user=admin password=anhnguyen0809 dbname=rin_admin port=5432 sslmode=disable",
	})

	if err != nil {
		panic(err)
	}

	gormAdapter = adapter.NewDefaultAdapter(uow.NewUnitOfWork(db), memory.NewMemoryCache(0))
}

func initSettingDefinations() {
	settingDefinitions = []SettingDefinition{
		{
			Name:             "app:name",
			DefaultValue:     "rin-echo Application",
			AllowedProviders: []string{scope.GlobalSettingProviderName},
		},
		{
			Name:        "email:username",
			DisplayName: "Email Username",
		},
		{
			Name:        "email:password",
			DisplayName: "email password",
			IsEncrypted: true,
		},
		{
			Name:             "app:theme",
			DefaultValue:     "dark",
			AllowedProviders: []string{scope.UserSettingProviderName},
		},
	}
}

func initScopeProviders() {
	userScopeProvider = scope.NewUserScopeProviderWithContext(gormAdapter, ctx)
	globalScopeProvider = scope.NewGlobalScopeProvider(gormAdapter)
}

func TestProvider_Get(t *testing.T) {
	value, err := settingProvider.Get("email:username")
	assert.Nil(t, err)

	err = settingProvider.Set("email:username", "admin@rin-echo.com")
	assert.Nil(t, err)
	value, err = settingProvider.Get("email:username")
	assert.Nil(t, err)
	assert.Equal(t, "admin@rin-echo.com", value)
}

func TestProvider_GetMulti(t *testing.T) {
	names := []string{"email:username", "email:password", "app:theme"}
	values, err := settingProvider.GetMulti(names)
	assert.Nil(t, err)

	assert.Equal(t, len(names), len(values))
}

func TestProvider_DefaultValue(t *testing.T) {
	name := "app:name"
	defaultValue := "rin-echo Application"
	value, err := settingProvider.Get(name)
	assert.Nil(t, err)
	if value != defaultValue {
		err = settingProvider.Set(name, defaultValue) // it will be removed from store
		assert.Nil(t, err)

		// check it's removed from store
		valueOfGlobal := globalScopeProvider.GetOrInit(name)
		assert.Equal(t, "", valueOfGlobal)

		// check value of name is default value
		value, err = settingProvider.Get(name)
		assert.Equal(t, defaultValue, value)
	}
}

func TestProvider_Encryption(t *testing.T) {
	name := "email:password"
	value, err := settingProvider.Get(name)
	assert.Nil(t, err)

	plain_value := "123456"
	err = settingProvider.Set(name, plain_value)
	assert.Nil(t, err)
	value, err = settingProvider.Get(name)
	assert.Nil(t, err)
	assert.Equal(t, plain_value, value)
	fmt.Print(value)
}

func TestProvider_For_CurrentUser(t *testing.T) {
	name := "app:theme"
	value, err := settingProvider.Get(name)
	assert.Nil(t, err)

	if value == "dark" {
		err = settingProvider.Set(name, "light")
		assert.Nil(t, err)
		value, err = settingProvider.Get(name)
		assert.Nil(t, err)
	}

	valueOfUser := userScopeProvider.GetOrInit(name)
	assert.Equal(t, valueOfUser, value)

	valueOfGlobal := globalScopeProvider.GetOrInit(name)
	assert.NotEqual(t, valueOfGlobal, value)
	assert.Equal(t, "", valueOfGlobal)
}

func TestProvider_For_Global(t *testing.T) {
	name := "app:name"
	defaultValue := "rin-echo Application"
	value, err := settingProvider.Get(name)
	assert.Nil(t, err)

	if value == defaultValue {
		err = settingProvider.Set(name, "new value")
		assert.Nil(t, err)
		value, err = settingProvider.Get(name)
		assert.Nil(t, err)
	}

	valueOfGlobal := globalScopeProvider.GetOrInit(name)
	assert.Equal(t, valueOfGlobal, value)

	valueOfUser := userScopeProvider.GetOrInit(name)
	assert.NotEqual(t, valueOfUser, value)
	assert.Equal(t, "", valueOfUser)
}
