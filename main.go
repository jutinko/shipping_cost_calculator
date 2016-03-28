package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/currency_converter"
	"github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/utilities"
)

func main() {
	shippingCalculator := &calculator.FiveOneParcelCalculator{}

	productStore := product_store.NewProductStore()
	initProductStore(productStore)

	currencyConverter := currency_converter.NewCurrencyConverter("https://api.fixer.io/latest?base=GBP")

	orderCalculator := calculator.NewOrderCalculator(productStore, shippingCalculator, currencyConverter)

	router := mux.NewRouter()
	router.
		HandleFunc("/get_order_price", OrderListRequestHandler(orderCalculator.GetPrice)).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%v", getPort()), router); err != nil {
		log.Fatalln(err)
	}
}

func getPort() string {
	if configuredPort := os.Getenv("PORT"); configuredPort == "" {
		return "3000"
	} else {
		return configuredPort
	}
}

func initProductStore(productStore *product_store.ProductStore) {
	products, err := utilities.ParseFile("data/prices.csv")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, product := range products {
		productStore.Put(product.Sku, product)
	}
}
