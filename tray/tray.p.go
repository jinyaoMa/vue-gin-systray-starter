package tray

import (
	"app/tray/locales"
	"app/tray/menus"
	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

var tray *Tray

func (t *Tray) onReady() {
	if !t.isRunning {
		return
	}

	systray.SetIcon(icon)
	t.resetTooltipLanguage()

	t.menus.server.Init()
	t.menus.server.Watch(&menus.ServerListener{
		OnStart: func() (ok bool) {
			ok = t.server.Start(false, false)
			t.config.StartServer = true
			t.config.EnableSwag = false
			t.resetTooltipLanguage()
			return
		},
		OnStartSwag: func() (ok bool) {
			ok = t.server.Start(true, false)
			t.config.StartServer = true
			t.config.EnableSwag = true
			t.resetTooltipLanguage()
			return
		},
		OnStop: func() (ok bool) {
			ok = t.server.Stop()
			t.config.StartServer = false
			t.config.EnableSwag = false
			t.resetTooltipLanguage()
			return
		},
	})

	systray.AddSeparator()

	t.menus.language.Init()
	t.menus.language.Watch(&menus.LanguageListener{
		OnLanguageChange: func(lang locales.Lang) (ok bool) {
			if ok = locales.Set(lang); ok {
				t.config.Locale = lang
				t.resetTooltipLanguage()

				t.menus.server.ResetLanguage()
				t.menus.language.ResetLanguage()
				t.menus.quit.ResetLanguage()
			}
			return
		},
	})

	systray.AddSeparator()

	t.menus.quit.Init()
	t.menus.quit.Watch(&menus.QuitListener{
		OnQuit: func() {
			t.Stop()
		},
	})

	t.handleConfig()
	t.logger.Println("Tray is running!")
}

func (t *Tray) onExit() {
	t.menus.server.StopWatch()
	t.menus.language.StopWatch()
	t.menus.quit.StopWatch()
	t.server.Stop()

	t.logger.Println("Tray is exiting!")
}

func (t *Tray) handleConfig() {
	var signal struct{}

	if t.config.StartServer {
		if t.config.EnableSwag {
			t.menus.server.StartSwag.ClickedCh <- signal
		} else {
			t.menus.server.Start.ClickedCh <- signal
		}
	} else {
		t.menus.server.Stop.ClickedCh <- signal
	}

	switch t.config.Locale {
	case locales.En:
		t.menus.language.En.ClickedCh <- signal
	case locales.Zh:
		t.menus.language.Zh.ClickedCh <- signal
	}
}

func (t *Tray) resetTooltipLanguage() {
	var src = locales.Get()

	var cState string
	if t.config.StartServer {
		if t.config.EnableSwag {
			cState = src.Server.StartWithSwag.String()
		} else {
			cState = src.Server.Start.String()
		}
	} else {
		cState = src.Server.Stop.String()
	}

	var cLocale string
	switch t.config.Locale {
	case locales.En:
		cLocale = src.Language.En.String()
	case locales.Zh:
		cLocale = src.Language.Zh.String()
	}

	systray.SetTooltip(src.Tooltip.String(cState, cLocale))
}
