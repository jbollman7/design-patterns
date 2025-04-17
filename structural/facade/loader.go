package main

import "fmt"

type Loader struct {
	load map[string]string
}

func newLoader(input map[string]string) *Loader {
	return &Loader{
		load: input,
	}
}

func (l *Loader) loadData(dataSource map[string]string) error {
	if fmt.Sprint(l.load) == fmt.Sprint((dataSource)) {
		return fmt.Errorf("Source is incorrect")
	}

	fmt.Printf("\nData Source is :%s", dataSource)
	return nil
}
