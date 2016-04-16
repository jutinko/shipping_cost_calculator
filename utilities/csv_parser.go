package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CsvParser struct {
	SellMargin      float64
	WholeSellMargin float64
}

func (c *CsvParser) Parse(data string) (*Product, error) {
	fields := strings.Split(data, ",")

	if len(fields) < 5 {
		return nil, errors.New(fmt.Sprintf("missing field: %s", data))
	}

	sku, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}
	name := fields[1]

	originalPrice, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}

	weightF, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return nil, err
	}

	volumeF, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}

	return &Product{
		Sku:        sku,
		Name:       name,
		Weight:     Weight(weightF),
		Volume:     Volume(volumeF),
		Price:      originalPrice * c.SellMargin,
		WholePrice: originalPrice * c.WholeSellMargin,
	}, nil
}

func (c *CsvParser) ParseFile(filename string) ([]*Product, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var products []*Product

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p, err := c.Parse(scanner.Text())
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
