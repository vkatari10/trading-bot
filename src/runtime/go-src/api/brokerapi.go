package api

// External API methods to place buy/sell orders on brokerage

import (
	"io"
	"net/http"
    "fmt"
	"strings"
)

// Returns aspects of the Brokerage Account probably to use in API
func Acct(data string) any {
	url := "https://paper-api.alpaca.markets/v2/account"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

    return body

} // Acct

// PlaceMarketOrder places an "buy" or "sell" signal based on the 
// input string
func PlaceMarketOrder(ticker string, shares int, side string) {
	url := "https://paper-api.alpaca.markets/v2/orders"

	jsonString := fmt.Sprintf("{\"type\":\"market\",\"time_in_force\":\"day\",\"symbol\":\"%s\",\"qty\":\"%d\",\"side\":\"%s\"}", ticker, shares, side)

	//fmt.Println(jsonString) // DEBUG for payload

	payload := strings.NewReader(jsonString)

	req, _ := http.NewRequest("POST", url, payload)
			
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
} // PlaceMarketOrder



// FUTURE --> Add these functions
/*
Positions
ssss1. get all positions
2. close all positions
3. get position
4. close position
4. close position by shares
*/
