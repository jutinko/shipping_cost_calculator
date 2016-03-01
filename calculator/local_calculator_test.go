package calculator_test

import (
	. "github.com/jutinko/go_practice/shipping_cost_calculator/calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LocalCalculator", func() {
	Describe("Calculate", func() {
		var calculator *Calculator

		BeforeEach(func() {
			calculator = &Calculator{}
		})

		Context("when the weight is less than 0", func() {
			It("should return 0", func() {
				price := calculator.Calculate(-3)
				Expect(price).Should(BeNumerically("==", 0))
			})
		})
	})

})
