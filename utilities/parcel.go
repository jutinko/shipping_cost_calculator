package utilities

type Parcel struct {
	Weight Weight
	Volume Volume
}

type Weight float64
type Volume float64

func NewParcel(weight Weight, volume Volume) *Parcel {
	return &Parcel{
		Weight: weight,
		Volume: volume,
	}
}
