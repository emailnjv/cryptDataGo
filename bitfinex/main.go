package bitfinex

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"sync"
	"fmt"
)

type (

	BaseCurrencies struct {
		BaseCurrency []Currency
	}



	Currency struct{
		Symbol string
		Used int
	}

	Market struct {
		Pair    string `json:"pair"`
		Price_precision    int `json:"price_precision"`
		Initial_margin    string `json:"initial_margin"`
		Minimum_margin    string `json:"minimum_margin"`
		Maximum_order_size    string `json:"maximum_order_size"`
		Minimum_order_size    string `json:"minimum_order_size"`
		Expiration    string `json:"expiration"`

	}

)
var baseCurrencies []Currency


func stringInArray(currencySym string, array []Currency) []int {
	for i, c := range array {

		if c.Symbol == currencySym {
			answer := []int{1,i}
			return answer
		}
	}
	answer := []int{0,0}
	return answer
}
func symbolParse(mark Market, wg *sync.WaitGroup){
	symEnd := mark.Pair[3:]
	indexCheck := stringInArray(symEnd, baseCurrencies)
	if indexCheck[0] == 1 {
		baseCurrencies[indexCheck[1]].Used += 1
		wg.Done()
	}else {
		newCurrency := Currency{symEnd,1}
		baseCurrencies = append(baseCurrencies, newCurrency)
		wg.Done()
	}
}


func GetCurrencies() {
	//var currencyList BaseCurrencies
	var myClient = &http.Client{Timeout: 10 * time.Second}

	response, err := myClient.Get("https://api.bitfinex.com/v1/symbols_details")
	if err != nil {
		log.Fatal(err)
	} else {
		var markets []Market
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}
		errr := json.Unmarshal(body, &markets)
		if errr != nil{
			println(errr)
		}


		var wg sync.WaitGroup
		wg.Add(len(markets))

		for _, mark := range markets{
			go symbolParse(mark, &wg)
		}
			wg.Wait()
			for _, q := range baseCurrencies{
				fmt.Printf("%v", q)			}


	}
}
