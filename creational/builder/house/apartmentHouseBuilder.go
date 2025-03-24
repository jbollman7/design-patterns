package main

type ApartmentHouseBuilder struct {
	materialType              string
	size                      string
	hasBasement               bool
	numberOfWindows           int8
	hasEnergyEfficientWindows bool
	yardSize                  int16
	garageSize                int8
	roofMaterial              string
	hasAC                     bool
	numberOfPeaks             int8
	exteriorColor             string
	interiorColor             string
}

func newApartmentBuilder() *ApartmentHouseBuilder {
	return &ApartmentHouseBuilder{}
}

func (a *ApartmentHouseBuilder) setMaterialType() {
	a.materialType = "Steel"
}

func (a *ApartmentHouseBuilder) setSize() {
	a.size = "Huge"
}

func (a *ApartmentHouseBuilder) setHasBasement() {
	a.hasBasement = true
}

func (a *ApartmentHouseBuilder) setNumberOfWindows() {
	a.numberOfWindows = 96
}

func (a *ApartmentHouseBuilder) setHasEnergyEfficientWindows() {
	a.hasEnergyEfficientWindows = false
}

func (a *ApartmentHouseBuilder) setYardSize() {
	a.yardSize = 0
}

func (a *ApartmentHouseBuilder) setGarageSize() {
	a.garageSize = 24
}

func (a *ApartmentHouseBuilder) setRoofMaterial() {
	a.roofMaterial = "asphalt"
}

func (a *ApartmentHouseBuilder) setHasAC() {
	a.hasAC = true
}

func (a *ApartmentHouseBuilder) setNumberOfPeaks() {
	a.numberOfPeaks = 31
}

func (a *ApartmentHouseBuilder) setExteriorColor() {
	a.exteriorColor = "yellow"
}

func (a *ApartmentHouseBuilder) setInteriorColor() {
	a.interiorColor = "light-tan"
}

func (a *ApartmentHouseBuilder) getHouse() House {
	return House{
		materialType:              a.materialType,
		size:                      a.size,
		hasBasement:               a.hasBasement,
		numberOfWindows:           a.numberOfWindows,
		hasEnergyEfficientWindows: a.hasEnergyEfficientWindows,
		yardSize:                  a.yardSize,
		garageSize:                a.garageSize,
		roofMaterial:              a.roofMaterial,
		hasAC:                     a.hasAC,
		numberOfPeaks:             a.numberOfPeaks,
		exteriorColor:             a.exteriorColor,
		interiorColor:             a.interiorColor,
	}
}
