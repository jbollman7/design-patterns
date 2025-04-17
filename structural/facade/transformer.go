package main

import "fmt"

type Transformer struct {
}

func newTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) transform(input string) error {
	fmt.Printf("\nInput passed to transformer: %s", input)
	return nil
}
