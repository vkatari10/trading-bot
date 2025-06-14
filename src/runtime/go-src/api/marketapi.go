package api

// This file works with the market API to get real time market data

import (
	"net/http"
	"io"
    "encoding/json"
    "log"
)

// GetQuote returns the OHLCV bars of a given ticker (15m Delay)
func GetQuote(ticker string) ([5]float64, error) {

	url := "https://data.alpaca.markets/v2/stocks/bars/latest?symbols="
    url += ticker

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", getAlpacaKey())
	req.Header.Add("APCA-API-SECRET-KEY", getAlpacaSecret())

	res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Println(err)
    }

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