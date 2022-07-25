package inject

import (
	"time"

	casbinx "github.com/rinnguyen1614/rin-echo/internal/core/casbin"

	"github.com/casbin/casbin/v2"
	gormadapder "github.com/casbin/gorm-adapter/v3"
)

func GetCasbin() *casbin.SyncedEnforcer {
	if service.enforcer == nil {
		cfg := GetConfig()
		logger := GetLogger()
		db := GetDB()

		a, err := gormadapder.NewAdapterByDBUseTableName(db, "", "casbin_rules")
		if err != nil {
			panic(err)
		}

		e, err := casbin.NewSyncedEnforcer(cfg.Casbin.ModelPath, a)
		if err != nil {
			panic(err)
		}

		err = e.LoadPolicy()
		if err != nil {
			panic(err)
		}

		//e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch2)
		e.SetLogger(casbinx.NewLogger(logger, 0))

		if cfg.Casbin.AutoLoad {
			e.StartAutoLoadPolicy(time.Duration(cfg.Casbin.AutoLoadInternal) * time.Second)
		}
		e.LoadPolicy()

		service.enforcer = e
	}
	return service.enforcer
}
