package searchers

import "github.com/rodkranz/FindLinksWeb/interfaces"

func NewYahoo() *interfaces.Gas {
	return &interfaces.Gas{Title: "Yahoo",
		Url:   `https://search.yahoo.com/search?p=[WORD]&&ei=UTF-8&b=[PAGE]1`,
		Regex: `" ac-algo ac-21th lh-15" href="(.*?)" target="_blank`}
}