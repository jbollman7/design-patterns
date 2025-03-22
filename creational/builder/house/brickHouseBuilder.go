package main

type BrickHouseBuilder struct {
    materialType string
    size string
    hasBasement bool
}


func newBrickHouseBuilder() *BrickHouseBuilder{
    return &BrickHouseBuilder{}
}

func (h *BrickHouseBuilder)setMaterialType() {
    h.materialType = "Brick"
}

func (h *BrickHouseBuilder)setSize() {
    h.size = "Medium"
}

func (h *BrickHouseBuilder)setHasBasement() {
    h.hasBasement = true
}

func (h *BrickHouseBuilder)getHouse() House {
    return House {
        materialType:  h.materialType,
        size:  h.size,
        hasBasement: h.hasBasement,
    }
}
