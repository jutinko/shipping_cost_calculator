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

const api string = "https://api.fixer.io/latest?base=GBP"

var _ = Describe("CurrencyConverter", func() {
	var currencyConverter *CurrencyConverter
	BeforeEach(func() {
		currencyConverter = &CurrencyConverter{Api: api}
	})

	It("can get the gbp rate", func() {
		Expect(currencyConverter.Exchange(3).GBP).To(BeNumerically("==", 3))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3).RMB).To(BeNumerically(">", 24))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3).RMB).To(BeNumerically("<", 30))
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
		Expect(actual.RMB).To(BeNumerically("~", 3*(rate.Rates["CNY"]*1.04), 0.1))
	})

	It("can get the gbp-eur rate", func() {
		Expect(currencyConverter.Exchange(3).EUR).To(BeNumerically(">", 3))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3).EUR).To(BeNumerically("<", 5))
	})

	It("adds EXTRARATE to every pound we convert to eur", func() {
		response, err := http.Get(api)
		Expect(err).NotTo(HaveOccurred())

		defer response.Body.Close()
		// untested code!
		contents, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())

		var rate utilities.ForexRate
		json.Unmarshal(contents, &rate)

		actual := currencyConverter.Exchange(3)
		Expect(actual.EUR).To(BeNumerically("~", 3*(rate.Rates["EUR"]*1.04), 0.1))
	})

	It("can get the gbp-usd rate", func() {
		Expect(currencyConverter.Exchange(3).USD).To(BeNumerically(">", 3))
	})

	It("can get the gbp-rmb rate", func() {
		Expect(currencyConverter.Exchange(3).USD).To(BeNumerically("<", 6))
	})

	It("adds EXTRARATE to every pound we convert to usd", func() {
		response, err := http.Get(api)
		Expect(err).NotTo(HaveOccurred())

		defer response.Body.Close()
		// untested code!
		contents, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())

		var rate utilities.ForexRate
		json.Unmarshal(contents, &rate)

		actual := currencyConverter.Exchange(3)
		Expect(actual.USD).To(BeNumerically("~", 3*(rate.Rates["USD"]*1.04), 0.1))
	})

	Context("when the internet is down", func() {
		FIt("uses the ballback rate for USD", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&HJHH"}
			Expect(currencyConverter.Exchange(3).RMB).To(BeNumerically("==", 29.64))
		})

		It("uses the ballback rate", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&JJJ"}
			Expect(currencyConverter.Exchange(3).GBP).To(BeNumerically("==", 3))
		})

		It("uses the ballback rate", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&FDS"}
			Expect(currencyConverter.Exchange(3).EUR).To(BeNumerically("==", 4.368))
		})

		It("uses the ballback rate", func() {
			currencyConverter = &CurrencyConverter{Api: "httpi.fixer.io/latest?base=GBP&III"}
			Expect(currencyConverter.Exchange(3).RMB).To(BeNumerically("==", 4.68))
		})
	})
})
