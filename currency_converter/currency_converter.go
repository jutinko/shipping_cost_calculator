package currency_converter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

const BACKOFFRATE float64 = 9.5

type CurrencyConverter struct {
	Api string
}

func (c *CurrencyConverter) Exchange(pounds float64) float64 {
	return pounds * c.getRate()
}

func (c *CurrencyConverter) getRate() float64 {
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

	var rate utilities.ForexRate
	json.Unmarshal(contents, &rate)
	return rate.Rates["CNY"]
}
