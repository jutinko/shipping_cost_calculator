package utilities_test

import (
	"io/ioutil"
	"os"

	. "github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csv_parser", func() {
	var (
		product         *Product
		parser          *CsvParser
		sellMargin      float64 = 1.3
		wholeSellMargin float64 = 1.45
	)

	BeforeEach(func() {
		parser = &CsvParser{
			SellMargin:      sellMargin,
			WholeSellMargin: wholeSellMargin,
		}
	})

	Describe("Parse", func() {
		It("can parse the line to a parcel", func() {
			product, _ = parser.Parse("10072,H&B MEGA VITAMIN TEENS,6.31925,0.3,300")
			Expect(product.Sku).To(Equal(10072))
			Expect(product.Name).To(Equal("H&B MEGA VITAMIN TEENS"))
			Expect(product.Weight).To(BeNumerically("==", 0.3))
			Expect(product.Volume).To(BeNumerically("==", 300))
			Expect(product.Price).To(BeNumerically("==", 6.31925*sellMargin))
			Expect(product.WholePrice).To(BeNumerically("==", 6.31925*wholeSellMargin))
		})

		Context("when there is not enough parsable fields", func() {
			It("returns parse error", func() {
				_, err := parser.Parse("H&B MEGA VITAMIN TEENS,6.31925,0.3,300")
				Expect(err).To(MatchError("missing field: H&B MEGA VITAMIN TEENS,6.31925,0.3,300"))
			})
		})

		Context("when there are fields that cannot be parsed", func() {
			It("returns parse error", func() {
				_, err := parser.Parse("s,H&B MEGA VITAMIN TEENS,6.31925,0.3,300")
				Expect(err.Error()).To(ContainSubstring("parsing \"s\": invalid syntax"))
			})

			It("returns parse error", func() {
				_, err := parser.Parse("10072,H&B MEGA VITAMIN TEENS,h.31925,0.3,300")
				Expect(err.Error()).To(ContainSubstring("parsing \"h.31925\": invalid syntax"))
			})
		})
	})

	Describe("Parsefile", func() {
		var (
			tmp *os.File
			err error
		)

		BeforeEach(func() {
			tmp, err = ioutil.TempFile("", "test_data")
			Expect(err).NotTo(HaveOccurred())

			err = ioutil.WriteFile(tmp.Name(), []byte("10072,H&B MEGA VITAMIN TEENS,6.31925,0.3,300\n10073,H&C MEGA VITAMIN TEENS,6.31925,0.3,300"), 0755)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(os.Remove(tmp.Name())).To(Succeed())
		})

		It("can parse a file", func() {
			products, _ := parser.ParseFile(tmp.Name())
			Expect(products[0].Sku).To(Equal(10072))
			Expect(products[1].Volume).To(BeNumerically("==", 300))
		})

		Context("when the file does not exist", func() {
			It("return file non-exist error", func() {
				_, err := parser.ParseFile(tmp.Name() + "spiderman")
				Expect(err.Error()).To(ContainSubstring("open"))
			})
		})

		Context("when the file data is not valid", func() {
			It("return parsing error", func() {
				err = ioutil.WriteFile(tmp.Name(), []byte("s,H&B MEGA VITAMIN TEENS,6.31925,6.594,0.3,300\n10073,H&C MEGA VITAMIN TEENS,6.31925,6.594,0.3,300"), 0755)
				Expect(err).NotTo(HaveOccurred())
				_, err := parser.ParseFile(tmp.Name())
				Expect(err.Error()).To(ContainSubstring("invalid"))
			})
		})
	})
})
