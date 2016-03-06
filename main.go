package main

import (
	"os"

	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/utilities"
)

func main() {
	shippingCalculator := &calculator.FiveOneParcelCalculator{}

	productStore := &product_store.ProductStore{}

	orderCalculator := calculator.NewOrderCalculator(productStore, shippingCalculator, currencyConverter)
}

func initProductStore(productStore *product_store.ProductStore) {
	products, err := utilities.ParseFile("data/redisImport.csv")
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for _, product := range products {
		productStore.Put(product.Sku, product)
	}
}
