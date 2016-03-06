// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/jutinko/shipping_cost_calculator/calculator"
)

type FakeCurrencyConverter struct {
	ExchangeStub        func(float64) float64
	exchangeMutex       sync.RWMutex
	exchangeArgsForCall []struct {
		arg1 float64
	}
	exchangeReturns struct {
		result1 float64
	}
}

func (fake *FakeCurrencyConverter) Exchange(arg1 float64) float64 {
	fake.exchangeMutex.Lock()
	fake.exchangeArgsForCall = append(fake.exchangeArgsForCall, struct {
		arg1 float64
	}{arg1})
	fake.exchangeMutex.Unlock()
	if fake.ExchangeStub != nil {
		return fake.ExchangeStub(arg1)
	} else {
		return fake.exchangeReturns.result1
	}
}

func (fake *FakeCurrencyConverter) ExchangeCallCount() int {
	fake.exchangeMutex.RLock()
	defer fake.exchangeMutex.RUnlock()
	return len(fake.exchangeArgsForCall)
}

func (fake *FakeCurrencyConverter) ExchangeArgsForCall(i int) float64 {
	fake.exchangeMutex.RLock()
	defer fake.exchangeMutex.RUnlock()
	return fake.exchangeArgsForCall[i].arg1
}

func (fake *FakeCurrencyConverter) ExchangeReturns(result1 float64) {
	fake.ExchangeStub = nil
	fake.exchangeReturns = struct {
		result1 float64
	}{result1}
}

var _ calculator.CurrencyConverter = new(FakeCurrencyConverter)