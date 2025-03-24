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
	birdHouse := director.buildHouse()

	// test
	fmt.Printf("Material Type of house to build: %s\n", birdHouse.materialType)
	fmt.Printf("Size of house to build: %s\n", birdHouse.size)
	fmt.Printf("Hosue has a basement: %t\n", birdHouse.hasBasement)
	fmt.Printf("Number of windows: %d\n", birdHouse.numberOfWindows)
	fmt.Printf("House has Energy Efficient Windows: %t\n", birdHouse.hasBasement)
	fmt.Printf("Yard size: %d\n", birdHouse.yardSize)
	fmt.Printf("Number of Cars that fit in garage: %d\n", birdHouse.yardSize)
	fmt.Printf("Roof material: %s\n", birdHouse.roofMaterial)
	fmt.Printf("House has AC: %t\n", birdHouse.hasAC)
	fmt.Printf("Number of Peaks: %d\n", birdHouse.numberOfPeaks)
	fmt.Printf("Exterior color of house: %s\n", birdHouse.exteriorColor)
	fmt.Printf("Interior color of house: %s\n", birdHouse.interiorColor)

	// Change director to different builder
	director.setBuilder(strongHouseBuilder)

	// build
	brickHouse := director.buildHouse()

	// test
	fmt.Printf("Material Type of house to build: %s\n", brickHouse.materialType)
	fmt.Printf("Size of house to build: %s\n", brickHouse.size)
	fmt.Printf("Hosue has a basement: %t\n", brickHouse.hasBasement)
	fmt.Printf("Number of windows: %d\n", brickHouse.numberOfWindows)
	fmt.Printf("House has Energy Efficient Windows: %t\n", brickHouse.hasBasement)
	fmt.Printf("Yard size: %d\n", brickHouse.yardSize)
	fmt.Printf("Number of Cars that fit in garage: %d\n", brickHouse.yardSize)
	fmt.Printf("Roof material: %s\n", brickHouse.roofMaterial)
	fmt.Printf("House has AC: %t\n", brickHouse.hasAC)
	fmt.Printf("Number of Peaks: %d\n", brickHouse.numberOfPeaks)
	fmt.Printf("Exterior color of house: %s\n", brickHouse.exteriorColor)
	fmt.Printf("Interior color of house: %s\n", brickHouse.interiorColor)

	// Test default, pass no Builder type explicit

	apartmentBuilder := newApartmentBuilder()
	director.setBuilder(apartmentBuilder)

	// build
	apartmentHouse := director.buildHouse()
	// test

	fmt.Printf("Material Type of house to build: %s\n", apartmentHouse.materialType)
	fmt.Printf("Size of house to build: %s\n", apartmentHouse.size)
	fmt.Printf("Hosue has a basement: %t\n", apartmentHouse.hasBasement)
	fmt.Printf("Number of windows: %d\n", apartmentHouse.numberOfWindows)
	fmt.Printf("House has Energy Efficient Windows: %t\n", apartmentHouse.hasBasement)
	fmt.Printf("Yard size: %d\n", apartmentHouse.yardSize)
	fmt.Printf("Number of Cars that fit in garage: %d\n", apartmentHouse.yardSize)
	fmt.Printf("Roof material: %s\n", apartmentHouse.roofMaterial)
	fmt.Printf("House has AC: %t\n", apartmentHouse.hasAC)
	fmt.Printf("Number of Peaks: %d\n", apartmentHouse.numberOfPeaks)
	fmt.Printf("Exterior color of house: %s\n", apartmentHouse.exteriorColor)
	fmt.Printf("Interior color of house: %s\n", apartmentHouse.interiorColor)
}
