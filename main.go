package main

import (
	"flag"
	"fmt"

	"github.com/jutinko/go_practice/shipping_cost_calculator/calculator"
)

var weight = flag.Float64(
	"weight",
	0,
	"the weight of the parcel",
)

func main() {
	calculator := &calculator.Calculator{}
	cost := calculator.Calculate(*weight)
	fmt.Printf("%f", cost)
}
