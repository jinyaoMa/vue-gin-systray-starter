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

func Init(logger *log.Logger, config *config.Database) {
	database = &Database{
		logger: logger,
		config: config,
	}
	database.initDB()
}
