package main

import (
	"./bitrex"
)

type (


	//ProcessedCurrency struct {
	//	Market string
	//	BaseCurrency string
	//	MarketCurrency string
	//	BuyOrders []BuyOrder
	//	SellOrders []SellOrder
	//	Volume int
	//}

)


func main() {
	bitrex.GetCurrencies()
}
