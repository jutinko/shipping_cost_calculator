package utilities_test

import (
	. "github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parcel", func() {
	Describe("Volume", func() {
		It("gets the correct volume", func() {
			parcel := NewParcel(Weight(0), Length(10), Length(20), Length(10))
			Expect(parcel.Volume()).To(BeNumerically("==", 2000))
		})
	})
})
