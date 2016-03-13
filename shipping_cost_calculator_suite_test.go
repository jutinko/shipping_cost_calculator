package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShippingCostCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShippingCostCalculator Suite")
}
