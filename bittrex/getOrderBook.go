package bittrex

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
)

type (

	//OrderResult is the struct to the result field
	OrderResult struct {
		Buy  []BuyOrder  `json:"buy"`
		Sell []SellOrder `json:"sell"`
	}

	//OrderResponse contains all info returned in results
	OrderResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Result  OrderResult `json:"result"`
	}

	//Buy field contains all buy orders' quantity and price rate
	BuyOrder struct {
		Quantity float32 `json:"quantity"`
		Rate     float32 `json:"rate"`
	}

	//Sell field contains all sell orders' quantity and price rate
	SellOrder struct {
		Quantity float32 `json:"quantity"`
		Rate     float32 `json:"rate"`
	}
)

//GetOrderBook returns all of the orders contained in two groups buy's and sell's
func GetOrderBook(marketSym string) OrderResponse {

	//Cretes HTTP client, followed by setting headers, and firing
	var myClient= &http.Client{}
	queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getorderbook?market=", string(marketSym), "&type=both"}, "")
	req, err := http.NewRequest("GET", queryURL, nil)
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	response, err := myClient.Do(req)
	var orderResp OrderResponse
	if err != nil {
		log.Fatal(err)
	}

	//Read, and close response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	response.Body.Close()

	//Unmarshall JSON
	errr := json.Unmarshal(body, &orderResp)
	if errr != nil {
		println(errr)
	}

	//Return response
	return orderResp
}