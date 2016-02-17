package interfaces

import "regexp"

type EngineInterface interface {
	GetTitle()  string
	SetData(url string)
	SetDataBundle(url []string)
	GetData()   []string
	GetUrl()    string
	GetRegex()  *regexp.Regexp
}
