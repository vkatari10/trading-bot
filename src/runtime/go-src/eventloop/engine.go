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
	
	userConfigFile := "../../../config/" + featuresFile
	userIndicators, err := engine.InitUserLogic(userConfigFile) // Load user defined technicals
	if err != nil {
		go SendPayload(map[string]any{
			"msg": fmt.Sprintf("Could not properly intiaialize the JSON from %s", featuresFile),
		}, logLink)
	} // if

	burn, burnQuote := BurnIn(thisBurnTime, thisTicker, thisRefreshRate) // intialize burn in data

	engine.LoadBurnData(&userIndicators, burn) // Intialize values for technical indicators
	engine.UpdateOHLCVDeltas(&userIndicators, burnQuote)

	apiBuf := newAPIBuffer() // store logging info in here

	runtime.GC() // force GC before starting main loop

	go SendPayload(map[string]any{
		"msg": "STAGE: LIVE",
	}, logLink)

	// Main Runtime Loop
	i := 0
	for i < thisRunTime {

		newQuote, err := api.GetQuote(thisTicker)
		if err != nil {
			go SendPayload(map[string]any {
				"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // 

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

		handlePrediction(apiBuf, pred, thisTicker) // decide if we need to buy or sell

		//log.Printf("STAGE: WAIT (%d seconds)\n", thisRefreshRate)
		go apiBuf.enqueue(map[string]any{ 
			"msg": fmt.Sprintf("STAGE: WAIT (%d seconds)", thisRefreshRate),
		}, logLink)

		// dump whatever we enqueued to the frontend 
		// issues may arise from using go here bc something could be enqueuing (use wait groups maybe?)
		go apiBuf.flush(6, time.Duration(thisBufferFlushTime)) // items, milliseconds buffer 
		go sendTechnicalData(userIndicators) // send new technical data

		
		time.Sleep(time.Duration(thisRefreshRate) * time.Second)
		i++
	} // for

	go SendPayload(map[string]any{
		"msg": "STAGE: STOP",
	}, logLink)

} // eventLoop

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