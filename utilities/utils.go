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
