package utilities_test

import (
	. "github.com/jutinko/shipping_cost_calculator/utilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvReader", func() {
	var (
		validJson []byte
		envReader *EnvReader
	)

	BeforeEach(func() {
		validJson = []byte(`{"p-redis":[{"name":"product-info","label":"p-redis","tags":["pivotal","redis"],"plan":"shared-vm","credentials":{"host":"192.168","password":"iiiiii8c-bc18-j5d2-b728-d83ef331e72f","port":40592}}]}`)
		envReader = NewEnvReader(validJson)
	})

	It("should return the host", func() {
		Expect(envReader.GetHost()).To(Equal("192.168"))
	})

	It("should return the password", func() {
		Expect(envReader.GetPassword()).To(Equal("iiiiii8c-bc18-j5d2-b728-d83ef331e72f"))
	})

	It("should return the port", func() {
		Expect(envReader.GetPort()).To(Equal("40592"))
	})
})
