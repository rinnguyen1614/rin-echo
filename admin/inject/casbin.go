package inject

import (
	casbinx "rin-echo/common/casbin"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func GetCasbin() *casbin.Enforcer {
	if service.enforcer == nil {
		cfg := GetConfig()
		logger := GetLogger()
		e, err := casbin.NewEnforcer(cfg.Casbin.ModelPath)
		if err != nil {
			panic(err)
		}
		e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch2)
		e.SetLogger(casbinx.NewLogger(logger, 0))
		service.enforcer = e
	}
	return service.enforcer
}
