package interfaces

import "regexp"

type EngineInterface interface {
	GetUrl()    string
	GetRegex()  *regexp.Regexp
}