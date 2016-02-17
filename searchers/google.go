package searchers

import "github.com/rodkranz/FindLinksWeb/interfaces"

func NewGoogle() *interfaces.Gas {
	return &interfaces.Gas{Title: "Google",
		Url:   `https://www.google.com.br/search?q=[WORD]&q=[WORD]&start=[PAGE]0`,
		Regex: `"><a href="/url\?q=(.*?)&amp;sa=U&amp;`}
}