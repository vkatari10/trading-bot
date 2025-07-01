package eventloop

import (
	"time"
	"runtime"
	"fmt"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

func EventLoop() {
	thisRunTime := int(390 - thisBurnTime)
	go sendEnvironmentData() // send env variables to front end
	
	// Load User JSON -> Convert to Go Struct
	userConfigFile := "../../../" + getFileName()
	userIndicators, err := engine.ParseLogicJSON(userConfigFile)
	if err != nil {
		SendPayload(map[string]any{
			"msg": "ERROR CODE: 1 [See ERRORS.md]",
		}, logLink)
		return
	} // if

	// fmt.Println(userIndicators)

	// Burn In Process
	var burnQuote [5]float64
	var burn []float64
	if thisBurnInOverride {
		burnQuote = [5]float64{100, 95, 105, 120, 80}
		burn = overrideBurnIn(thisBurnTime)
	} else {
		burn, burnQuote = BurnIn(thisBurnTime, thisTicker, thisRefreshRate)
	} // if-else

	// Initialize Technical Values
	engine.LoadBurnData(&userIndicators, burn)
	engine.UpdateOHLCVDeltas(&userIndicators, burnQuote)

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

		engine.UpdateOHLCVDeltas(&userIndicators, newQuote)

		go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("QUOTE: $%.2f", newQuote[0]),
			}, logLink)
		
		engine.UpdateTechnicals(&userIndicators, newQuote[0])  // Close values
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
				"msg": "UPDATE: Prediction recieved from ML API",
			}, logLink)

		// Decide Buy/Sell/Hold 
		handlePrediction(apiBuf, pred, thisTicker)

		go apiBuf.enqueue(map[string]any{ 
			"msg": fmt.Sprintf("STAGE: WAIT (%d seconds)", thisRefreshRate),
		}, logLink)

		// Flush all messages to logLink
		go apiBuf.flush(6, time.Duration(thisBufferFlushTime)) // items, milliseconds buffer
		
		sendTechnicalData(userIndicators) // send new technical data
		time.Sleep(time.Duration(thisRefreshRate) * time.Second)
		i++
	} // for

	go SendPayload(map[string]any{
		"msg": "STAGE: STOP",
	}, logLink)
} // eventLoop