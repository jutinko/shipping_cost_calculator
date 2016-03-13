package currency_converter_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/jutinko/shipping_cost_calculator/currency_converter"
	"github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const api string = "https://api.fixer.io/latest?base=GBP&symbols=CNY"

var _ = Describe("CurrencyConverter", func() {
	var currencyConverter *CurrencyConverter
	BeforeEach(func() {
		currencyConverter = &CurrencyConverter{Api: api}
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3)).To(BeNumerically(">", 3))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3)).To(BeNumerically("<", 30))
	})

	It("adds EXTRARATE to every pound we convert to rmb", func() {
		response, err := http.Get(api)
		Expect(err).NotTo(HaveOccurred())

		defer response.Body.Close()
		// untested code!
		contents, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())

		var rate utilities.ForexRate
		json.Unmarshal(contents, &rate)

		actual := currencyConverter.Exchange(3)
		Expect(actual).To(BeNumerically("~", 3*(rate.Rates["CNY"]+0.2), 1*0.2))
	})

	Context("when the internet is down", func() {
		It("uses the ballback rate", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&symbols=CNY"}
			Expect(currencyConverter.Exchange(3)).To(BeNumerically("==", 28.5))
		})
	})
})
