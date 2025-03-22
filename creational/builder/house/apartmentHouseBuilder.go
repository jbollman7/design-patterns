package main

type ApartmentHouseBuilder struct {
    materialType string
    size string
    hasBasement bool
}

func newApartmentBuilder()* ApartmentHouseBuilder {
    return &ApartmentHouseBuilder{}
}

func(a *ApartmentHouseBuilder) setMaterialType() {
    a.materialType = "Steel"
}

func (a *ApartmentHouseBuilder) setSize() {
    a.size = "Huge"
}

func (a *ApartmentHouseBuilder) setHasBasement() {
    a.hasBasement = true
}

func (a *ApartmentHouseBuilder) getHouse() House {
    return House {
        materialType:  a.materialType,
        size: a.size,
        hasBasement: a.hasBasement,
    }
}
