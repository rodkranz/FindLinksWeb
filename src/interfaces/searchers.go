package interfaces

import "fmt"

type Searchers []string

func (i *Searchers) String() string {
	return fmt.Sprintf("%v", &i);
}

func (i *Searchers) Set(value string) error {
	*i = append(*i, value)
	return nil
}
