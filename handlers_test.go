package main_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/jutinko/shipping_cost_calculator"
	"github.com/jutinko/shipping_cost_calculator/calculator"
	"github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func stubbedOrderListRequestCalculator(price *utilities.Price, err error) OrderListRequestCalculator {
	return func(_ []*calculator.ProductOrder) (*utilities.Price, error) {
		return price, err
	}
}

var _ = Describe("Handlers", func() {
	singleOrder := func() io.Reader { return strings.NewReader(`[{"sku": "30", "quantity": "2"}]`) }
	var fakePrice *utilities.Price

	BeforeEach(func() {
		fakePrice = &utilities.Price{
			EUR: 1,
			GBP: 2,
			RMB: 3,
			USD: 4,
		}
	})

	Describe("getting the final price", func() {
		It("responds with 200", func() {
			handle := OrderListRequestHandler(stubbedOrderListRequestCalculator(fakePrice, nil))

			resp := httptest.NewRecorder()
			req, err := http.NewRequest(
				"GET", "/get_order_price", singleOrder(),
			)

			handle(resp, req)

			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Code).To(Equal(200))
		})

		Context("when the calculator returns an error", func() {
			It("responds with 500", func() {
				handle := OrderListRequestHandler(stubbedOrderListRequestCalculator(fakePrice, errors.New("bad sku")))

				resp := httptest.NewRecorder()
				req, err := http.NewRequest(
					"GET", "/get_order_price", singleOrder(),
				)

				handle(resp, req)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Code).To(Equal(500))
			})
		})

		Context("when there are multiple orders", func() {
			multipleOrder := func() io.Reader {
				return strings.NewReader(`[{"sku": "30", "quantity": "2"}, {"sku": "2", "quantity": "4"}]`)
			}

			It("responds with 500", func() {
				handle := OrderListRequestHandler(stubbedOrderListRequestCalculator(fakePrice, nil))

				resp := httptest.NewRecorder()
				req, err := http.NewRequest(
					"GET", "/get_order_price", multipleOrder(),
				)

				handle(resp, req)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Code).To(Equal(200))
			})
		})
	})
})
