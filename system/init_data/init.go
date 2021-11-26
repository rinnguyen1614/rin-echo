package initdata

import (
	"rin-echo/system/adapters"
	"rin-echo/system/inject"
)

func Init() {

	var (
		uow = adapters.NewUnitOfWork(inject.GetDB())
	)

	Migrate(uow.DB())
}
