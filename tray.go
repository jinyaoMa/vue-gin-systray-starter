package main

import (
	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed public/favicon.ico
var icon []byte

var tray *Tray

type Tray struct {
	isRunning bool
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
	systray.Run(t.onTrayReady, t.onTrayExit)
}

func (t *Tray) Stop() {
	if !t.isRunning {
		return
	}
	t.isRunning = false
	systray.Quit()
}

func (t *Tray) onTrayReady() {
	systray.SetIcon(icon)
}

func (t *Tray) onTrayExit() {

}
