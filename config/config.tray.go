package config

import "app/tray/locales"

type Tray struct {
	Locale      locales.Lang `comment:"Locale options: en, zh"`
	StartServer bool
	EnableSwag  bool
}
