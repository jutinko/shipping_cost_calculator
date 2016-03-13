package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jutinko/shipping_cost_calculator/calculator"
)

type OrderListRequestCalculator func([]*calculator.ProductOrder) (float64, error)

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

		price, err := calc(orders)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(strconv.FormatFloat(price, 'f', 2, 64)))
	}
}
