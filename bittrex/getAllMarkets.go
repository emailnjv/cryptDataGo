package bittrex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	MarketRespone struct {
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Result  []Market `json:"result"`
	}

	// Market is the basic response result for all markets from the exchange API
	// Accepts no parameters, and returns a market response struct
	Market struct {
		MarketCurrency     string  `json:"marketCurrency"`
		BaseCurrency       string  `json:"baseCurrency"`
		MarketCurrencyLong string  `json:"marketCurrencyLong"`
		BaseCurrencyLong   string  `json:"baseCurrencyLong"`
		MinTradeSize       float32 `json:"minTradeSize"`
		MarketName         string  `json:"marketName"`
		IsActive           bool    `json:"isActive"`
		Created            string  `json:"created"`
		Notice             string  `json:"notice"`
		IsSponsored        bool    `json:"isSponsored"`
		LogoURL            string  `json:"logoURL"`
	}
)

// GetMakets function to get all markets
func GetMakets() MarketRespone {

	// Create HTTP client for request
	var myClient = &http.Client{}
	req, err := http.NewRequest("GET", "https://www.bittrex.com/api/v1.1/public/getmarkets", nil)
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	// Execute request
	response, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// create response object
	var mResponse MarketRespone

	// write response to object
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// Unmarshall response, followed by closing the responses
	errr := json.Unmarshal(body, &mResponse)
	if errr != nil {
		println(errr)
	}

	// Close the response
	response.Body.Close()

	// Return result
	return mResponse
}
