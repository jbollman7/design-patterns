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

func (d *Director) buildPie() Pie {
    d.builder.setPieType()
    d.builder.setBakeTime()
    d.builder.setIsFruit()
    return d.builder.getPie()
}
