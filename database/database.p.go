package database

import (
	"app/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var database *Database

func (d *Database) initDB() {
	switch d.config.Driver {
	case config.DriverSqlite:
		d.initSqlite()
	case config.DriverMysql:
		d.initMysql()
	case config.DriverPostgres:
		d.initPostgres()
	}
}

func (d *Database) initSqlite() {
	var err error
	DB, err = gorm.Open(sqlite.Open(d.config.Database), &gorm.Config{
		Logger: d.getSqlLogger(),
	})
	if err != nil {
		d.logger.Fatalf("database (sqlite) connect error %v\n", err)
	}
	if DB.Error != nil {
		d.logger.Fatalf("database (sqlite) error %v\n", DB.Error)
	}
}

func (d *Database) initMysql() {
	var err error
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&%s",
		d.config.User,
		d.config.Password,
		d.config.Host,
		d.config.Port,
		d.config.Database,
		d.config.Tail)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: d.getSqlLogger(),
	})
	if err != nil {
		d.logger.Fatalf("database (mysql) connect error %v\n", err)
	}
	if DB.Error != nil {
		d.logger.Fatalf("database (mysql) error %v\n", DB.Error)
	}
}

func (d *Database) initPostgres() {
	var err error
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
		d.config.Host,
		d.config.User,
		d.config.Password,
		d.config.Database,
		d.config.Port,
		d.config.Tail)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: d.getSqlLogger(),
	})
	if err != nil {
		d.logger.Fatalf("database (postgres) connect error %v\n", err)
	}
	if DB.Error != nil {
		d.logger.Fatalf("database (postgres) error %v\n", DB.Error)
	}
}

func (d *Database) getSqlLogger() gormLogger.Interface {
	return gormLogger.New(
		d.logger,
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}
