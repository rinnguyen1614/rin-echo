package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
	"gorm.io/gorm"
)

func newDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gormx.OpenWithDb(gormx.DriverPostgresql, db)
	if err != nil {
		return gormDB, mock, err
	}
	return gormDB, mock, err
}
