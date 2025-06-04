package api

// This file works with the market API to get real time market data

import (
	"net/http"
	"io"
    "encoding/json"
    "log"
)

// GetQuote Gets the current price of a given ticker using the 
// Alpaca API
func GetQuote(ticker string) (float64, error) {

	url := "https://data.alpaca.markets/v2/stocks/quotes/latest?symbols="

    url += ticker

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", "PK6VHZZLDHZSNSV1ISJR")
	req.Header.Add("APCA-API-SECRET-KEY", "GWFHZ3FRlXsgHvblS5NQMK1oLbFhY660aMLMcn7A")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

    var jsonMap map[string]any

    json.Unmarshal(body, &jsonMap) // CAN add goroutines here with mutex at each step

    quoteMap, ok := jsonMap["quotes"].(map[string]any)
    if !ok {
        log.Println("ERROR:  Market JSON 1st parse failed")
    } // if

    tickerMap, ok := quoteMap[ticker].(map[string]any)
    if !ok {
        log.Println("ERROR: Market JSON 2nd parse failed")
    } // if

    result, ok := tickerMap["ap"].(float64)
    if !ok {
        log.Println("ERROR: Market JSON 3rd parse failed")
    } // if

    return result, nil
} // GetQuote