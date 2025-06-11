package eventloop

import (
	"log"
	"time"
	"runtime"
	"fmt"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

var ( 
	Ticker string = "AAPL"
	BurnInWindow int = 31 // If using delta/diff add one to account for its equation of the highest window value (does not check)
	TotalUpTime = 450 - BurnInWindow
	TickTime = time.Duration(10) // seconds (use 1 for testing)
) // Environment Variables 


func Run() {
	//burn := BurnIn() // intialize burn in data

	// dummy data to avoid waiting for input data
	burn := []float64{163.42, 199.08, 184.21, 216.77, 152.93,
					189.35, 173.88, 201.67, 218.19, 167.04,
					153.21, 174.66, 211.05, 197.48, 158.89,
					205.76, 161.57, 182.10, 194.33, 159.62,
					212.98, 188.71, 168.25, 200.83, 178.55,
					215.60, 166.09, 209.40, 170.46, 185.79, 145.5}

	burnQuote, err := api.GetQuote(Ticker)	 // can remove later (use when not calling BurnIn())
	if err != nil {
		log.Printf("ERROR: market data could not be pulled")
	} // if		

	// COULD move this up before the burn in data to intialize the OHCLV Deltas better 
	userIndicators, err := engine.InitUserLogic("features.json") // Load user defined technicals
	if err != nil {
		log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
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
	for i < TotalUpTime {
		
		newQuote, err := api.GetQuote(Ticker)
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
		api.SendData(&userIndicators, Ticker)
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
			log.Printf("DECIDE: Buy 1 share of %s\n", Ticker)
			api.PlaceMarketOrder(Ticker, 1, "buy")
		} else if pred < 0 { // sell
			log.Printf("DECIDE: Sell 1 share of %s\n", Ticker)
			api.PlaceMarketOrder(Ticker, 1, "sell")
		} else {
			log.Printf("DECIDE: Do nothing\n")
			go SendPayload(map[string]any{
				"msg": "DECISION: HOLD",
			}, logLink)
		} // if-else

		go apiBuf.offload(3, 100) // milliseconds
			
		log.Printf("STAGE: WAIT (%d seconds)\n", TickTime)
		time.Sleep(TickTime * time.Second)
		i++
	} // for

	log.Println("STAGE: STOP")
	go apiBuf.enqueue(
		map[string]any{
			"msg": "STAGE: STOP",
		}, logLink)

} // eventLoop

//BurnIn Loads the Burn in Data to intialize technical indicators
func BurnIn() (arr []float64, finalQuote [5]float64) {
	go SendPayload(map[string]any{
		"msg": "STAGE: BURN IN",
		}, logLink)
	log.Println("STAGE: BURN IN")

	// stores burn data
	var burn []float64 = make([]float64, BurnInWindow);

	// stores latest quotes
	var newQuote [5]float64

	for i := range burn {
		newQuote, err := api.GetQuote(Ticker)
		if err != nil {
			log.Printf("ERROR: market data could not be pulled")
			go SendPayload(map[string]any {
			"msg" : "ERROR: Could not get market data",
			}, logLink)
		} // if
		burn[i] = newQuote[0]
		log.Printf("QUOTE: %f", newQuote)
		time.Sleep(TickTime * time.Second) // burn in rate at same tick time for main loop
	} // for

	return burn, newQuote
} // BurnIn