package adapter

import (
	"errors"
	"fmt"
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/core/cache"
	model "github.com/rinnguyen1614/rin-echo/internal/core/setting/model"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"gorm.io/gorm"
)

var (
	Expiration         = time.Second * 0
	ErrSettingExists   = errors.New("setting existes in system")
	ErrSettingNotFound = errors.New("setting not found")
)

type settingCacheItem struct {
	Key   string
	Value string
}

type DefaultAdapter struct {
	uow   iuow.UnitOfWork
	cache cache.Cache
	ctx   core.Context
}

func NewDefaultAdapter(uow iuow.UnitOfWork, cache cache.Cache) Adapter {
	return &DefaultAdapter{
		uow:   uow,
		cache: cache,
	}
}

func (a *DefaultAdapter) WithContext(ctx core.Context) Adapter {
	if ctx == nil {
		panic("nil context")
	}
	a2 := new(DefaultAdapter)
	*a2 = *a
	a2.ctx = ctx
	a2.uow = a.uow.WithContext(ctx)

	return a2
}

// Get matched record or initialize a new instance with given name, providerName, providerKey
func (a *DefaultAdapter) GetOrInit(name, providerName, providerKey string) string {
	return a.getOrInitCacheItem(name, providerName, providerKey)
}

func (a *DefaultAdapter) GetMulti(names []string, providerName, providerKey string) map[string]string {
	return a.getMultiCacheItem(names, providerName, providerKey)
}

func (a *DefaultAdapter) GetAll(providerName, providerKey string) map[string]string {
	var (
		settingValues model.Settings
		values        = make(map[string]string)
	)

	// get from store
	a.filter(providerName, providerKey).Find(&settingValues)

	if settingValues != nil {
		values = settingValues.ToValuesByName()
	}

	return values
}

func (a *DefaultAdapter) Set(name, value, providerName, providerKey string) error {
	return a.uow.TransactionUnitOfWork(func(uow iuow.UnitOfWork) error {
		var (
			tx           = a.filter(providerName, providerKey, name)
			settingValue model.Setting
			update       bool
			err          error
		)
		// check exists in database
		if err = tx.First(&settingValue).Error; err == nil {
			update = true
		}

		if !update {
			err = a.uow.DB().Create(model.Setting{
				Name:         name,
				Value:        value,
				ProviderKey:  providerKey,
				ProviderName: providerName,
			}).Error

		} else if settingValue.Value != value {
			// if value is changed, then update record in database
			err = tx.Updates(model.Setting{Value: value}).Error
		}

		if err != nil {
			return err
		}

		// update item in cache by name & providerKey & providerName
		return a.setCacheItem(name, value, providerName, providerKey)
	})
}

func (a *DefaultAdapter) Delete(name, providerName, providerKey string) error {

	return a.uow.TransactionUnitOfWork(func(uow iuow.UnitOfWork) error {
		var (
			tx           = a.filter(providerName, providerKey, name)
			settingValue model.Setting
		)
		// check exists in database
		if err := tx.First(&settingValue).Error; err != nil {
			return err
		}

		if err := tx.Delete(&settingValue).Error; err != nil {
			return err
		}

		// update item in cache by name & providerKey & providerName
		return a.deleteCacheItem(name, providerName, providerKey)
	})

}

func (a *DefaultAdapter) FilterName(tx *gorm.DB, names ...string) *gorm.DB {
	return tx.Where("name IN ?", names)
}

func (a *DefaultAdapter) FilterProviderName(tx *gorm.DB, providerNames ...string) *gorm.DB {
	return tx.Where("provider_name IN ?", providerNames)
}

func (a *DefaultAdapter) FilterProviderKey(tx *gorm.DB, providerKeys ...string) *gorm.DB {
	return tx.Where("provider_key IN ?", providerKeys)
}

func (a *DefaultAdapter) filter(providerName, providerKey string, names ...string) *gorm.DB {
	var tx = a.uow.DB()

	if len(names) != 0 {
		tx = a.FilterName(tx, names...)
	}

	return a.FilterProviderKey(
		a.FilterProviderName(
			tx,
			providerName),
		providerKey,
	)
}

func (a *DefaultAdapter) getOrInitCacheItem(name, providerName, providerKey string) string {
	var (
		settingValue model.Setting
		key          = getKey(name, providerName, providerKey)
		item, err    = cache.GetOrCreate(a.cache, a.ctx, key, func(key string) (value interface{}) {
			a.filter(providerName, providerKey, name).First(&settingValue)

			return settingCacheItem{
				Key:   settingValue.Name,
				Value: settingValue.Value,
			}
		})
	)

	if err != nil {
		return ""
	}
	return item.(settingCacheItem).Value
}

func (a *DefaultAdapter) getMultiCacheItem(names []string, providerName, providerKey string) map[string]string {
	var (
		values             = make(map[string]string)
		settingValuesByKey = make(map[string]*model.Setting)
		keys               []string
		settingValues      model.Settings
	)

	if len(names) == 0 {
		return values
	}

	if len(names) == 1 {
		values[names[0]] = a.getOrInitCacheItem(names[0], providerName, providerKey)
		return values
	}

	// get from store
	a.filter(providerName, providerKey, names...).Find(&settingValues)
	settingValuesByName := settingValues.ToMapByName()
	for _, name := range names {
		key := getKey(name, providerName, providerKey)
		keys = append(keys, key)
		settingValuesByKey[key] = settingValuesByName[name]
	}

	cacheItems, _ := cache.GetOrCreateMulti(a.cache, a.ctx, keys, func(key string) (value interface{}) {
		settingValue := settingValuesByKey[key]
		return settingCacheItem{
			Key:   settingValue.Name,
			Value: settingValue.Value,
		}
	})

	for _, item := range cacheItems {
		values[item.(settingCacheItem).Key] = item.(settingCacheItem).Value
	}

	return values
}

func (a *DefaultAdapter) setCacheItem(name, value, providerName, providerKey string) error {
	return a.cache.Set(a.ctx, getKey(name, providerName, providerKey), settingCacheItem{Key: name, Value: value}, Expiration)
}

func (a *DefaultAdapter) deleteCacheItem(name, providerName, providerKey string) error {
	return a.cache.Delete(a.ctx, getKey(name, providerName, providerKey))
}

func getKey(name, providerName, providerKey string) string {
	return fmt.Sprintf("%s_%s_%s", name, providerName, providerKey)
}
