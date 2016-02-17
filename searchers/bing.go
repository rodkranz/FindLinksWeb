package searchers

import (
	"regexp"
)
const bing_url   = "https://search.yahoo.com/search?p=[WORD]&&ei=UTF-8&b=[PAGE]1"
const bing_regex = "\" ac-algo ac-21th lh-15\" href=\"(.*?)\" target=\"_blank"

func NewBing() *Bing {
	return &Bing{}
}

type Bing struct {}

func (b *Bing) GetUrl() string {
	return bing_url
}

func (b *Bing) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(bing_regex)
}

