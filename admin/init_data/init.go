package initdata

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/inject"
	"rin-echo/common/log"
	iuow "rin-echo/common/uow/interfaces"
)

func Init() {

	var (
		uow    = adapters.NewUnitOfWork(inject.GetDB())
		rbac   = adapters.NewRBACCasbin(inject.GetCasbin())
		config = inject.GetConfig()
		logger = inject.GetLogger()
	)

	Migrate(uow.DB())

	err := uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		if err := initResources(ux, config.App.Resources["resource"]); err != nil {
			return err
		}

		if err := initRoleAndUser(ux, &rbac); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Log(logger, log.DebugLevel, err.Error())
	}
}
