package main

import (
    "fmt"
)

func main() {

    // Create Builder
    // (Make a concrete builder that implements interface)
    // Pumpkin pie builder
    pumpkinBuilder := getBuilder("pumpkin")


    // Make director for pumpkin
    director := newDirector(pumpkinBuilder)

    pumpkinPie := director.buildPie()

    fmt.Printf("Pumpkin Pie pie type: %s\n", pumpkinPie.pieType)
    fmt.Printf("Pumpkin Pie has fruit?: %t\n", pumpkinPie.isFruit)
    fmt.Printf("Pumpkin Pie bake time: %d\n", pumpkinPie.bakeTime)

    // Custard pie builder
    custardBuilder := getBuilder("custard")

    // switch director for custard
    director.setBuilder(custardBuilder)

    custardPie := director.buildPie()

    fmt.Printf("Custard Pie pie type: %s\n", custardPie.pieType)
    fmt.Printf("Custard Pie has fruit?: %t\n", custardPie.isFruit)
    fmt.Printf("Custard Pie bake time: %d\n", custardPie.bakeTime)
} 
