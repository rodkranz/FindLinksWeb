package main

import (
	"flag"
	"github.com/rodkranz/FindLinksWeb/interfaces"
	"github.com/rodkranz/FindLinksWeb/searchers"
	"github.com/rodkranz/FindLinksWeb/engine"
	"os"
)

var config 			 = &interfaces.Configuration{}
var SearchersDefault = []string{"google","yahoo"}

func init() {
	flag.StringVar(&config.Text, 	    "t", 		 "any news", 	        "text that you looking for.")
	flag.StringVar(&config.Output, 	    "out", 		 "output.log", 	        "Define that result will be save at file.")
	flag.StringVar(&config.UserAgent,   "use-agent", "Go-http-client/1.1",  "Change the browser that you will you in your search.")
	flag.IntVar(&config.Page, 		    "p", 		 1,				        "Pagination number.")
	flag.Var(&config.Searchers, 	    "searchers", "Web Search")

	flag.Parse()

	if len(config.Searchers) == 0 {
		config.Searchers = SearchersDefault;
	}
}

func main() {
	bing 	:= searchers.NewBing()
	engine 	:= engine.NewEngine(config)

	engine.AddEngine(bing)
	engine.Run()

	os.Exit(0)
}

