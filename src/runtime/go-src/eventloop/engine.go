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
		log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
	} // if

	//burn := BurnIn(thisBurnTime, thisTicker, thisRefreshRate) // intialize burn in data

	// dummy data to avoid waiting for input data
	burn := []float64{163.42, 199.08, 184.21, 216.77, 152.93,
					189.35, 173.88, 201.67, 218.19, 167.04,
					153.21, 174.66, 211.05, 197.48, 158.89,
					205.76, 161.57, 182.10, 194.33, 159.62,
					212.98, 188.71, 168.25, 200.83, 178.55,
					215.60, 166.09, 209.40, 170.46, 185.79, 145.5}

	burnQuote, err := api.GetQuote(thisTicker)	 // can remove later (use when not calling BurnIn())
	if err != nil {
		log.Printf("ERROR: market data could not be pulled")
	} // if		

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

		if pred > 0 { // buy
			log.Printf("DECIDE: Buy 1 share of %s\n", thisTicker)
			go api.PlaceMarketOrder(thisTicker, 1, "buy")
			go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: BUY %s", thisTicker),
			}, logLink)
		} else if pred < 0 { // sell
			log.Printf("DECIDE: Sell 1 share of %s\n", thisTicker)
			go api.PlaceMarketOrder(thisTicker, 1, "sell")
			go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: SELL %s", thisTicker),
			}, logLink)
		} else {
			log.Printf("DECIDE: Do nothing\n")
			go apiBuf.enqueue(
			map[string]any{
				"msg": fmt.Sprintf("DECIDE: HOLD %s", thisTicker),
			}, logLink)
		} // if-else
			
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
			log.Printf("ERROR: market data could not be pulled")
			go SendPayload(map[string]any {
			"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // if
		burn[i] = newQuote[0]
		log.Printf("QUOTE: %f", newQuote)
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
}