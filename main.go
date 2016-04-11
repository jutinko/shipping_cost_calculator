package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/currency_converter"
	"github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/utilities"
)

func main() {
	shippingCalculator := &calculator.FiveOneParcelCalculator{}

	dialOption := redis.DialPassword("d7c2b08c-bc18-45d2-b728-d83ef331e72f")

	client, err := redis.Dial("tcp", "192.168.8.153:46593", dialOption)
	// client, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(fmt.Errorf("failed to connect to redis: %s", err))
	}

	println("Connected to redis")

	productStore := product_store.NewProductStore(client)
	//initProductStore(productStore)

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
	priceLists := []string{"data/pricesNovember2015.csv", "data/pricesApril2016.csv"}
	for _, priceList := range priceLists {
		products, err := utilities.ParseFile(priceList)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		for _, product := range products {
			err := productStore.Put(product.Sku, product)
			if err != nil {
				panic(fmt.Errorf("Failed to put %#v to store: %s", product, err))
			}
		}
	}
}
