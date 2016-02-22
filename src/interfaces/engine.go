package interfaces

import "regexp"

type EngineInterface interface {
	SetData(url string)
	SetDataBundle(url []string)
	GetTitle() string
	GetData() []string
	GetUrl() string
	GetRegex() *regexp.Regexp
	GetWord() string
}
