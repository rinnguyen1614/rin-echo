package gorm

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	//"gorm.io/driver/sqlite"
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

func Open(dialect string, dns string, options ...gorm.Option) (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	if dialect == DriverPostgresql {
		db, err = gorm.Open(postgres.Open(dns), options...)
	} else if dialect == DriverMysql {
		db, err = gorm.Open(mysql.Open(dns), options...)
	} else if dialect == DriverMssql {
		db, err = gorm.Open(sqlserver.Open(dns), options...)
	} else if dialect == DriverSqlite {
		//db, err = gorm.Open(sqlite.Open(dns), options...)
	} else {
		return nil, errors.New("database dialect is not supported")
	}
	if err != nil {
		return nil, err
	}

	return db, err
}

func OpenWithConfig(config Database) (*gorm.DB, error) {
	return Open(config.Driver, config.GetDNS(), &gorm.Config{
		CreateBatchSize:        config.BatchSize,
		PrepareStmt:            config.PrepareStmt,
		SkipDefaultTransaction: true,
	})
}
