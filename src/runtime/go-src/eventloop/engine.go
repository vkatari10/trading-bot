package eventloop

import (
	"log"
	"time"
	"runtime"
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

	log.Println("STAGE: BURN IN")
	
	//burn := BurnIn()

	// dummy data 
	burn := []float64{163.42, 199.08, 184.21, 216.77, 152.93,
189.35, 173.88, 201.67, 218.19, 167.04,
153.21, 174.66, 211.05, 197.48, 158.89,
205.76, 161.57, 182.10, 194.33, 159.62,
212.98, 188.71, 168.25, 200.83, 178.55,
215.60, 166.09, 209.40, 170.46, 185.79, 145.5}

	userIndicators, err := engine.InitUserLogic("features.json") // Load user defined technicals
	if err != nil {
		log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
	} // if

	burnQuote, err := api.GetQuote(Ticker)
	if err != nil {
		log.Printf("ERROR: market data could not be pulled")
	} // if
	
	engine.LoadBurnData(&userIndicators, burn) // Intialize values for technical indicators
	engine.UpdateOHLCVDeltas(&userIndicators, burnQuote)

	log.Println(userIndicators)

	runtime.GC() // force GC before starting main loop

	log.Println("STAGE: LIVE")

	// Main Runtime Loop
	i := 0
	for i < TotalUpTime {

		// use close price to update technicals
		
		newQuote, err := api.GetQuote(Ticker)
		if err != nil {
			log.Print("ERROR: market data could not be pulled")
		} // 

		engine.UpdateOHLCVDeltas(&userIndicators, newQuote)
		log.Printf("QUOTE: %f\n", newQuote)
		
		// call GetNew methods on each indicator
		log.Println("UPDATE: Updated Technicals")
		engine.UpdateTechnicals(&userIndicators, newQuote[0])  // Close values
		
		// DEBUG for seeing live updates of technicals
		// for j := range userIndicators.Techs {
		// 	log.Println(userIndicators.Ind[j])
		// }

		// Send JSON of features to ML API
		log.Println("UPDATE: Sent Features to ML model")
		api.SendData(&userIndicators, Ticker)

		// Get prediction back as JSON
		log.Println("UPDATE: Got prediction from ML model")
		pred := api.GetPrediction()

		if pred > 0 {
			log.Printf("DECIDE: Buy 1 share of %s\n", Ticker)
			api.PlaceMarketOrder(Ticker, 1, "buy")
		} else if pred < 0 {
			log.Printf("DECIDE: Sell 1 share of %s\n", Ticker)
			api.PlaceMarketOrder(Ticker, 1, "sell")
		} else {
			log.Printf("DECIDE: Do nothing\n")
		} // if-else
			
		log.Printf("STAGE: WAIT (%d seconds)\n", TickTime)
		time.Sleep(TickTime * time.Second)
		i++
	} // for

	log.Println("STAGE: STOP")

} // eventLoop

//BurnIn Loads the Burn in Data to intialize technical indicators
func BurnIn() ([]float64) {

	var burn []float64 = make([]float64, BurnInWindow);

	for i := range burn {
		newQuote, err := api.GetQuote(Ticker)
		if err != nil {
			log.Printf("ERROR: market data could not be pulled")
		} // if
		burn[i] = newQuote[0]
		log.Printf("QUOTE: %f", newQuote)
		time.Sleep(TickTime * time.Second) // wait 60 till next tick 
	} // for

	return burn
} // BurnIn