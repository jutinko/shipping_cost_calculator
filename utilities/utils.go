package utilities

type Weight float64
type Volume float64

type Parcel struct {
	Weight Weight
	Volume Volume
}

func NewParcel(weight Weight, volume Volume) *Parcel {
	return &Parcel{
		Weight: weight,
		Volume: volume,
	}
}

type SimpleOrder struct {
	Sku       int
	Quantity  int
	Name      string
	SellPrice *Price
}

type Product struct {
	Sku        int
	Name       string
	Weight     Weight  // in kg
	Volume     Volume  // in cm^3
	WholePrice float64 // in gbp
	Price      float64 // in gbp
}

type ForexRate struct {
	Base  string
	Date  string
	Rates map[string]float64
}

type Price struct {
	EUR float64
	GBP float64
	USD float64
	RMB float64
}

type FinalPrice struct {
	Orders []*SimpleOrder
	Price  *Price
}
