package searchers

import "github.com/rodkranz/FindLinksWeb/interfaces"

func NewBing() *interfaces.Gas {
	return &interfaces.Gas{Title: "Bing",
		Url:   "http://www.bing.com/search?q=[WORD]&first=[PAGE]1",
		Regex: `<li class="b_algo"><h2><a href="(.*?)" h="ID=SERP,[0-9|.]+">`}
}