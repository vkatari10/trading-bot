package eventloop

// This file contains methods to help assist the main eventloop

import (
	"time"
	"fmt"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	"math/rand"
)

//BurnIn Loads the Burn in Data to intialize technical indicators
func BurnIn(burnTime int, ticker string, refresh int) (arr []float64, finalQuote [5]float64) {
	go SendPayload(map[string]any{
		"msg": "STAGE: BURN IN",
		}, logLink)

	stopTime := time.Duration(refresh)

	// stores burn data
	burn := []float64{}

	// stores latest quotes
	var newQuote [5]float64

	for i := range burnTime {
		newQuote, err := api.GetQuote(ticker)
		if err != nil {
			go SendPayload(map[string]any {
			"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // if
		burn = append(burn, newQuote[0])
	
		go SendPayload(map[string]any {
			"msg": fmt.Sprintf("QUOTE: $%.2f (%d / %d) burned in", newQuote[0], i + 1, burnTime),
		}, logLink)
		
		time.Sleep(stopTime * time.Second)
	} // for
	//fmt.Println(burn) check on faster time 

	return burn, newQuote
} // BurnIn

// handlePrediction handles the prediction made by the ML model by 
// working with the broker API
func handlePrediction(apiBuffer *APIBuffer, prediction float64, ticker string) {

	decisionMsg := "DECIDE: "
		var decision string;

		if prediction > 0 { // buy
			decision = "buy"
			go api.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "BUY"
		} else if prediction < 0 { // sell
			decision = "sell"
			go api.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "SELL"
		} else {
			decisionMsg += "HOLD"
		} // if-else
		
		go apiBuffer.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: %s %s", decisionMsg, ticker),
			}, logLink)
		go sendBrokerData()

} // handlePrediction

// overrideBurnIn overides the burn in by creating random values of the specified window size 
func overrideBurnIn(windowSize int) []float64 {
	result := make([]float64, 0)
	for i := 0; i < windowSize; i++ {
		result = append(result, rand.Float64() * 10)
	} // for
	return result
} // bypassBurnIn