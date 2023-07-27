package gateways_test

import (
	"CurrencyExchangeService/gateways"
	"fmt"
	"testing"
)

func TestGateway(t *testing.T) {
	result := gateways.GetOpenExchangeRate()
	fmt.Printf("Got Result:  %+v", result)
}
