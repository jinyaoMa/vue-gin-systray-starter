package config

import (
	"app/tray/locales"
	"os"
	"path/filepath"
)

type Configs struct {
	IniPath string `ini:"-"`
	IsDev   bool
	Logger  *Logger
	Tray    *Tray
	Server  *Server
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
	}
}
