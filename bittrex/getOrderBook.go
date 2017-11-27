package bittrex

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
)

type (
	OrderResult struct {
		Buy  []BuyOrder  `json:"buy"`
		Sell []SellOrder `json:"sell"`
	}

	OrderResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Result  OrderResult `json:"result"`
	}

	BuyOrder struct {
		Quantity float32 `json:"quantity"`
		Rate     float32 `json:"rate"`
	}

	SellOrder struct {
		Quantity float32 `json:"quantity"`
		Rate     float32 `json:"rate"`
	}
)

func GetOrderBook(marketSym string) OrderResponse {

	var myClient= &http.Client{}
	queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getorderbook?market=", string(marketSym), "&type=both"}, "")
	req, err := http.NewRequest("GET", queryURL, nil)
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	response, err := myClient.Do(req)
	var orderResp OrderResponse

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}
		errr := json.Unmarshal(body, &orderResp)
		if errr != nil {
			println(errr)
		}

	}

	return orderResp
}