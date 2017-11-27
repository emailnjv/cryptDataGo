package tests

import (
	"testing"
	"../bittrex"
)

func TestGetMarketHistory(t *testing.T) {
	expected := true
	actual := bittrex.GetOrderBook("USDT-XRP")
	if actual.Success != true {
		t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual.Success)
	}
}