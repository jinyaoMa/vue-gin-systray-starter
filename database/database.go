package database

import (
	"app/config"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	logger *log.Logger
	config *config.Database
}

func Connect(logger *log.Logger, config *config.Database) *gorm.DB {
	database = &Database{
		logger: logger,
		config: config,
	}
	DB = database.initDB()
	return DB
}
