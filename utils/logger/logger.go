package logger

import (
	"app/config"
	"io"
	"log"
	"os"
)

const (
	LoggerLabelTray     = "TRAY: "
	LoggerLabelServer   = "SERVER: "
	LoggerLabelDatabase = "DATABASE: "
)

var logger *Logger

type Logger struct {
	Tray     *log.Logger
	Server   *log.Logger
	Database *log.Logger
}

func GetInstance() (*Logger, bool) {
	return logger, logger != nil
}

func New(config *config.Logger, isDev bool) *Logger {
	if logger != nil {
		return logger
	}

	logger = &Logger{}
	if isDev {
		logger.Tray = newLog(os.Stdout, LoggerLabelTray)
		logger.Server = newLog(os.Stdout, LoggerLabelServer)
		logger.Database = newLog(os.Stdout, LoggerLabelDatabase)
	} else {
		logTray := getLogFile(config.TrayPath)
		logServer := getLogFile(config.ServerPath)
		logDatabase := getLogFile(config.DatabasePath)
		logger.Tray = newLog(io.MultiWriter(logTray, os.Stdout), LoggerLabelTray)
		logger.Server = newLog(io.MultiWriter(logServer, os.Stdout), LoggerLabelServer)
		logger.Database = newLog(io.MultiWriter(logDatabase, os.Stdout), LoggerLabelDatabase)
	}
	return logger
}
