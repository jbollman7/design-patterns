package main

import "fmt"

type Extractor struct {
	dataSource string
}

func newExtractor(dataSource string) *Extractor {
	return &Extractor{
		dataSource: dataSource,
	}
}

func (e *Extractor) extractData(dataSource string) error {
	if e.dataSource != dataSource {
		return fmt.Errorf("Source is incorrect")
	}
	fmt.Printf("\nData Source is: %s", dataSource)
	return nil
}
