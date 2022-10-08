package initdata

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/manager"

	"log"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters"
	"github.com/rinnguyen1614/rin-echo/internal/system/inject"
)

func Init() {
	var (
		uow               = adapters.NewUnitOfWork(inject.GetDB())
		permissionManager = manager.NewPermissionManager(inject.GetCasbin())
		config            = inject.GetConfig()
	)

	Migrate(config.Database.MigrationURL, config.Database.URL)

	if config.Database.InitData {
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
			log.Fatal(err.Error())
		}
	}
}
