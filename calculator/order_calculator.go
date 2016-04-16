package calculator

import (
	"strconv"

	"github.com/jutinko/shipping_cost_calculator/utilities"
)

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
	Exchange(float64) *utilities.Price
	NewRates()
}

type ProductOrder struct {
	Sku      int `json:"sku,string"`
	Quantity int `json:"quantity,string"`
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

func (o *OrderCalculator) GetPrice(orders []*ProductOrder) (*utilities.FinalPrice, error) {
	// make sure we use the most up to date rates
	o.currencyConverter.NewRates()

	var (
		price         float64
		shipping      float64
		wholePrice    float64
		totalQuantity int
		multiplier    float64
		exchange      *utilities.Price
	)

	simpleOrders := make(map[int]*utilities.SimpleOrder)
	simpleOrdersWhole := make(map[int]*utilities.SimpleOrder)

	for _, order := range orders {
		product, err := o.productStore.Get(order.Sku)
		if err != nil {
			return nil, err
		}

		multiplier = float64(order.Quantity)
		shipping = o.shippingCalculator.Calculate(&utilities.Parcel{
			Weight: product.Weight,
			Volume: product.Volume,
		})

		price = price + (product.Price+shipping)*multiplier
		wholePrice = wholePrice + (product.WholePrice+shipping)*multiplier

		i, ok := simpleOrders[product.Sku]
		if !ok {
			simpleOrders[product.Sku] = &utilities.SimpleOrder{
				Sku:       product.Sku,
				Quantity:  order.Quantity,
				Name:      product.Name,
				SellPrice: formatPrice(o.currencyConverter.Exchange(product.Price + shipping)),
			}
		} else {
			simpleOrders[product.Sku] = &utilities.SimpleOrder{
				Sku:       product.Sku,
				Quantity:  i.Quantity + order.Quantity,
				Name:      product.Name,
				SellPrice: formatPrice(o.currencyConverter.Exchange(product.Price + shipping)),
			}
		}

		i, ok = simpleOrdersWhole[product.Sku]
		if !ok {
			simpleOrdersWhole[product.Sku] = &utilities.SimpleOrder{
				Sku:       product.Sku,
				Quantity:  order.Quantity,
				Name:      product.Name,
				SellPrice: formatPrice(o.currencyConverter.Exchange(product.WholePrice + shipping)),
			}
		} else {
			simpleOrdersWhole[product.Sku] = &utilities.SimpleOrder{
				Sku:       product.Sku,
				Quantity:  i.Quantity + order.Quantity,
				Name:      product.Name,
				SellPrice: formatPrice(o.currencyConverter.Exchange(product.WholePrice + shipping)),
			}
		}
		totalQuantity = totalQuantity + order.Quantity
	}

	if totalQuantity < WholesaleThreshhold {
		exchange = o.currencyConverter.Exchange(price)
	} else {
		exchange = o.currencyConverter.Exchange(wholePrice)
		simpleOrders = simpleOrdersWhole
	}

	simpleOrdersSlice := []*utilities.SimpleOrder{}
	for _, value := range simpleOrders {
		simpleOrdersSlice = append(simpleOrdersSlice, value)
	}

	return &utilities.FinalPrice{
		Orders: simpleOrdersSlice,
		Price:  formatPrice(exchange),
	}, nil
}

func formatPrice(price *utilities.Price) *utilities.Price {
	if price == nil {
		return price
	}

	EUR, _ := strconv.ParseFloat(strconv.FormatFloat(price.EUR, 'f', 2, 64), 64)
	GBP, _ := strconv.ParseFloat(strconv.FormatFloat(price.GBP, 'f', 2, 64), 64)
	RMB, _ := strconv.ParseFloat(strconv.FormatFloat(price.RMB, 'f', 2, 64), 64)
	USD, _ := strconv.ParseFloat(strconv.FormatFloat(price.USD, 'f', 2, 64), 64)

	return &utilities.Price{
		EUR: EUR,
		GBP: GBP,
		RMB: RMB,
		USD: USD,
	}
}
