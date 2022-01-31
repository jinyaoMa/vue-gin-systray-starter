package locales

const En Lang = "en"

func init() {
	locales[En] = &Strings{
		Title:   "gin-systray-starter-v2",
		Tooltip: "Server: {1}\nLanguage: {2}",
		Quit:    "Quit",
		Server: Server{
			Title:         "Server",
			Start:         "Start",
			StartWithSwag: "Start with Swagger",
			Stop:          "Stop",
		},
		Language: Language{
			Title: "Language",
			En:    "English",
			Zh:    "中文",
		},
	}
}
