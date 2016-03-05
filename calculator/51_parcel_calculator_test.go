package calculator_test

import (
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/utilities"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("51_parcel_calculator", func() {
	Describe("CalculateCoreCost", func() {
		var calc *calculator.FiveOneParcelCalculator

		BeforeEach(func() {
			calc = &calculator.FiveOneParcelCalculator{}
		})

		Context("when the weight is less than 0", func() {
			It("should return 0", func() {
				price := calc.CalculateCoreCost(-3)
				Expect(price).Should(BeNumerically("==", 0))
			})
		})

		Context("when the weight is less than 5", func() {
			It("should return 23.39", func() {
				price := calc.CalculateCoreCost(3)
				Expect(price).Should(BeNumerically("==", 23.39))
			})
		})

		Context("when the weight is less than 7", func() {
			It("should return 25.99", func() {
				price := calc.CalculateCoreCost(6.8)
				Expect(price).Should(BeNumerically("==", 25.99))
			})
		})

		Context("when the weight is less than 10", func() {
			It("should return 30.96", func() {
				price := calc.CalculateCoreCost(8.8)
				Expect(price).Should(BeNumerically("==", 30.96))
			})

			It("should return 30.96", func() {
				price := calc.CalculateCoreCost(9.8)
				Expect(price).Should(BeNumerically("==", 30.96))
			})
		})

		Context("when the weight is less than 13", func() {
			It("should return 35.94", func() {
				price := calc.CalculateCoreCost(12.8)
				Expect(price).Should(BeNumerically("==", 35.94))
			})
		})

		Context("when the weight is less than 15", func() {
			It("should return 39.25", func() {
				price := calc.CalculateCoreCost(14.25)
				Expect(price).Should(BeNumerically("==", 39.25))
			})
		})

		Context("when the weight is more than 15", func() {
			It("should return 500", func() {
				price := calc.CalculateCoreCost(20.25)
				Expect(price).Should(BeNumerically("==", 500))
			})
		})
	})

	Describe("Calculate", func() {
		var (
			parcel *utilities.Parcel
			calc   *calculator.FiveOneParcelCalculator
		)

		Context("when the real cost is higher", func() {
			It("should honor the real cost", func() {
				parcel = utilities.NewParcel(utilities.Weight(6), utilities.Length(32), utilities.Length(12), utilities.Length(24))

				price := calc.Calculate(parcel)
				Expect(price).Should(BeNumerically("==", 25.99+calculator.PickUpCost))
			})
		})

		Context("when the volumetric cost is higher", func() {
			It("should honor the volumetric cost", func() {
				parcel = utilities.NewParcel(utilities.Weight(6), utilities.Length(62), utilities.Length(42), utilities.Length(24))

				price := calc.Calculate(parcel)
				Expect(price).Should(BeNumerically("==", 35.94+calculator.PickUpCost))
			})
		})
	})
})
