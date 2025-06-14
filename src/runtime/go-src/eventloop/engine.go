package eventloop

import (
	"log"
	"time"
	"runtime"
	"fmt"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

func Run() {
	go sendEnvironmentData() // send env variables to front end

	thisBurnTime, thisRefreshRate, err := intializeVariables()
	if err != nil {
		log.Fatal(err)
	}
	thisRunTime := int(450 - thisBurnTime)
	thisTicker := getTicker()

	// COULD move this up before the burn in data to intialize the OHCLV Deltas better 
	userIndicators, err := engine.InitUserLogic("features.json") // Load user defined technicals
	if err != nil {
		// log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
		go SendPayload(map[string]any{
			"msg": "ERROR: could not parse user defined JSON in src/logic properly",
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
		log.Printf("QUOTE: $%.2f\n", newQuote[0])

		go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("QUOTE: $%.2f", newQuote[0]),
			}, logLink)
		
		engine.UpdateTechnicals(&userIndicators, newQuote[0])  // Close values
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Updated Technicals",
			}, logLink)
		log.Println("UPDATE: Updated Technicals")
		
		// DEBUG for seeing live updates of technicals
		// for j := range userIndicators.Techs {
		// 	log.Println(userIndicators.Ind[j])
		// }

		// Send JSON of features to ML API
		api.SendData(&userIndicators, thisTicker)
		log.Println("UPDATE: Sent Features to ML model")
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Sent New Features to ML API",
			}, logLink)
		

		// Get prediction back as JSON
	
		pred := api.GetPrediction()
		log.Println("UPDATE: Got prediction from ML model")
		go apiBuf.enqueue(
			map[string]any{
				"msg": "UPDATE: Prediction recieved from ML API",
			}, logLink)

		handlePrediction(apiBuf, pred, thisTicker) // decide if we need to buy or sell

		log.Printf("STAGE: WAIT (%d seconds)\n", thisRefreshRate)
		go apiBuf.enqueue(map[string]any{ 
			"msg": fmt.Sprintf("STAGE: WAIT (%d seconds)", thisRefreshRate),
		}, logLink)

		// dump whatever we enqueued to the frontend 
		// issues may arise from using go here bc something could be enqueuing (use wait groups maybe?)
		go apiBuf.offload(6, 2000) // milliseconds
		go sendTechnicalData(userIndicators) // send new technical data

		time.Sleep(time.Duration(thisRefreshRate) * time.Second)
		i++
	} // for

	log.Println("STAGE: STOP")
	go SendPayload(map[string]any{
		"msg": "STAGE: STOP",
	}, logLink)

} // eventLoop

//BurnIn Loads the Burn in Data to intialize technical indicators
func BurnIn(burnTime int, ticker string, refresh time.Duration) (arr []float64, finalQuote [5]float64) {
	go SendPayload(map[string]any{
		"msg": "STAGE: BURN IN",
		}, logLink)
	log.Println("STAGE: BURN IN")

	// stores burn data
	var burn []float64 = make([]float64, burnTime);

	// stores latest quotes
	var newQuote [5]float64

	for i := range burn {
		newQuote, err := api.GetQuote(ticker)
		if err != nil {
			// log.Printf("ERROR: market data could not be pulled")
			go SendPayload(map[string]any {
			"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // if
		burn[i] = newQuote[0]
		// log.Printf("QUOTE: %f", newQuote)
		go SendPayload(map[string]any {
			"msg": fmt.Sprintf("QUOTE: $%.2f", newQuote[0]),
		}, logLink)
		time.Sleep(refresh * time.Second) // burn in rate at same tick time for main loop
	} // for

	return burn, newQuote
} // BurnIn

// initializeVariables returns the burn window time and refresh rate
// by calling the values from the .env file
func intializeVariables() (int, time.Duration, error) {
	burn, err := getBurnWindow()
	if err != nil {
		return 0, 0, fmt.Errorf("%v", err)
	} // if

	tick, err := getRefreshRate()
	if err != nil {
		return 0, 0, fmt.Errorf("%v", err)
	} // if

	return burn, tick, nil
} // intializeVariables

// handlePrediction handles the prediction made by the ML model by 
// working with the broker API
func handlePrediction(apiBuffer *APIBuffer, prediction float64, ticker string) {

	decisionMsg := "DECIDE: "
		var decision string;

		if prediction > 0 { // buy
			// log.Printf("DECIDE: Buy 1 share of %s\n", thisTicker)
			decision = "buy"
			go api.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "BUY"
		} else if prediction < 0 { // sell
			// log.Printf("DECIDE: Sell 1 share of %s\n", thisTicker)
			decision = "sell"
			go api.PlaceMarketOrder(ticker, 1, decision)
			decisionMsg += "SELL"
		} else {
			// log.Printf("DECIDE: Do nothing\n")
			decisionMsg += "HOLD"
		} // if-else
		
		go apiBuffer.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: %s %s", decisionMsg, ticker),
			}, logLink)
		go sendBrokerData()

} // handlePrediction