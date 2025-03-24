package main

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setMaterialType()
	d.builder.setSize()
	d.builder.setHasBasement()
	d.builder.setNumberOfWindows()
	d.builder.setHasEnergyEfficientWindows()
	d.builder.setYardSize()
	d.builder.setGarageSize()
	d.builder.setRoofMaterial()
	d.builder.setHasAC()
	d.builder.setNumberOfPeaks()
	d.builder.setExteriorColor()
	d.builder.setInteriorColor()
	return d.builder.getHouse()
}
