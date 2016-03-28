package currency_converter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

var BACKOFFRATE = map[string]float64{
	"EUR": 1.4,
	"GBP": 0,
	"CNY": 9.5,
	"USD": 1.5,
}

// this constant is used to protect from extreme forex senarios
const EXTRARATE float64 = 1.04

type CurrencyConverter struct {
	Api   string
	rates map[string]float64
}

func NewCurrencyConverter(api string) *CurrencyConverter {
	return &CurrencyConverter{
		Api:   api,
		rates: getRates(api),
	}
}

func (c *CurrencyConverter) NewRates() {
	c.rates = getRates(c.Api)
}

func (c *CurrencyConverter) Exchange(pounds float64) *utilities.Price {
	return &utilities.Price{
		EUR: pounds * c.rates["EUR"] * EXTRARATE,
		GBP: pounds,
		RMB: pounds * c.rates["CNY"] * EXTRARATE,
		USD: pounds * c.rates["USD"] * EXTRARATE,
	}
}

func getRates(api string) map[string]float64 {
	response, err := http.Get(api)
	if err != nil {
		return BACKOFFRATE
	}

	defer response.Body.Close()
	// untested code!
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return BACKOFFRATE
	}

	var rate *utilities.ForexRate
	json.Unmarshal(contents, &rate)
	return rate.Rates
}
