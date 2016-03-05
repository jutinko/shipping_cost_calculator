package utilities

import "sort"

type Parcel struct {
	Weight Weight
	length Length
	height Length
	width  Length
}

type Weight float64
type Length float64

func NewParcel(weight Weight, a, b, c Length) *Parcel {
	dim := []float64{float64(a), float64(b), float64(c)}

	sort.Float64s(dim)
	return &Parcel{
		Weight: weight,
		length: Length(dim[2]),
		height: Length(dim[1]),
		width:  Length(dim[0]),
	}
}

func (p *Parcel) Volume() float64 {
	return float64(p.length * p.height * p.width)
}
