package bittrex

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
)

type (

	//MHistory contains ID, Timestamp, Quantity, Price, Total, Filltype, OrderType
 	MHistory struct {
 		ID        float32 `json:"id"`
 		TimeStamp string  `json:"timeStamp"`
 		Quantity  float32 `json:"quantity"`
 		Price     float32 `json:"price"`
 		Total     float32 `json:"total"`
 		FillType  string  `json:"fillType"`
 		OrderType string  `json:"orderType"`
 	}

	//HistoryResponse is basic struct for history api response
 	HistoryResponse struct {
 		Success bool       `json:"success"`
 		Message string     `json:"message"`
 		Result  []MHistory `json:"result"`
 	}
)

//GetMarketHistory returns specified market history
//Takes in market spits out response struct
func GetMarketHistory(marketSym string) HistoryResponse {

	//Creates HTTP client, set headers, and fire request
	var myClient= &http.Client{}
	queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getmarkethistory?market=", string(marketSym)}, "")
	req, err := http.NewRequest("GET", queryURL, nil)
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	response, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//Read response and un-marshall into struct
	var marketHist HistoryResponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	errr := json.Unmarshal(body, &marketHist)
	if errr != nil {
		println(errr)
	}

	response.Body.Close()

	//Return response struct
	return marketHist

}
