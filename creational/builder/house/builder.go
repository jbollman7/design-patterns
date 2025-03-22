package main

type Builder interface {
    setMaterialType()
    setSize()
    setHasBasement()
    getHouse() House
}

// Construct new Builder
func getNewBuilder(houseType string) Builder {
    if (houseType == "Bird") {
        return newBirdHouseBuilder()
    }

    if (houseType == "Strong") {
        return newBrickHouseBuilder()
    }

    if (houseType == "Huge") {
        return newApartmentBuilder()
    }

    return nil
}
