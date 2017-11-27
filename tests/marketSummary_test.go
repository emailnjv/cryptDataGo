package tests

import (
	"testing"
	"../bittrex"
)

func TestGetSummary(t *testing.T) {
	expected := true
	actual := bittrex.GetSummary("USDT-XRP")
	if actual.Success != true {
		t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual.Success)
	}
}
