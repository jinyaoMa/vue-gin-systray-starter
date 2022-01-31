package locales

const Zh Lang = "zh"

func init() {
	locales[Zh] = &Strings{
		Title:   "我的 APP",
		Tooltip: "服务器: {1}\n语言: {2}",
		Quit:    "退出",
		Server: Server{
			Title:         "服务器",
			Start:         "开启",
			StartWithSwag: "开启（有 Swagger）",
			Stop:          "关闭",
		},
		Language: Language{
			Title: "语言",
			En:    "English",
			Zh:    "中文",
		},
	}
}
