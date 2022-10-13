package gorm

import (
	"database/sql"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	// DriverMysql is a const value of mysql driver.
	DriverMysql = "mysql"
	// DriverSqlite is a const value of sqlite driver.
	DriverSqlite = "sqlite"
	// DriverPostgresql is a const value of postgresql driver.
	DriverPostgresql = "postgresql"
	// DriverMssql is a const value of mssql driver.
	DriverMssql = "mssql"
)

func Open(driver, dns string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector, err := GetDialector(driver, dns)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenWithDb(driver string, dbExisting *sql.DB, opts ...gorm.Option) (*gorm.DB, error) {
	dialector, err := GetDialectorWithDb(driver, dbExisting)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDialector(driver, dns string) (gorm.Dialector, error) {
	var (
		err       error
		dialector gorm.Dialector
	)

	switch driver {
	case DriverPostgresql:
		dialector = postgres.Open(dns)
	case DriverMysql:
		dialector = mysql.Open(dns)
	case DriverMssql:
		dialector = sqlserver.Open(dns)
	case DriverSqlite:
		dialector = sqlite.Open(dns)
	default:
		err = errors.New("database dialect is not supported")
	}

	if err != nil {
		return nil, err
	}

	return dialector, err
}

func GetDialectorWithDb(driver string, db *sql.DB) (gorm.Dialector, error) {
	var (
		err       error
		dialector gorm.Dialector
	)

	switch driver {
	case DriverPostgresql:
		dialector = postgres.New(postgres.Config{Conn: db})
	case DriverMysql:
		dialector = mysql.New(mysql.Config{Conn: db})
	case DriverMssql:
		dialector = sqlserver.New(sqlserver.Config{Conn: db})
	default:
		err = errors.New("database dialect is not supported")
	}

	if err != nil {
		return nil, err
	}

	return dialector, err
}
