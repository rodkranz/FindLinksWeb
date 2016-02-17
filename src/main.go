package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/rodkranz/FindLinksWeb/src/engine"
	"github.com/rodkranz/FindLinksWeb/src/interfaces"
	"github.com/rodkranz/FindLinksWeb/src/searchers"
)

var config = &interfaces.Configuration{}
var SearchersDefault = []string{"google", "yahoo"}

func init() {
	flag.StringVar(&config.Text, "t", "any news", "text that you looking for.")
	flag.StringVar(&config.Output, "out", "output.log", "Define that result will be save at file.")
	flag.StringVar(&config.UserAgent, "use-agent", "Go-http-client/1.2", "Change the browser that you will you in your search.")
	flag.IntVar(&config.Page, "p", 1, "Pagination number.")
	flag.Var(&config.Searchers, "searchers", "Web Search")

	flag.Parse()

	if len(config.Searchers) == 0 {
		config.Searchers = SearchersDefault
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	yahoo := searchers.NewYahoo()
	bing := searchers.NewBing()
	google := searchers.NewGoogle()
	duck := searchers.NewDuck()

	engine := engine.NewEngine(config)

	engine.AddEngine(bing, google, yahoo, duck)
	engine.Run()
	engine.ShowResult()

	os.Exit(0)
}
