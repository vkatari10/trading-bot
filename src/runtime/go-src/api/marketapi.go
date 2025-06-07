package api

// This file works with the market API to get real time market data

import (
	"net/http"
	"io"
    "encoding/json"
    "log"
)

// GetQuote returns the OHLCV bars of a given ticker (15m Delay)
func GetQuote(ticker string, barType string) ([5]float64, error) {

	url := "https://data.alpaca.markets/v2/stocks/bars/latest?symbols="
    url += ticker

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

    var jsonMap map[string]any

    json.Unmarshal(body, &jsonMap) // CAN add goroutines here with mutex at each step

    quoteMap, ok := jsonMap["bars"].(map[string]any)
    if !ok {
        log.Println("ERROR:  Market JSON 1st parse failed")
    } // if

    tickerMap, ok := quoteMap[ticker].(map[string]any)
    if !ok {
        log.Println("ERROR: Market JSON 2nd parse failed")
    } // if

    bars := [5]string{"o", "h", "c", "l", "v"}
    finalBars := [5]float64{}

    for i := range bars {
        result, ok := tickerMap[bars[i]].(float64)
        if !ok  {
            log.Printf("ERROR: Could not get %c bar value\n", bars[i])
        } // if

        finalBars[i] = result
    } // for
    
    //log.Println(finalBars)
    return finalBars, nil
} // GetQuote