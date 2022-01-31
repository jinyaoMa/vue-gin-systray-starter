package main

import (
	"app/config"
	"app/database"
	"app/server"
	"app/tray"
	"app/utils/logger"
	"flag"
)

var hasTray = flag.Int("t", 1, "set to enable system tray")

func main() {
	flag.Parse()

	configs := config.LoadConfigs()

	logger := logger.New(configs.Logger, configs.IsDev)

	database.Connect(logger.Database, configs.Database)

	server := server.New(logger.Server, configs.Server, configs.IsDev)

	if *hasTray == 1 {
		tray := tray.New(logger.Tray, configs.Tray, server)
		tray.Run()
	} else {
		server.Start(true, true)
	}
}
