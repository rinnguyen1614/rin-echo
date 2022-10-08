package inject

import (
	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
	"github.com/rinnguyen1614/rin-echo/internal/core/log"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	if service.db == nil {
		cfg := GetConfig()
		logger := GetLogger()
		db, err := gormx.OpenWithConfig(gormx.Database{
			Driver:    cfg.Database.Driver,
			DNS:       cfg.Database.URL,
			BatchSize: cfg.Database.BatchSize,
		})

		if err != nil {
			panic(err)
		}
		db.Logger = gormx.NewLogger(logger, gormx.LoggerConfig{LogLevel: log.WarnLevel})
		db.Use(&gormx.RinPlugin{})
		service.db = db
	}
	return service.db
}
