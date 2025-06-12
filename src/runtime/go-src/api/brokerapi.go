package api

// External API methods to place buy/sell orders on brokerage

import (
	"io"
	"net/http"
    "fmt"
	"strings"
	"encoding/json"
	"strconv"
)

// GetCashValue gets the brokerage account value of cash and 
// account total value
func GetCashValue() (cashAvail float64, totalValue float64, err error) {
	url := "https://paper-api.alpaca.markets/v2/account"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data map[string]any;

	json.Unmarshal(body, &data)

	cash, err := getDoubleValue(data, "cash")
	if err != nil {
		return 0, 0, fmt.Errorf("%v", err)
	} // if

	value, err := getDoubleValue(data, "portfolio_value")
	if err != nil {
		return 0, 0, fmt.Errorf("%v", err)
	} // if

	return cash, value, nil
} // GetCashValue

// GetPositions gets the current open positons in the broker account
// Right now implemented to only do the first position since the bot only
// can track one stock at the time
func GetPositions() (quantity float64, averageCost float64, marketValue float64, err error) {

	url := "https://paper-api.alpaca.markets/v2/positions"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data []map[string]any;

	json.Unmarshal(body, &data)

	// this is an array

	qty, err := getDoubleValue(data[0], "qty")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	} // if

	avgCost, err := getDoubleValue(data[0], "avg_entry_price")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	} // if

	value, err := getDoubleValue(data[0], "market_value")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	} // if

	return qty, avgCost, value, nil 

} // GetPositions


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

// getDoubleValue helps type assert and convert JSON values 
// back to float64 values
func getDoubleValue(json map[string]any, key string) (result float64, err error) {
	strResult, ok := json[key].(string)
	if !ok {
		return 0, fmt.Errorf("%v", err)
	} // if

	result, err = strconv.ParseFloat(strResult, 64)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	} // if
	
	return result, nil
} // getDoubleValue

// FUTURE --> Add these functions
/*
Positions
ssss1. get all positions
2. close all positions
3. get position
4. close position
4. close position by shares
*/
