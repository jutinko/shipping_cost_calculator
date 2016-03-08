package currency_converter_test

import (
	. "github.com/jutinko/shipping_cost_calculator/currency_converter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CurrencyConverter", func() {
	var currencyConverter *CurrencyConverter
	BeforeEach(func() {
		currencyConverter = &CurrencyConverter{Api: "https://api.fixer.io/latest?base=GBP&symbols=CNY"}
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3)).To(BeNumerically(">", 3))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3)).To(BeNumerically("<", 30))
	})

	Context("when the internet is down", func() {
		It("uses the ballback rate", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&symbols=CNY"}
			Expect(currencyConverter.Exchange(3)).To(BeNumerically("==", 28.5))
		})
	})
})
