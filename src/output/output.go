package output

import (
	"github.com/fatih/color"

	"fmt"
	"strings"
	"github.com/rodkranz/FindLinksWeb/src/engine"
)

var (
	plus		= color.New(color.FgCyan).SprintFunc()
	owner 		= color.New(color.FgYellow).SprintFunc()
	number 		= color.New(color.FgGreen).SprintFunc()
	searchers	= color.New(color.FgBlue).SprintFunc()
	links		= color.New(color.FgBlue).SprintFunc()
)

func Banner() {
	fmt.Printf("%s Find Web Link v%s By %s [%s].\n", plus("+"), number("1.0.0"), owner("Rodrigo Lopes"), owner("dev.rodrigo.lopes@gmail.com"))
	fmt.Println()
}

func Searching() {
	fmt.Printf("%s We are searching, please wait...", plus("+"))
}

func ShowEngines(eng *engine.Engine){

	fmt.Printf("%s Searchers availables\n", plus("+"))
	for _, e := range eng.GetEngines() {
		fmt.Printf("%s %v \t result(s) found in %v.\n", plus("+"), number(len(e.GetData())), searchers(e.GetTitle()))
	}

	fmt.Println()
	for _, e := range eng.GetEngines() {
		if len(e.GetData()) == 0 {
			continue
		}
		ShowResult(e.GetTitle(),  eng.GetWord(), e.GetData())
	}
}


func ShowResult(title, word string, data []string) {
	fmt.Printf("%s %v: \n", plus("+"), searchers(title))
	for i, v := range data {
		fmt.Printf("%s %v\t %v\n", plus("+"), number((i+1)), GetRange(v, word))
	}
	fmt.Println()
}

func GetRange(line, word string) string {
	index := strings.Index(line, word);
	if index == -1 {
		return line
	}

	return fmt.Sprintf("%s%s%s", line[0:index], links(word), line[index+len(word):])
}

func ClearLine() {
	fmt.Printf("\r%s\r", strings.Repeat(" ", 100))
}