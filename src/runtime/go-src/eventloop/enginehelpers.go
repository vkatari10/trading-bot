package eventloop

// This file contains methods to help assist the main eventloop

import (
	"time"
	"fmt"
	alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
	"math/rand"
	"os"
	apibuf "github.com/vkatari10/trading-bot/src/runtime/go-src/apibuffer"
)

// DEPRECATED
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
		newQuote, err := alpaca.GetQuote(ticker)
		if err != nil {
			go SendPayload(map[string]any {
			"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // if
		burn = append(burn, newQuote[0])
	
		go SendPayload(map[string]any {
			"msg": fmt.Sprintf("%s: $%.2f (%d / %d) burned in", ticker, newQuote[0], i + 1, burnTime),
		}, logLink)
		
		time.Sleep(stopTime * time.Second)
	} // for
	//fmt.Println(burn) check on faster time 

	return burn, newQuote
} // BurnIn

// TODO: Move to risk package once rewrite done
// handlePrediction handles the prediction made by the ML model by 
// working with the broker API
func handlePrediction(apiBuffer *apibuf.APIBuffer, prediction float64, ticker string) {

	decisionMsg := "DECIDE: "
		var decision string;

		if prediction > 0 { // buy
			decision = "buy"
			go alpaca.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "BUY"
		} else if prediction < 0 { // sell
			decision = "sell"
			go alpaca.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "SELL"
		} else {
			decisionMsg += "HOLD"
		} // if-else
		
		go apiBuffer.Enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: %s %s", decisionMsg, ticker),
			}, logLink)
		go sendBrokerData()

} // handlePrediction

// DEPRECATED
// overrideBurnIn overides the burn in by creating random values of the specified window size 
func overrideBurnIn(windowSize int) []float64 {
	result := make([]float64, 0)
	for i := 0; i < windowSize; i++ {
		result = append(result, rand.Float64() * 10)
	} // for
	return result
} // bypassBurnIn

// getFileName gets the config file name used from the CLI
func getFileName() string {
	argv := os.Args
	return argv[1]
} // getFileName