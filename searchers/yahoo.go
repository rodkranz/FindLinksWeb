package searchers

import (
	"regexp"
)

const yahoo_url   = "https://search.yahoo.com/search?p=[WORD]&&ei=UTF-8&b=[PAGE]1"
const yahoo_regex = "\" ac-algo ac-21th lh-15\" href=\"(.*?)\" target=\"_blank"

func NewYahoo() *Yahoo {
	return &Yahoo{}
}

type Yahoo struct{}

func (b *Yahoo) GetUrl() string {
	return yahoo_url
}

func (b *Yahoo) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(yahoo_regex)
}

