package tray

import (
	"app/config"
	"app/server"
	"app/tray/menus"
	"log"

	"github.com/getlantern/systray"
)

type Tray struct {
	isRunning bool
	logger    *log.Logger
	config    *config.Tray
	menus     TrayMenus
	server    *server.Server
}

type TrayMenus struct {
	server   *menus.Server
	language *menus.Language
	quit     *menus.Quit
}

func GetInstance() (*Tray, bool) {
	return tray, tray != nil
}

func New(logger *log.Logger, config *config.Tray, server *server.Server) *Tray {
	if tray != nil {
		if !tray.isRunning {
			tray.logger = logger
			tray.config = config
			tray.server = server
		}
		return tray
	}

	tray = &Tray{
		logger: logger,
		config: config,
		server: server,
	}
	return tray
}

func (t *Tray) Run() {
	if tray == nil {
		tray = t
	} else {
		t = tray
	}

	if t.isRunning {
		return
	}
	t.isRunning = true

	t.menus.server = &menus.Server{}
	t.menus.language = &menus.Language{}
	t.menus.quit = &menus.Quit{}

	systray.Run(t.onReady, t.onExit)
}

func (t *Tray) Stop() {
	if !t.isRunning {
		return
	}

	t.isRunning = false

	systray.Quit()
}
