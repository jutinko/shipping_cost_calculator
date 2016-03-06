package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

const FIXERIO string = "https://api.fixer.io/latest?base=GBP&symbols=CNY"
const BACKOFFRATE float64 = 0.95

func Exchange(pounds float64) float64 {
	return pounds * getRate()
}

func getRate() float64 {
	// Url, err := url.Parse(serverbase)
	// if err != nil {
	// 	fmt.Printf("Parse failed: %s", err)
	// 	os.Exit(1)
	// }

	// Url.Path += "requestpair/"
	// params := url.Values{}
	// params.Add("name", name)
	// params.Add("location", location)
	// Url.RawQuery = params.Encode()
	response, err := http.Get(FIXERIO)

	if err != nil {
		fmt.Printf("Get failed: %s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Read content failed: %s", err)
			os.Exit(1)
		}

		var rate utilities.ForexRate
		json.Unmarshal(contents, &rate)
		fmt.Printf("%+v\n", rate.Rates["CNY"])
	}
	return 0
}

func main() {
	getRate()
}
