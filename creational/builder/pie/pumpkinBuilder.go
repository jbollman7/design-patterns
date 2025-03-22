package main

type PumpkinBuilder struct {
    pieType string 
    bakeTime int
    isFruit bool
}

func newPumpkinBuilder() *PumpkinBuilder {
    return &PumpkinBuilder{}
}

func (p *PumpkinBuilder) setPieType() {
    p.pieType = "Pumpkin"
}

func (p *PumpkinBuilder) setBakeTime() {
    p.bakeTime = 23
}

func (p *PumpkinBuilder) setIsFruit() {
    p.isFruit = true
}

func (p *PumpkinBuilder)getPie() Pie {
    return Pie {
        pieType: p.pieType,
        bakeTime: p.bakeTime,
        isFruit: p.isFruit,
    }
}
