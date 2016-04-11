package product_store

import (
	"encoding/json"
	"fmt"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

//go:generate counterfeiter -o fakes/FakeRedisConnection.go . RedisConnection
type RedisConnection interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

type ProductStore struct {
	Table  map[int]*utilities.Product
	client RedisConnection
}

func NewProductStore(client RedisConnection) *ProductStore {
	return &ProductStore{
		Table:  make(map[int]*utilities.Product),
		client: client,
	}
}

func (p *ProductStore) Put(sku int, product *utilities.Product) error {
	jsonBytes, _ := json.Marshal(product)
	_, err := p.client.Do("SET", sku, jsonBytes)
	if err != nil {
		return fmt.Errorf("failed to put: %d:%s", sku, err)
	}

	p.Table[sku] = product
	return nil
}

func (p *ProductStore) Get(sku int) (*utilities.Product, error) {
	value, err := p.client.Do("GET", sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %d:%s", sku, err)
	}
	var product *utilities.Product

	if productString, ok := value.([]byte); ok {
		json.Unmarshal([]byte(productString), &product)
	}
	return product, nil
}
