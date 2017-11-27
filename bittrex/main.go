package bittrex

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"
// 	"sync"
// )

// type (
// 	ProcessedCurrency struct {
// 		Name           string      `bson:"name"`
// 		Base           string      `bson:"base"`
// 		BaseCurrency   string      `bson:"baseCurrency"`
// 		MarketCurrency string      `bson:"marketCurrency"`
// 		BuyOrders      []BuyOrder  `bson:"buyOrders"`
// 		SellOrders     []SellOrder `bson:"sellOrders"`
// 		History        []MHistory  `bson:"history"`
// 		Volume         float32     `bson:"volume"`
// 		BaseVolume     float32     `bson:"baseVolume"`
// 		Time           string      `bson:"time"`
// 		Price          float32     `bson:"price"`
// 	}

// 	MHistory struct {
// 		ID        float32 `json:"id"`
// 		TimeStamp string  `json:"timeStamp"`
// 		Quantity  float32 `json:"quantity"`
// 		Price     float32 `json:"price"`
// 		Total     float32 `json:"total"`
// 		FillType  string  `json:"fillType"`
// 		OrderType string  `json:"orderType"`
// 	}

// 	historyResponse struct {
// 		Success bool       `json:"success"`
// 		Message string     `json:"message"`
// 		Result  []MHistory `json:"result"`
// 	}

// 	// mSummary struct {
// 	// 	MarketName     string  `json:"marketName"`
// 	// 	High           float32 `json:"high"`
// 	// 	Low            float32 `json:"low"`
// 	// 	Volume         float32 `json:"volume"`
// 	// 	Last           float32 `json:"last"`
// 	// 	BaseVolume     float32 `json:"baseVolume"`
// 	// 	TimeStamp      string  `json:"timeStamp"`
// 	// 	Bid            float32 `json:"bid"`
// 	// 	Ask            float32 `json:"ask"`
// 	// 	OpenBuyOrders  float32 `json:"openBuyOrders"`
// 	// 	OpenSellOrders float32 `json:"openSellOrders"`
// 	// 	PrevDay        float32 `json:"prevDay"`
// 	// 	Created        string  `json:"created"`
// 	// }

// 	// summaryResponse struct {
// 	// 	Success bool       `json:"success"`
// 	// 	Message string     `json:"message"`
// 	// 	Result  []mSummary `json:"result"`
// 	// }

// 	currency struct {
// 		Symbol string  `json:"symbol"`
// 		Used   float32 `json:"used"`
// 	}

// 	BuyOrder struct {
// 		Quantity float32 `json:"quantity"`
// 		Rate     float32 `json:"rate"`
// 	}

// 	SellOrder struct {
// 		Quantity float32 `json:"quantity"`
// 		Rate     float32 `json:"rate"`
// 	}

// 	OrderResult struct {
// 		Buy  []BuyOrder  `json:"buy"`
// 		Sell []SellOrder `json:"sell"`
// 	}

// 	orderResponse struct {
// 		Success bool        `json:"success"`
// 		Message string      `json:"message"`
// 		Result  OrderResult `json:"result"`
// 	}

// 	// marketRespone struct {
// 	// 	Success bool     `json:"success"`
// 	// 	Message string   `json:"message"`
// 	// 	Result  []market `json:"result"`
// 	// }

// 	// market struct {
// 	// 	MarketCurrency     string  `json:"marketCurrency"`
// 	// 	BaseCurrency       string  `json:"baseCurrency"`
// 	// 	MarketCurrencyLong string  `json:"marketCurrencyLong"`
// 	// 	BaseCurrencyLong   string  `json:"baseCurrencyLong"`
// 	// 	MinTradeSize       float32 `json:"minTradeSize"`
// 	// 	MarketName         string  `json:"marketName"`
// 	// 	IsActive           bool    `json:"isActive"`
// 	// 	Created            string  `json:"created"`
// 	// 	Notice             string  `json:"notice"`
// 	// 	IsSponsored        bool    `json:"isSponsored"`
// 	// 	LogoURL            string  `json:"logoURL"`
// 	// }
// )

// // Entry Point
// func GetCurrencies() chan ProcessedCurrency {
// 	var myClient = &http.Client{}
// 	req, err := http.NewRequest("GET", "https://www.bittrex.com/api/v1.1/public/getmarkets", nil)
// 	req.Close = true
// 	req.Header.Add("Content-Type", "application/json")
// 	response, err := myClient.Do(req)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var mResponse marketRespone
// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	errr := json.Unmarshal(body, &mResponse)
// 	if errr != nil {
// 		println(errr)
// 	}

// 	var mainWg sync.WaitGroup
// 	results := mResponse.Result
// 	mainWg.Add(len(results))

// 	out := make(chan ProcessedCurrency, len(results))

// 	go func(in *sync.WaitGroup) {

// 		in.Wait()
// 		close(out)
// 	}(&mainWg)
// 	for _, markets := range results {

// 		var wg sync.WaitGroup
// 		summary := make(chan summaryResponse, 1)
// 		orderBook := make(chan orderResponse, 1)
// 		marketHistory := make(chan historyResponse, 1)
// 		wg.Add(3)

// 		go func(marketSym string) {
// 			var myClient = &http.Client{}
// 			queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getmarketsummary?market=", string(marketSym)}, "")
// 			req, err := http.NewRequest("GET", queryURL, nil)
// 			req.Close = true
// 			req.Header.Add("Content-Type", "application/json")
// 			response, err := myClient.Do(req)
// 			if err != nil {
// 				log.Fatal(err)
// 			} else {

// 				var sumResponse summaryResponse
// 				defer response.Body.Close()
// 				body, err := ioutil.ReadAll(response.Body)
// 				if err != nil {
// 					panic(err.Error())
// 				}
// 				errr := json.Unmarshal(body, &sumResponse)
// 				if errr != nil {
// 					println(errr)
// 				}
// 				summary <- sumResponse
// 				wg.Done()

// 			}
// 		}(markets.MarketName)
// 		go func(marketSym string) {
// 			var myClient = &http.Client{}
// 			queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getorderbook?market=", string(marketSym), "&type=both"}, "")
// 			req, err := http.NewRequest("GET", queryURL, nil)
// 			req.Close = true
// 			req.Header.Add("Content-Type", "application/json")
// 			response, err := myClient.Do(req)
// 			if err != nil {
// 				log.Fatal(err)
// 			} else {
// 				var orderResp orderResponse
// 				defer response.Body.Close()
// 				body, err := ioutil.ReadAll(response.Body)
// 				if err != nil {
// 					panic(err.Error())
// 				}
// 				errr := json.Unmarshal(body, &orderResp)
// 				if errr != nil {
// 					println(errr)
// 				}
// 				orderBook <- orderResp
// 				wg.Done()

// 			}
// 		}(markets.MarketName)
// 		go func(marketSym string) {
// 			var myClient = &http.Client{}
// 			queryURL := strings.Join([]string{"https://www.bittrex.com/api/v1.1/public/getmarkethistory?market=", string(marketSym)}, "")
// 			req, err := http.NewRequest("GET", queryURL, nil)
// 			req.Close = true
// 			req.Header.Add("Content-Type", "application/json")
// 			response, err := myClient.Do(req)
// 			if err != nil {
// 				log.Fatal(err)
// 			} else {
// 				var marketHist historyResponse
// 				defer response.Body.Close()
// 				body, err := ioutil.ReadAll(response.Body)
// 				if err != nil {
// 					panic(err.Error())
// 				}
// 				errr := json.Unmarshal(body, &marketHist)
// 				if errr != nil {
// 					println(errr)
// 				}
// 				marketHistory <- marketHist
// 				wg.Done()

// 			}
// 		}(markets.MarketName)
// 		go func(marketSym market) {

// 			var summ summaryResponse
// 			var orde orderResponse
// 			var mark historyResponse

// 			for result := range summary {
// 				summ = result
// 				close(summary)
// 			}
// 			for result := range orderBook {
// 				orde = result
// 				close(orderBook)
// 			}
// 			for result := range marketHistory {
// 				mark = result
// 				close(marketHistory)
// 			}
// 			wg.Wait()
// 			parsedResult := ProcessedCurrency{marketSym.MarketCurrencyLong, marketSym.BaseCurrencyLong, marketSym.BaseCurrency, marketSym.MarketCurrency, orde.Result.Buy, orde.Result.Sell, mark.Result, summ.Result[0].Volume, summ.Result[0].BaseVolume, summ.Result[0].TimeStamp, summ.Result[0].Last}

// 			out <- parsedResult
// 			mainWg.Done()

// 		}(markets)

// 	}

// 	return out

// }
