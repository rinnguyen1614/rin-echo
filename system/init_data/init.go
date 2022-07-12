package initdata

import (
	"rin-echo/system/adapters"
	"rin-echo/system/adapters/manager"
	"rin-echo/system/inject"

	"github.com/rinnguyen1614/rin-echo-core/log"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
)

func Init() {

	var (
		uow               = adapters.NewUnitOfWork(inject.GetDB())
		permissionManager = manager.NewPermissionManager(inject.GetCasbin())
		logger            = inject.GetLogger()
		config            = inject.GetConfig()
	)

	Migrate(uow.DB())

	if err := uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		if err := initResources(ux, config.App.Resources["resource"]); err != nil {
			return err
		}

		if err := initMenus(ux, config.App.Resources["menu"]); err != nil {
			return err
		}

		if err := initRoleAndUser(ux, permissionManager); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Log(logger, log.DebugLevel, err.Error())
	}

}
