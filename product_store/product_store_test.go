package product_store_test

import (
	"encoding/json"
	"errors"

	. "github.com/jutinko/shipping_cost_calculator/product_store"
	"github.com/jutinko/shipping_cost_calculator/product_store/fakes"
	"github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductStore", func() {
	var (
		productStore        *ProductStore
		fakeRedisConnection *fakes.FakeRedisConnection
		product             *utilities.Product
	)

	BeforeEach(func() {
		fakeRedisConnection = new(fakes.FakeRedisConnection)
		productStore = NewProductStore(fakeRedisConnection)
		product = &utilities.Product{Sku: 2, Name: "nihao", Price: 3}
		fakeRedisConnection.DoStub = func(cmd string, args ...interface{}) (interface{}, error) {
			if cmd == "SET" && args[0] == -2 {
				return nil, errors.New("dodo failed")
			} else if cmd == "GET" && args[0] == -2 {
				return nil, errors.New("no-product: -2")
			} else if cmd == "GET" && args[0] == 2 {
				jsonBytes, _ := json.Marshal(product)
				return jsonBytes, nil
			}
			return nil, nil
		}
	})

	Describe("Put", func() {
		It("can put", func() {
			productStore.Put(2, product)

			Expect(fakeRedisConnection.DoCallCount()).To(Equal(1))
			cmd, params := fakeRedisConnection.DoArgsForCall(0)
			Expect(cmd).To(Equal("SET"))
			Expect(params[0]).To(Equal(2))
			jsonString, _ := json.Marshal(product)
			Expect(params[1]).To(Equal(jsonString))
		})

		Context("when the client fails to set", func() {
			It("returns the error", func() {
				err := productStore.Put(-2, product)
				Expect(err).To(MatchError(ContainSubstring("dodo failed")))
			})
		})
	})

	Describe("Get", func() {
		It("deligates the call to the client", func() {
			actualProduct, err := productStore.Get(2)

			Expect(err).NotTo(HaveOccurred())
			Expect(actualProduct).To(Equal(product))
			cmd, params := fakeRedisConnection.DoArgsForCall(0)
			Expect(cmd).To(Equal("GET"))
			Expect(params[0]).To(Equal(2))
		})

		Context("when the sku does not exist", func() {
			It("returns an error", func() {
				_, err := productStore.Get(-2)
				Expect(err).To(MatchError(ContainSubstring("no-product: -2")))
			})
		})
	})
})
