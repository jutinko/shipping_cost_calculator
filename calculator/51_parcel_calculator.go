package calculator

import "github.com/jutinko/shipping_cost_calculator/utilities"

// the average shipping price for one kilo
const shippingPricePerKilo float64 = float64(13) / 3
const volumetricConversion utilities.Volume = 4000

type FiveOneParcelCalculator struct{}

func (c *FiveOneParcelCalculator) Calculate(p *utilities.Parcel) float64 {
	volumetricWeight := float64(p.Volume / volumetricConversion)
	if volumetricWeight > float64(p.Weight) {
		return c.calculateCoreCost(volumetricWeight)
	} else {
		return c.calculateCoreCost(float64(p.Weight))
	}
}

func (c *FiveOneParcelCalculator) calculateCoreCost(weight float64) float64 {
	return shippingPricePerKilo * weight
}
