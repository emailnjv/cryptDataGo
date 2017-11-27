package bittrex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type (

	// MSummary is the basic structure for the summary api call
	MSummary struct {
		MarketName     string  `json:"marketName"`
		High           float32 `json:"high"`
		Low            float32 `json:"low"`
		Volume         float32 `json:"volume"`
		Last           float32 `json:"last"`
		BaseVolume     float32 `json:"baseVolume"`
		TimeStamp      string  `json:"timeStamp"`
		Bid            float32 `json:"bid"`
		Ask            float32 `json:"ask"`
		OpenBuyOrders  float32 `json:"openBuyOrders"`
		OpenSellOrders float32 `json:"openSellOrders"`
		PrevDay        float32 `json:"prevDay"`
		Created        string  `json:"created"`
	}

	SummaryResponse struct {
		Success bool       `json:"success"`
		Message string     `json:"message"`
		Result  []MSummary `json:"result"`
	}
)

// GetSummary returns the summary of the plugged in market
// takes in the market symbol ex. USDT-XRP, and returns the response struct
func GetSummary(marketSym string) SummaryResponse {

	// Create HTTP client, and perform get request to api endpoint
	var myClient = &http.Client{}
	queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getmarketsummary?market=", string(marketSym)}, "")
	req, err := http.NewRequest("GET", queryURL, nil)
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	response, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read the result and unmarshal to struct
	var sumResponse SummaryResponse
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	errr := json.Unmarshal(body, &sumResponse)
	if errr != nil {
		println(errr)
	}

	return sumResponse
}
