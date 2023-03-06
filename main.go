package main

import (
	"fmt"

	"github.com/startdusk/go-generic-demo/business"
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

	windArr := []business.Wind{{Name: "wind1", Netto: 10.99}, {Name: "wind2", Netto: 9.98}}
	fmt.Println(business.PrintSlice(windArr))

	solarArr := []business.Solar{{Name: "solar1", Netto: 100}, {Name: "solar2", Netto: 10}}
	fmt.Println(business.PrintSlice(solarArr))

	fmt.Println(business.Cost(10, solarTest.Netto))

	solarArr2 := business.SolarSlice{{Name: "solar3", Netto: 100}}
	fmt.Println(business.PrintSlice(solarArr2))
	fmt.Println(business.PrintSlice2(solarArr2))
}
