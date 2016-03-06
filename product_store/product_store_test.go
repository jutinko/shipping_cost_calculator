package product_store_test

import (
	. "github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductStore", func() {
	var productStore *ProductStore

	BeforeEach(func() {
		productStore = &ProductStore{
			Table: make(map[int]*utilities.Product),
		}
	})

	Describe("Put", func() {
		It("can put", func() {
			productStore.Put(2, &utilities.Product{Sku: 2, Name: "nihao", Price: 3})
			product, _ := productStore.Get(2)
			Expect(product).To(Equal(&utilities.Product{Sku: 2, Name: "nihao", Price: 3}))
		})
	})

	Describe("Get", func() {
		Context("when the sku does not exist", func() {
			It("returns an error", func() {
				_, err := productStore.Get(321)
				Expect(err).To(MatchError("no-product"))
			})
		})
	})
})
