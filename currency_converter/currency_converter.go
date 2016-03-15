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
	Api string
}

func (c *CurrencyConverter) Exchange(pounds float64) *utilities.Price {
	rates := c.getRates()

	return &utilities.Price{
		EUR: pounds * rates["EUR"] * EXTRARATE,
		GBP: pounds,
		RMB: pounds * rates["CNY"] * EXTRARATE,
		USD: pounds * rates["USD"] * EXTRARATE,
	}
}

func (c *CurrencyConverter) getRates() map[string]float64 {
	response, err := http.Get(c.Api)
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
