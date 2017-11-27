package tests

import (
	"testing"
	"../bittrex"
)

func TestGetAllMarkets(t *testing.T) {
	expected := true
	actual := bittrex.GetMakets()
	if actual.Success != true {
		t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual.Success)
	}
}

