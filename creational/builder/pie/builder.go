package main

type Builder interface {
    setPieType()
    setBakeTime()
    setIsFruit()
    getPie() Pie
}

func getBuilder(builderType string) Builder {
    if builderType == "pumpkin" {
        return newPumpkinBuilder()
    }

    if builderType == "custard" {
        return newCustardBuilder()
    }

    return nil
}
