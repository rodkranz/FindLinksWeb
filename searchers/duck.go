package searchers

import "github.com/rodkranz/FindLinksWeb/interfaces"

func NewDuck() *interfaces.Gas {
	return &interfaces.Gas{Title: "Duck Duck",
		Url:   `https://duckduckgo.com/html/?q=[WORD]`,
		Regex: `<a rel=\"nofollow\" href=\"(.*?)\">`}
}