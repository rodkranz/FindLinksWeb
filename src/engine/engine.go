package engine

import (
	"github.com/rodkranz/FindLinksWeb/src/interfaces"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type Engine struct {
	config  *interfaces.Configuration
	engines []interfaces.EngineInterface
}

func NewEngine(conf *interfaces.Configuration) *Engine {
	return &Engine{config: conf}
}

func (e *Engine) AddEngine(eng ...interfaces.EngineInterface) {
	e.engines = append(e.engines, eng...)
}

func (e *Engine) Run() {
	var wg sync.WaitGroup

	for _, eng := range e.engines {
		wg.Add(1)

		go func(eng interfaces.EngineInterface) {
			defer wg.Done()
			htmlData := e.downloadHTML(eng)
			list := e.parseHTML(htmlData, eng)
			eng.SetDataBundle(list)
		}(eng)
	}

	wg.Wait()
}

func (e *Engine) ShowResult() {
	fmt.Println("[+] Find Web Link V1 By rodlopes <dev.rodrigo.lopes@gmail.com>. \n")
	fmt.Println("[+] Searchers availables")

	for _, eng := range e.engines {
		fmt.Printf("[+] %v \t result(s) found in %v.\n", len(eng.GetData()), eng.GetTitle())
	}

	fmt.Println()
	for _, eng := range e.engines {
		if len(eng.GetData()) == 0 {
			continue
		}

		fmt.Printf("[+] %v \n", eng.GetTitle())
		for i, v := range eng.GetData() {
			fmt.Printf("[+] %v\t %v\n", i, v)
		}
		fmt.Println()
	}
}

func (e *Engine) downloadHTML(eng interfaces.EngineInterface) string {
	uri := e.makeUrl(eng)
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		log.Fatalf("Error to create a new request for [%s] -> %s\n", uri, err.Error())
	}

	req.Header.Set("User-Agent", e.config.UserAgent)
	res, err := client.Do(req)
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
	matched := eng.GetRegex().FindAllStringSubmatch(html, -1)

	var list []string
	for _, link := range matched {
		if !strings.HasPrefix(strings.ToLower(link[1]), "http") {
			continue
		}
		list = append(list, link[1])
	}

	return list
}
