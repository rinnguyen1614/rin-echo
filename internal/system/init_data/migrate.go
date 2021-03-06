package initdata

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202104050000",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.AuditLog{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.AuditLog{})
			},
		},
		{
			ID: "202104050001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.SecurityLog{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.SecurityLog{})
			},
		},
		{
			ID: "202104050002",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Setting{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Setting{})
			},
		},
		{
			ID: "202104050100",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Resource{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Resource{})
			},
		},
		{
			ID: "202104050101",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Menu{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Menu{})
			},
		},
		{
			ID: "202104050112",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.User{})
			},
		},

		{
			ID: "202104050103",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Role{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Role{})
			},
		},
		{
			ID: "202104050104",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.UserRole{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.UserRole{})
			},
		},
		{
			ID: "202104050105",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Permission{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Permission{})
			},
		},
		{
			ID: "202104050106",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Address{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Address{})
			},
		},
		{
			ID: "202104050107",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.AddressLocation{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.AddressLocation{})
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Printf("Could not migrate: %v", err)
	} else {
		log.Printf("Migration did run successfully")
	}

}
