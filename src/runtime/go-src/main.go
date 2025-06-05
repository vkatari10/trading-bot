package main

// Execution file for the runtime environemnt

import (
	"runtime"
	"time"
	"github.com/vkatari10/trading-bot/src/runtime/go-src/api" // api
	"github.com/vkatari10/trading-bot/src/runtime/go-src/engine" // engine
	//"sync"
	"log"
)	

var ( 
	Ticker string = "AAPL"
	BurnInWindow int = 30
	TotalUpTime = 450 - BurnInWindow
	TickTime = time.Duration(1) // seconds
) // Environment Variables 

//Main Runtime Engine should be placed here
func main() {

	log.Println("STAGE: BURN IN")
	
	var burn []float64 = make([]float64, BurnInWindow) // Burn in for 30 minutes

	for i := range burn {
		newQuote, err := api.GetQuote(Ticker, "c")
		if err != nil {
			log.Printf("ERROR: market data could not be pulled")
		} // if
		burn[i] = newQuote
		log.Printf("QUOTE: %f", newQuote)
		time.Sleep(TickTime * time.Second) // wait 60 till next tick 
	} // for

	userIndicators, err := engine.InitUserLogic("features.json") // Load user defined technicals
	if err != nil {
		log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
	} // if

	engine.LoadBurnData(&userIndicators, burn) // Intialize values for technical indicators

	runtime.GC() // force GC before starting main loop

	log.Println("STAGE: LIVE")

	// Main Runtime Loop
	i := 0
	for i < TotalUpTime {

		// use close price to update technicals
		
		newQuote, err := api.GetQuote(Ticker, "c")
		if err != nil {
			log.Print("ERROR: market data could not be pulled")
		} // if

		log.Printf("QUOTE: %f\n", newQuote)
		
		// call GetNew methods on each indicator
		log.Println("UPDATE: Updated Technicals")
		engine.UpdateTechnicals(&userIndicators, newQuote)
		
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
			log.Printf("DECIDE: Buy 1 share of %s\n", Ticker)
			api.PlaceMarketOrder(Ticker, 1, "sell")
		} else {
			log.Printf("DECIDE: Do nothing\n")
		} // if-else
			
		log.Printf("STAGE: WAIT (%d seconds)\n", TickTime)
		time.Sleep(TickTime * time.Second)
		i++
	} // for

	log.Println("STAGE: STOP")

} // main


