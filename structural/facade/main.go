package main

import "fmt"

// Facade class! this is what interfaces with all
type ETLProcessor struct {
	extractor   *Extractor
	transformer *Transformer
	loader      *Loader
}

// func to create ponter to new instance of newEtlProcessor
func newEtlProcessor(source string, target map[string]string) *ETLProcessor {

	// call extractor
	extractor := newExtractor(source)

	// perform tranformating

	transform := newTransformer()

	// call load
	// loader requires the map
	load := newLoader(target)

	// in here i need to generate data to return pointer to ETLPRocessor

	return &ETLProcessor{
		extractor:   extractor,
		transformer: transform,
		loader:      load,
	}
}

func (e *ETLProcessor) getData(dataSource string, target map[string]string) {
	fmt.Printf("\nCalling getData on facade")
	e.extractor.extractData(dataSource)
	e.transformer.transform("transformation undertaking")
	e.loader.loadData(target)
	fmt.Printf("\nETL process is complete")
}

func main() {
	// In here, i dont want to touch the n number of subsystems. just use the facade.
	// The facade is a simplified interface to do my work that i want done
	const dataSource = "MySql"
	const dataTarget = "Posgres"
	loader := make(map[string]string)
	loader[dataTarget] = "yes"

	// use facade
	// We new up the Facade class
	// Then we only call methods on facade, never the underlying classes
	etl := newEtlProcessor(dataSource, loader)

	etl.getData(dataSource, loader)
}
