package main

import (
	"encoding/json"
	"net/http"

	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/utilities"
)

type OrderListRequestCalculator func([]*calculator.ProductOrder) (*utilities.FinalPrice, error)

func OrderListRequestHandler(calc OrderListRequestCalculator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		orders := make([]*calculator.ProductOrder, 0)

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&orders); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			println(err.Error())
			return
		}

		w.WriteHeader(200)
		price, err := calc(orders)
		if err != nil {
			result, _ := json.Marshal(err.Error())
			w.Write(result)
			return
		}

		w.Write(convPrice(price))
	}
}

func convPrice(price *utilities.FinalPrice) []byte {
	result, _ := json.Marshal(price)
	return result
}
