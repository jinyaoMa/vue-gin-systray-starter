package database

import (
	"app/config"
	"errors"
	"log"

	"gorm.io/gorm"
)

type Database struct {
	logger *log.Logger
	config *config.Database
	DB     *gorm.DB
}

func GetInstance() (*Database, bool) {
	return database, database != nil
}

func Connect(logger *log.Logger, config *config.Database, models ...interface{}) {
	if database != nil {
		if database.DB.Error != nil && errors.Is(database.DB.Error, gorm.ErrInvalidDB) {
			database.logger = logger
			database.config = config
			database.initDriver(models...)
		}
	}

	database = &Database{
		logger: logger,
		config: config,
	}
	database.initDriver(models...)
}
