package main


type BirdHouseBuilder struct {
    materialType string
    size string
    hasBasement bool
}

func newBirdHouseBuilder() *BirdHouseBuilder {
    return &BirdHouseBuilder{}
}

func (h *BirdHouseBuilder)setMaterialType() {
    h.materialType = "Wood"
}

func (h *BirdHouseBuilder)setSize() {
    h.size = "Small"
}

func (h *BirdHouseBuilder)setHasBasement() {
    h.hasBasement = false
}

func (h *BirdHouseBuilder)getHouse() House {
    return House {
        materialType:  h.materialType,
        size:          h.size,
        hasBasement:   h.hasBasement,
    }
}
