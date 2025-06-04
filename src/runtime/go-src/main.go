package main

// Execution file for the runtime environemnt

import (
	"runtime"
	"time"

	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	"github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

	// "sync"
	//import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
	"log"
)

var ticker string = "AAPL"
var burnInWindow int = 1

// Main Runtime Engine should be placed here
func main() {

	log.Println("STAGE: BURN IN")
	
	var burn []float64 = make([]float64, burnInWindow) // Burn in for 30 minutes

	for i := range burn {
		new_quote, err := api.GetQuote(ticker)
		if err != nil {
			log.Printf("ERROR: market data could not be pulled")
		}
		burn[i] = new_quote
		log.Printf("QUOTE: %f", new_quote)
		time.Sleep(1 * time.Second) // wait 60 till next tick 
	} // for

	userIndicators, err := engine.InitUserLogic("features.json") // Load user defined technicals
	if err != nil {
		log.Fatal("ERROR: could not parse user defined JSON in src/logic properly")
	} // if

	engine.LoadBurnData(&userIndicators, burn) // Intialize values for technical indicators

	runtime.GC() // force GC before starting main loop

	log.Println("STAGE: LIVE")
	

} // main
