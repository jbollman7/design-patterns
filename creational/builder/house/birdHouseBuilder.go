package main

type BirdHouseBuilder struct {
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

func newBirdHouseBuilder() *BirdHouseBuilder {
	return &BirdHouseBuilder{}
}

func (h *BirdHouseBuilder) setMaterialType() {
	h.materialType = "Wood"
}

func (h *BirdHouseBuilder) setSize() {
	h.size = "Small"
}

func (h *BirdHouseBuilder) setHasBasement() {
	h.hasBasement = false
}

func (h *BirdHouseBuilder) setNumberOfWindows() {
	h.numberOfWindows = 0
}

func (h *BirdHouseBuilder) setHasEnergyEfficientWindows() {
	h.hasEnergyEfficientWindows = false
}

func (h *BirdHouseBuilder) setYardSize() {
	h.yardSize = 0
}

func (h *BirdHouseBuilder) setGarageSize() {
	h.garageSize = 0
}

func (h *BirdHouseBuilder) setRoofMaterial() {
	h.roofMaterial = "wood"
}

func (h *BirdHouseBuilder) setHasAC() {
	h.hasAC = false
}

func (h *BirdHouseBuilder) setNumberOfPeaks() {
	h.numberOfPeaks = 1
}

func (h *BirdHouseBuilder) setExteriorColor() {
	h.exteriorColor = "blue"
}

func (h *BirdHouseBuilder) setInteriorColor() {
	h.interiorColor = "unpainted"
}

func (h *BirdHouseBuilder) getHouse() House {
	return House{
		materialType:              h.materialType,
		size:                      h.size,
		hasBasement:               h.hasBasement,
		numberOfWindows:           h.numberOfWindows,
		hasEnergyEfficientWindows: h.hasEnergyEfficientWindows,
		yardSize:                  h.yardSize,
		garageSize:                h.garageSize,
		roofMaterial:              h.roofMaterial,
		hasAC:                     h.hasAC,
		numberOfPeaks:             h.numberOfPeaks,
		exteriorColor:             h.exteriorColor,
		interiorColor:             h.interiorColor,
	}
}
