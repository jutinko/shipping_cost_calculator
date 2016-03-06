package product_store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProductStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductStore Suite")
}
