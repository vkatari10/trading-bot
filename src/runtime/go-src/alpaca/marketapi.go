package alpaca

// This file works with the market API to get real time market data
//
// Author: Vikas Katari
// Date: 05/30/2025

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AlpacaQuote struct {
    Bars map[string]any `json:"bars"`
}

type AlpacaBar struct {
    Open float64 `json:"o"`
    High float64 `json:"h"`
    Low float64 `json:"l"`
    Close float64 `json:"c"`
    Volume float64 `json:"v"`
    //VWAP float64 `json:"vw"`
}   

/*
Alpaca Response
{
  "bars": {
    "AAPL": {
      "c": 211.02,
      "h": 211.105,
      "l": 211.02,
      "n": 17,
      "o": 211.105,
      "t": "2025-07-18T19:24:00Z",
      "v": 913,
      "vw": 211.051429
    }
  }
}
*/

// GetQuote returns the OHLCV bars of a given ticker (15m Delay)
func GetQuote(ticker string) ([5]float64, error) {

	url := alpacaMarketDataLink
    url += ticker

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Println(err)
    } // if

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

    var jsonMap map[string]any

    json.Unmarshal(body, &jsonMap) // CAN add goroutines here with mutex at each step

    //log.Println(jsonMap)

    quoteMap, ok := jsonMap["bars"].(map[string]any)
    if !ok {
        log.Println("ERROR:  Market JSON 1st parse failed")
        return [5]float64{}, nil
    } // if

    tickerMap, ok := quoteMap[ticker].(map[string]any)
    if !ok {
        log.Println("ERROR: Market JSON 2nd parse failed")
        return [5]float64{}, nil
    } // if

    // format we want of the array (YFinance format)
    bars := [5]string{"c", "h", "l", "o", "v"}
    finalBars := [5]float64{}

    for i := range bars {
        result, ok := tickerMap[bars[i]].(float64)
        if !ok  {
            log.Printf("ERROR: Could not get %s bar value\n", bars[i])
        } // if

        finalBars[i] = result
    } // for
    
    //log.Println(finalBars)
    return finalBars, nil
} // GetQuote

// GetAlpacaBars returns OHLCV data given a ticker as input
// expects valid API and SECRET Alpaca Keys
func GetAlpacaBars(ticker string) (*[5]float64, error) {
    url := alpacaMarketDataLink
    url += ticker
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("alpaca market external API: %w", err)
    } // if

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
    
    var jsonMap map[string]any

    json.Unmarshal(body, &jsonMap) // CAN add goroutines here with mutex at each step

    //log.Println(jsonMap)

    quoteMap, ok := jsonMap["bars"].(map[string]any)
    if !ok {
        return &[5]float64{}, fmt.Errorf("alpaca market JSON first parse failed")
    } // if

    tickerMap, ok := quoteMap[ticker].(map[string]any)
    if !ok {
        return &[5]float64{}, fmt.Errorf("alpaca market JSON second parse failed")
    } // if

    bars := [5]string{"o", "h", "l", "c", "v"}
    finalBars := [5]float64{}

    for i := range bars {
        val, ok := tickerMap[bars[i]].(float64)
        if !ok {
            return nil, fmt.Errorf("alpaca market came back in unexpected order")
        }

        finalBars[i] = val
    }

    return &finalBars, nil
} // GetAlpacaBars

