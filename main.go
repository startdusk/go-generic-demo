package main

import (
	"fmt"

	"github.com/startdusk/go-generic-demo/business"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	windTest := business.Wind{
		Name:  "windTest",
		Netto: 100.00,
	}
	solarTest := business.Solar{
		Name:  "solarTest",
		Netto: 98.80,
	}
	fmt.Println(business.PrintGeneric(solarTest))
	fmt.Println(business.PrintGeneric(windTest))

	wind1 := business.Wind{Name: "wind1", Netto: 10.99}
	windArr := []business.Wind{wind1, {Name: "wind2", Netto: 9.98}}
	fmt.Println(business.PrintSlice(windArr))

	solarArr := []business.Solar{{Name: "solar1", Netto: 100}, {Name: "solar2", Netto: 10}}
	fmt.Println(business.PrintSlice(solarArr))

	fmt.Println(business.Cost(10, solarTest.Netto))
	solar3 := business.Solar{Name: "solar3", Netto: 100}
	solar4 := business.Solar{Name: "solar4", Netto: 10}
	solarArr2 := business.SolarSlice{solar3, solar4}
	fmt.Println(business.PrintSlice(solarArr2))
	fmt.Println(business.PrintSlice2(solarArr2))

	fmt.Printf("index: %d\n", slices.Index(solarArr2, solar4)) // 1
	business.SortByCost(solarArr2)
	fmt.Printf("index: %d\n", slices.Index(solarArr2, solar4)) // 0

	contracts := make(map[int]business.Solar)
	contracts[0] = solar3
	contracts[4] = solar4
	contractIDs := maps.Keys(contracts)
	fmt.Println(contractIDs)
}
