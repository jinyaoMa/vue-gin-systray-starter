package config

import (
	"app/tray/locales"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Configs struct {
	IniPath  string `ini:"-"`
	IsDev    bool
	Logger   *Logger
	Tray     *Tray
	Server   *Server
	Database *Database
}

func DefaultConfigs() *Configs {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return &Configs{
		IniPath: filepath.Join(filepath.Dir(path), "app.ini"),
		IsDev:   true,
		Logger: &Logger{
			TrayPath:     filepath.Join(filepath.Dir(path), "tray.log"),
			ServerPath:   filepath.Join(filepath.Dir(path), "server.log"),
			DatabasePath: filepath.Join(filepath.Dir(path), "database.log"),
		},
		Tray: &Tray{
			Locale:      locales.Zh,
			StartServer: true,
			EnableSwag:  true,
		},
		Server: &Server{
			Port:    8080,
			PortTls: 8443,
			CertDir: filepath.Dir(path),
		},
		Database: &Database{
			Driver:   DriverSqlite,
			Host:     "",
			Port:     0,
			User:     "",
			Password: "",
			Database: filepath.Join(filepath.Dir(path), "app.db"),
			Tail:     "",
		},
	}
}

func LoadConfigs() *Configs {
	config := DefaultConfigs()

	_, err := os.Stat(config.IniPath)
	if err != nil {
		if os.IsNotExist(err) {
			iniFile := ini.Empty()
			err := ini.ReflectFrom(iniFile, config)
			if err != nil {
				panic("Failed to create ini: " + config.IniPath)
			}
			iniFile.SaveTo(config.IniPath)
			return config
		}
		panic("Error when loading ini: " + err.Error())
	}

	iniFile, err := ini.Load(config.IniPath)
	if err != nil {
		panic("Failed to load ini: " + config.IniPath)
	}
	err = iniFile.MapTo(config)
	if err != nil {
		panic("Failed to map ini: " + config.IniPath)
	}
	return config
}
