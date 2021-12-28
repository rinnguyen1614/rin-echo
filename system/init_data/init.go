package initdata

import (
	"rin-echo/common/log"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/adapters"
	"rin-echo/system/adapters/manager"
	"rin-echo/system/inject"
)

func Init() {

	var (
		uow               = adapters.NewUnitOfWork(inject.GetDB())
		permissionManager = manager.NewPermissionManager(inject.GetCasbin())
		logger            = inject.GetLogger()
	)

	Migrate(uow.DB())

	if err := uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		if err := initRoleAndUser(ux, permissionManager); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Log(logger, log.DebugLevel, err.Error())
	}

}
