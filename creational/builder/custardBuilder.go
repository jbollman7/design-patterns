package main

type CustardBuilder struct {
    pieType string 
    bakeTime int
    isFruit bool
}

func newCustardBuilder() *CustardBuilder {
    return &CustardBuilder{}
}

func (p *CustardBuilder) setPieType() {
    p.pieType = "Custard"
}

func (p *CustardBuilder) setBakeTime() {
    p.bakeTime = 18
}

func (p *CustardBuilder) setIsFruit() {
    p.isFruit = false
}

func (p *CustardBuilder) getPie() Pie {
    return Pie {
        pieType: p.pieType,
        bakeTime: p.bakeTime,
        isFruit: p.isFruit,
    }
}
