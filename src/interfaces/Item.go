package interfaces

import (
	"regexp"
)

type Gas struct {
	Title string
	Regex string
	Url   string
	Data  []string
}

func (i *Gas) SetData(url string) {
	i.Data = append(i.Data, url)
}

func (i *Gas) SetDataBundle(url []string) {
	for _, u := range url {
		i.Data = append(i.Data, u)
	}
}

func (i *Gas) GetData()   []string {
	return i.Data
}

func (i *Gas) GetTitle() string {
	return i.Title
}

func (i *Gas) GetUrl() string {
	return i.Url
}

func (i *Gas) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(i.Regex)
}

