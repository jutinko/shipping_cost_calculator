package currency_converter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCurrencyConverter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CurrencyConverter Suite")
}
