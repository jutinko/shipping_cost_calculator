package product_store

import (
	"errors"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

type ProductStore struct {
	Table map[int]*utilities.Product
}

func (p *ProductStore) Put(sku int, product *utilities.Product) {
	p.Table[sku] = product
}

func (p *ProductStore) Get(sku int) (*utilities.Product, error) {
	if val, ok := p.Table[sku]; ok {
		return val, nil
	}
	return nil, errors.New("no-product")
}
