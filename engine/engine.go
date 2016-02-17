package engine

import (
	"github.com/rodkranz/FindLinksWeb/interfaces"

	"fmt"
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

type Engine struct {
	config     *interfaces.Configuration
	engines    []interfaces.EngineInterface
}

func NewEngine(conf *interfaces.Configuration) *Engine {
	return &Engine{config: conf}
}

func (e *Engine) AddEngine(eng interfaces.EngineInterface) {
	e.engines = append(e.engines, eng)
}


func (e *Engine) downloadHTML(eng interfaces.EngineInterface) string {
	uri         := eng.GetUrl();
	client      := &http.Client{}
	req, err    := http.NewRequest("GET", uri, nil)

	if err != nil {
		log.Fatalf("Error to create a new request for [%s] -> %s\n", uri, err.Error())
	}

	req.Header.Set("User-Agent", e.config.UserAgent)
	res, err  := client.Do(req)
	if err != nil {
		log.Fatalf("Error to read data of body -> %s\n", err.Error())
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	return string(data)
}



func (e *Engine) makeUrl(eng interfaces.EngineInterface) string {
	uri := eng.GetUrl()

	uri = strings.Replace(uri, "[WORD]", url.QueryEscape(e.config.Text), -1)
	uri = strings.Replace(uri, "[PAGE]", strconv.Itoa(e.config.Page), -1)

	return uri
}


func (e *Engine) parseHTML(html string, eng interfaces.EngineInterface) []string {
	matched := eng.GetRegex().FindAllStringSubmatch(html,-1)

	var list []string
	for _, link := range matched {
		list = append(list, link[1])
	}

	return list
}