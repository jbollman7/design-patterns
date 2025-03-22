package main

import (
    "fmt"
)

func main() {
    // Set bird builder
    birdHouseBuilder := newBirdHouseBuilder()

    strongHouseBuilder := newBrickHouseBuilder()

    // Declare director && Pass Builder to director
    director := newDirector(birdHouseBuilder)
    director.buildHouse()

    // test
    fmt.Printf("Type of house to build: %s\n", birdHouseBuilder.materialType)
    fmt.Printf("Type of house to build: %s\n", birdHouseBuilder.size)
    fmt.Printf("Type of house to build: %t\n", birdHouseBuilder.hasBasement)

    // Change director to different builder
    director.setBuilder(strongHouseBuilder)
    director.buildHouse()

    // test

    fmt.Printf("Type of house to build: %s\n", strongHouseBuilder.materialType)
    fmt.Printf("Type of house to build: %s\n", strongHouseBuilder.size)
    fmt.Printf("Type of house to build: %t\n", strongHouseBuilder.hasBasement)
}
