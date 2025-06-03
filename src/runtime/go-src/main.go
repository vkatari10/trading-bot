package main

// Execution file for the runtime environemnt

import (
	"fmt"
	"runtime"
	"time"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	// "sync"
	//import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
	"log"
)

var ticker string = "AAPL"

// Main Runtime Engine should be placed here
func main() {
	
	var burn []float32 = make([]float32, 30) // Burn in for 30 minutes

	for i := range burn {
		new_quote, err := api.GetQuote(ticker)
		if err != nil {
			log.Fatal("Market Data Streaming API Failure")
		}
		burn[i] = new_quote
		fmt.Println(new_quote)
		time.Sleep(60 * time.Second)
	} // for

	runtime.GC() // force GC before starting main loop



	// Load user defined technicals as a LiveIndicator struct
	// inds, err := engine.InitUserLogic("features.json")
	// if err != nil {
	// 	fmt.Errorf("%v", err.Error())
	// } // if


	// fmt.Println(inds.Techs)
	// fmt.Println(inds.Ind)

	// sma1, ok := inds.Ind[0].(*engine.SMA)

	// if !ok {
	// 	fmt.Println("Hello some error")
	// }

	// fmt.Println(sma1.Sum)

	deez, err := api.GetQuote("AAPL")
	if err != nil {
		fmt.Printf("%w", err)
	}

	fmt.Println(deez)

	


} // main
