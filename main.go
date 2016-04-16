package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/currency_converter"
	"github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/utilities"
)

func main() {
	shippingCalculator := &calculator.FiveOneParcelCalculator{}

	envReader := getEnvReader()

	dialOption := redis.DialPassword(envReader.GetPassword())
	client, err := redis.Dial("tcp", envReader.GetHost()+":"+envReader.GetPort(), dialOption)
	if err != nil {
		panic(fmt.Errorf("failed to connect to redis: %s", err))
	}

	println("Connected to redis")

	productStore := product_store.NewProductStore(client)
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

func getEnvReader() *utilities.EnvReader {
	var vcap_services string
	if vcap_services = os.Getenv("VCAP_SERVICES"); vcap_services == "" {
		panic("Failed to get environment variable VCAP_SERVICES")
	}

	envReader, err := utilities.NewEnvReader([]byte(vcap_services))
	if err != nil {
		panic(fmt.Errorf("Failed to parse the json for envReader: %s", err))
	}

	return envReader
}

func initProductStore(productStore *product_store.ProductStore) {
	priceLists := []string{"data/pricesNovember2015.csv", "data/pricesApril2016.csv"}

	sellMarginString := os.Getenv("sell_margin")
	sellMargin, err := strconv.ParseFloat(sellMarginString, 64)
	if err != nil {
		panic(fmt.Errorf("Failed to get sellMargin: %s", err))
	}

	wholeSellMarginString := os.Getenv("whole_sell_margin")
	wholeSellMargin, err := strconv.ParseFloat(wholeSellMarginString, 64)
	if err != nil {
		panic(fmt.Errorf("Failed to get wholeSellMargin: %s", err))
	}

	parser := &utilities.CsvParser{
		SellMargin:      sellMargin,
		WholeSellMargin: wholeSellMargin,
	}

	for _, priceList := range priceLists {
		products, err := parser.ParseFile(priceList)
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
