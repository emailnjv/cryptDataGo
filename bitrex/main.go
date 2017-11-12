package bitrex

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"sync"
	"strings"

	"fmt"
)

type (

	ProcessedCurrency struct {
		Name string
		Base string
		BaseCurrency string
		MarketCurrency string
		BuyOrders []BuyOrder
		SellOrders []SellOrder
		History []mHistory
		Volume float32
		BaseVolume float32
		Time string
		Price float32
	}

	mHistory struct {
		Id float32 `json:"id"`
		TimeStamp string `json:"timeStamp"`
		Quantity float32 `json:"quantity"`
		Price float32 `json:"price"`
		Total float32 `json:"total"`
		FillType string `json:"fillType"`
		OrderType string `json:"orderType"`
	}

	historyResponse struct {
		Success bool `json:"success"`
		Message string `json:"message"`
		Result []mHistory `json:"result"`
	}

	mSummary struct {
		MarketName string `json:"marketName"`
		High float32 `json:"high"`
		Low float32 `json:"low"`
		Volume float32 `json:"volume"`
		Last float32 `json:"last"`
		BaseVolume float32 `json:"baseVolume"`
		TimeStamp string `json:"timeStamp"`
		Bid float32 `json:"bid"`
		Ask float32 `json:"ask"`
		OpenBuyOrders float32 `json:"openBuyOrders"`
		OpenSellOrders float32 `json:"openSellOrders"`
		PrevDay float32 `json:"prevDay"`
		Created string `json:"created"`

	}

	summaryResponse struct {
		Success bool `json:"success"`
		Message string `json:"message"`
		Result []mSummary `json:"result"`
	}

	currency struct{
		Symbol string
		Used float32
	}

	BuyOrder struct {
		Quantity float32 `json:"quantity"`
		Rate float32 `json:"rate"`
	}

	SellOrder struct {
		Quantity float32 `json:"quantity"`
		Rate float32 `json:"rate"`
	}

	orderResult struct {
		Buy []BuyOrder
		Sell []SellOrder
	}

	orderResponse struct {
		Success bool `json:"success"`
		Message string `json:"message"`
		Result orderResult `json:"result"`
	}

	marketRespone struct {
		Success bool `json:"success"`
		Message string `json:"message"`
		Result []market `json:"result"`
	}

	market struct {
		MarketCurrency    string `json:"marketCurrency"`
		BaseCurrency    string `json:"baseCurrency"`
		MarketCurrencyLong    string `json:"marketCurrencyLong"`
		BaseCurrencyLong    string `json:"baseCurrencyLong"`
		MinTradeSize    float32 `json:"minTradeSize"`
		MarketName    string `json:"marketName"`
		IsActive    bool `json:"isActive"`
		Created    string `json:"created"`
		Notice    string `json:"notice"`
		IsSponsored    bool `json:"isSponsored"`
		LogoUrl    string `json:"logoUrl"`

	}

)



func getInfo(results []market) []ProcessedCurrency{
	var wg sync.WaitGroup
	var mainWg sync.WaitGroup
	out := make(chan ProcessedCurrency)
	//final := make(chan []ProcessedCurrency)

	mainWg.Add(len(results))
	for _, markets := range(results) {
		summary := make(chan summaryResponse)
		orderBook := make(chan orderResponse)
		marketHistory := make(chan historyResponse)
		wg.Add(3)
		go func(marketSym string) {
			var myClient = &http.Client{}
			queryUrl := strings.Join([]string{"https://bittrex.com/api/v1.1/public/getmarketsummary?market=",string(marketSym)}, "")
			req, err := http.NewRequest("GET", queryUrl, nil)
			req.Header.Add("Content-Type", "application/json")
			response, err := myClient.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {

				var sumResponse summaryResponse
				defer response.Body.Close()
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					panic(err.Error())
				}
				errr := json.Unmarshal(body, &sumResponse)
				if errr != nil{
					println(errr)
				}
				summary <- sumResponse
				wg.Done()

			}
		}(markets.MarketName)

		go func(marketSym string) {
			var myClient = &http.Client{}
			queryUrl := strings.Join([]string{"https://bittrex.com/api/v1.1/public/getorderbook?market=",string(marketSym),"&type=both"}, "")
			req, err := http.NewRequest("GET", queryUrl, nil)
			req.Header.Add("Content-Type", "application/json")
			response, err := myClient.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {
				var orderResp orderResponse
				defer response.Body.Close()
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					panic(err.Error())
				}
				errr := json.Unmarshal(body, &orderResp)
				if errr != nil{
					println(errr)
				}
				orderBook <- orderResp
				wg.Done()

			}
		}(markets.MarketName)

		go func(marketSym string) {
			var myClient = &http.Client{}
			queryUrl := strings.Join([]string{"https://bittrex.com/api/v1.1/public/getmarkethistory?market=",string(marketSym)}, "")
			req, err := http.NewRequest("GET", queryUrl, nil)
			req.Header.Add("Content-Type", "application/json")
			response, err := myClient.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {
				var marketHist historyResponse
				defer response.Body.Close()
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					panic(err.Error())
				}
				errr := json.Unmarshal(body, &marketHist)
				if errr != nil{
					println(errr)
				}
				marketHistory <- marketHist
				wg.Done()

			}
		}(markets.MarketName)
		go func(marketSym market) {
			wg.Wait()
			summ := <- summary
			close(summary)
			orde := <- orderBook
			close(orderBook)
			mark := <- marketHistory
			close(marketHistory)
			parsedResult := ProcessedCurrency{marketSym.MarketCurrencyLong, marketSym.BaseCurrencyLong, marketSym.BaseCurrency, marketSym.MarketCurrency, orde.Result.Buy, orde.Result.Sell, mark.Result, summ.Result[0].Volume, summ.Result[0].BaseVolume, summ.Result[0].TimeStamp, summ.Result[0].Last}

			out <- parsedResult
			mainWg.Done()
		}(markets)

	}
	var finalArr []ProcessedCurrency
	for result := range out{
		finalArr = append(finalArr, result)
	}

	go func() {
		mainWg.Wait()
		close(out)
	}()



	return finalArr


}


func GetCurrencies() {
	var myClient = &http.Client{}

	req, err := http.NewRequest("GET", "https://bittrex.com/api/v1.1/public/getmarkets", nil)
	req.Header.Add("Content-Type", "application/json")
	response, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		var mResponse marketRespone
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}
		errr := json.Unmarshal(body, &mResponse)
		if errr != nil{
			println(errr)
		}

		//var returnArray []ProcessedCurrency
		for _, result := range getInfo(mResponse.Result) {
			fmt.Printf("%T	\n",result)
		}





	}
}