package main

type BrickHouseBuilder struct {
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

func newBrickHouseBuilder() *BrickHouseBuilder {
	return &BrickHouseBuilder{}
}

func (h *BrickHouseBuilder) setMaterialType() {
	h.materialType = "Brick"
}

func (h *BrickHouseBuilder) setSize() {
	h.size = "Medium"
}

func (h *BrickHouseBuilder) setHasBasement() {
	h.hasBasement = true
}

func (h *BrickHouseBuilder) setNumberOfWindows() {
	h.numberOfWindows = 12
}

func (h *BrickHouseBuilder) setHasEnergyEfficientWindows() {
	h.hasEnergyEfficientWindows = true
}

func (h *BrickHouseBuilder) setYardSize() {
	h.yardSize = 5
}

func (h *BrickHouseBuilder) setGarageSize() {
	h.garageSize = 3
}

func (h *BrickHouseBuilder) setRoofMaterial() {
	h.roofMaterial = "steel"
}

func (h *BrickHouseBuilder) setHasAC() {
	h.hasAC = true
}

func (h *BrickHouseBuilder) setNumberOfPeaks() {
	h.numberOfPeaks = 4
}

func (h *BrickHouseBuilder) setExteriorColor() {
	h.exteriorColor = "red"
}

func (h *BrickHouseBuilder) setInteriorColor() {
	h.interiorColor = "light-gray"
}

func (h *BrickHouseBuilder) getHouse() House {
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
