package main

import (
	"flag"
	"fmt"

	"github.com/jutinko/shipping_cost_calculator/calculator"
)

var weight = flag.Float64(
	"weight",
	0,
	"the weight of the parcel",
)

func main() {
	calculator := &calculator.FiveOneParcelCalculator{}
	cost := calculator.CalculateCoreCost(*weight)
	fmt.Printf("%f", cost)
}
