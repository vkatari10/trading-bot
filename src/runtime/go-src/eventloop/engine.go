package eventloop

import (
	"time"
	"runtime"
	"fmt"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

var (
	us engine.RuntimeSettings // to use its values in the rest of the package
)

// tempTicker is how we will integrate specific tickers onto the eventloop
func EventLoop(tempTicker string) { 

	
	// Load User JSON -> Convert to Go Struct

	// TODO create a method to just return the map 
	// and then send the map ONLY to the parsers
	userConfigFile := "../../../" + getFileName()

	userSettings, err := engine.GetRuntimeSettings(userConfigFile)
	if err != nil {
		SendPayload(map[string]any{
			"msg": "ERROR CODE: 2 [See ERRORS.md]",
		}, logLink)
		fmt.Println(err)
		return 
	} // if 

	us = userSettings

	thisRunTime := int(390 - us.BurnTime) // market open time - burn time
	go sendEnvironmentData() // send env variables as JSON

	userIndicators, err := engine.ParseLogicJSON(userConfigFile)
	if err != nil {
		SendPayload(map[string]any{
			"msg": "ERROR CODE: 1 [See ERRORS.md]",
		}, logLink)
		return
	} // if

	// fmt.Println(userIndicators)
	// fmt.Println(userSettings)

	// Burn In Process
	var burnQuote [5]float64
	var burn []float64
	if userSettings.OverrideBurnIn {
		burnQuote = [5]float64{100, 95, 105, 120, 80}
		burn = overrideBurnIn(userSettings.BurnTime)
	} else {
		fmt.Printf("burn time -> %d, cycletime -> %d\n", userSettings.BurnTime, userSettings.CycleTime)
		burn, burnQuote = BurnIn(userSettings.BurnTime, thisTicker, userSettings.CycleTime)
	} // if-else

	// Initialize Technical Values
	LoadBurnData(&userIndicators, burn)
	UpdateOHLCVDeltas(&userIndicators, burnQuote)

	apiBuf := newAPIBuffer() // store logging info in here
	runtime.GC() // force GC before starting main loop

	go SendPayload(map[string]any{
		"msg": "STAGE: LIVE",
	}, logLink)

	// Main Runtime Loop
	i := 0
	for i < thisRunTime {

		// Pull new Quote
		newQuote, err := api.GetQuote(thisTicker)
		if err != nil {
			go SendPayload(map[string]any {
				"msg" : "ERROR: Could not get market data",
			}, logLink)
			break // Stop if we cannot get a quote
		} // if

		UpdateOHLCVDeltas(&userIndicators, newQuote)

		go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("QUOTE: $%.2f", newQuote[0]),
			}, logLink)
		
		UpdateTechnicals(&userIndicators, newQuote[0])  // Close values
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Updated Technicals",
			}, logLink)
		
		// DEBUG for seeing live updates of technicals
		// for j := range userIndicators.Techs {
		// 	log.Println(userIndicators.Ind[j])
		// }

		// Send JSON of features to ML API
		api.SendData(&userIndicators, thisTicker)
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Sent New Features to ML API",
			}, logLink)
		
		// Get prediction back as JSON
		pred := api.GetPrediction()
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Prediction received from ML API",
			}, logLink)

		// Decide Buy/Sell/Hold 
		handlePrediction(apiBuf, pred, thisTicker)

		go apiBuf.enqueue(map[string]any{ 
			"msg": fmt.Sprintf("STAGE: WAIT (%d seconds)", userSettings.BurnTime),
		}, logLink)

		// Flush all messages to logLink
		go apiBuf.flush(6, time.Duration(userSettings.LogAPIFlushTime)) // items, milliseconds buffer
		
		sendTechnicalData(userIndicators) // send new technical data
		time.Sleep(time.Duration(userSettings.CycleTime) * time.Second)
		i++
		
	} // for

	SendPayload(map[string]any{
		"msg": "STAGE: STOP",
	}, logLink)
} // eventLoop