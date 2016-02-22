package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"github.com/rodkranz/FindLinksWeb/src/engine"
	"github.com/rodkranz/FindLinksWeb/src/interfaces"
	"github.com/rodkranz/FindLinksWeb/src/searchers"
	"github.com/rodkranz/FindLinksWeb/src/output"
)

var config = &interfaces.Configuration{}
var SearchersDefault = []string{"yahoo", "bing", "google", "duck"}

func init() {
	flag.StringVar(&config.Text, "t", "any news", "text that you looking for.")
	flag.StringVar(&config.Output, "out", "output.log", "Def,ine that result will be save at file.")
	flag.StringVar(&config.UserAgent, "use-agent", "Go-http-client/1.2", "Change the browser that you will you in your search.")
	flag.IntVar(&config.Page, "p", 1, "Pagination number.")
	flag.Var(&config.Searchers, "searchers", "Web Search")

	flag.Parse()

	if len(config.Searchers) == 0 {
		config.Searchers = SearchersDefault
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	output.Banner()
	engine := engine.NewEngine(config)

	AddSearchers(engine)

	signalToContinue := make(chan bool)
	go engine.Run(signalToContinue)

	output.Searching()
	<-signalToContinue

	output.ClearLine()

	output.ShowEngines(engine)
	os.Exit(0)
}

func AddSearchers(e *engine.Engine) {
	for _, eng := range config.Searchers {
		switch strings.ToLower(eng) {
		case "yahoo":
			e.AddEngine(searchers.NewYahoo())
			break
		case "bing":
			e.AddEngine(searchers.NewBing())
			break
		case "google":
			e.AddEngine(searchers.NewGoogle())
			break
		case "duck":
			e.AddEngine(searchers.NewDuck())
			break
		}
	}
}