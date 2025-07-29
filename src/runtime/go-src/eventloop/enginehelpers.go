package eventloop

// This file contains methods to help assist the main eventloop

// THIS FILE IS DEPRECATED
// MOVE handlePrediction to the RISK module

import (
	"fmt"
	alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
	apibuf "github.com/vkatari10/trading-bot/src/runtime/go-src/apibuffer"
)

// TODO: Move to risk package once rewrite done
// handlePrediction handles the prediction made by the ML model by 
// working with the broker API
func handlePrediction(
	apiBuffer *apibuf.APIBuffer, 
	prediction float64, 
	ticker string,
	postLink string,
	) {

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
			}, postLink)
		// go sendBrokerData()

} // handlePrediction

