package main

type Director struct{
    builder Builder;
}

func newDirector(b Builder) *Director {
    return &Director {
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
    return d.builder.getHouse()
}
