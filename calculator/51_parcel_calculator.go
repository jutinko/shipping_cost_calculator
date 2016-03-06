package calculator

import "github.com/jutinko/shipping_cost_calculator/utilities"

const volumetricConversion utilities.Volume = 5000
const PickUpCost float64 = 5

type FiveOneParcelCalculator struct{}

func (c *FiveOneParcelCalculator) Calculate(p *utilities.Parcel) float64 {
	volumetricWeight := float64(p.Volume / volumetricConversion)
	if volumetricWeight > float64(p.Weight) {
		return c.CalculateCoreCost(volumetricWeight) + PickUpCost
	} else {
		return c.CalculateCoreCost(float64(p.Weight)) + PickUpCost
	}
}

func (c *FiveOneParcelCalculator) CalculateCoreCost(weight float64) float64 {
	if weight < 0 {
		return 0
	} else if weight < 5 {
		return 23.39
	} else if weight < 7 {
		return 25.99
	} else if weight < 10 {
		return 30.96
	} else if weight < 13 {
		return 35.94
	} else if weight < 15 {
		return 39.25
	}
	return 500
}
