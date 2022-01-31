package locales

import (
	"fmt"
	"strings"
)

type Strings struct {
	Title    Text
	Tooltip  Text
	Quit     Text
	Server   Server
	Language Language
}

type Server struct {
	Title         Text
	Start         Text
	StartWithSwag Text
	Stop          Text
}

type Language struct {
	Title Text
	En    Text
	Zh    Text
}

type Lang string
type Text string

func (t *Text) String(params ...string) string {
	text := string(*t)
	for i, param := range params {
		text = strings.Replace(text, fmt.Sprintf("{%d}", i+1), param, 1)
	}
	return text
}

var (
	locale  Lang              = En
	locales map[Lang]*Strings = make(map[Lang]*Strings, 2)
)

func Get() Strings {
	if strings, ok := locales[locale]; ok {
		return *strings
	}
	return Strings{}
}

func Set(lang Lang) (ok bool) {
	if _, ok = locales[lang]; ok {
		locale = lang
	}
	return
}
