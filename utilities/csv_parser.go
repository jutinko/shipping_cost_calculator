package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Sku        int
	Name       string
	Weight     Weight  // in kg
	Volume     Volume  // in cm^3
	WholePrice float64 // in gbp
	Price      float64 // in gbp
}

func Parse(data string) (*Product, error) {
	fields := strings.Split(data, ",")
	if len(fields) < 6 {
		return nil, errors.New(fmt.Sprintf("missing field: %s", data))
	}

	sku, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}
	name := fields[1]
	wholePrice, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return nil, err
	}
	weightF, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}
	volumeF, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		return nil, err
	}

	return &Product{
		Sku:        sku,
		Name:       name,
		Weight:     Weight(weightF),
		Volume:     Volume(volumeF),
		WholePrice: wholePrice,
		Price:      price,
	}, nil
}

func ParseFile(filename string) ([]*Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var products []*Product

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p, err := Parse(scanner.Text())
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
