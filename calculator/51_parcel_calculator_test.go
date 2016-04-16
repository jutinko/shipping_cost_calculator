package calculator_test

import (
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/utilities"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("51_parcel_calculator", func() {
	Describe("Calculate", func() {
		var (
			parcel *utilities.Parcel
			calc   *calculator.FiveOneParcelCalculator
		)

		Context("when the real cost is higher", func() {
			It("should honor the real cost", func() {
				parcel = utilities.NewParcel(utilities.Weight(1), utilities.Volume(2*12*24))

				price := calc.Calculate(parcel)
				Expect(price).Should(BeNumerically("==", 1.3))
			})
		})

		Context("when the volumetric cost is higher", func() {
			It("should honor the volumetric cost", func() {
				parcel = utilities.NewParcel(utilities.Weight(1), utilities.Volume(62*42*24))

				price := calc.Calculate(parcel)
				Expect(price).Should(BeNumerically("~", 20.311, 0.001))
			})
		})
	})
})
