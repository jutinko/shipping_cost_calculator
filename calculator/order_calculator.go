package calculator

import "github.com/jutinko/shipping_cost_calculator/utilities"

const WholesaleThreshhold int = 15

//go:generate counterfeiter -o fakes/FakeProductStore.go . ProductStore
type ProductStore interface {
	Get(int) (*utilities.Product, error)
}

//go:generate counterfeiter -o fakes/FakeShippingCalculator.go . ShippingCalculator
type ShippingCalculator interface {
	Calculate(*utilities.Parcel) float64
}

//go:generate counterfeiter -o fakes/FakeCurrencyConverter.go . CurrencyConverter
type CurrencyConverter interface {
	Exchange(float64) float64
}

type ProductOrder struct {
	Sku      int
	Quantity int
}

func NewProductOrder(sku, quantity int) *ProductOrder {
	return &ProductOrder{
		Sku:      sku,
		Quantity: quantity,
	}
}

type OrderCalculator struct {
	productStore       ProductStore
	shippingCalculator ShippingCalculator
	currencyConverter  CurrencyConverter
}

func NewOrderCalculator(productStore ProductStore, shippingCalculator ShippingCalculator, currencyConverter CurrencyConverter) *OrderCalculator {
	return &OrderCalculator{
		productStore:       productStore,
		shippingCalculator: shippingCalculator,
		currencyConverter:  currencyConverter,
	}
}

func (o *OrderCalculator) GetPrice(orders []*ProductOrder) (float64, error) {
	var (
		price         float64
		wholePrice    float64
		weight        float64
		volume        float64
		totalQuantity int
		multiplier    float64
	)

	for _, order := range orders {
		product, err := o.productStore.Get(order.Sku)
		if err != nil {
			return 0, err
		}

		multiplier = float64(order.Quantity)
		weight = weight + float64(product.Weight)*multiplier
		volume = volume + float64(product.Volume)*multiplier
		price = price + product.Price*multiplier
		wholePrice = wholePrice + product.WholePrice*multiplier
		totalQuantity = totalQuantity + order.Quantity
	}

	shippingBit := o.shippingCalculator.Calculate(
		&utilities.Parcel{
			Weight: utilities.Weight(weight),
			Volume: utilities.Volume(volume),
		})

	if totalQuantity < WholesaleThreshhold {
		return o.currencyConverter.Exchange(shippingBit + price), nil
	}
	return o.currencyConverter.Exchange(shippingBit + wholePrice), nil
}
