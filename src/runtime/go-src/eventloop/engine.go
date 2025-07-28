package eventloop

// UNCOMMENT SECTIONS THAT USE OLD RUNTIME SETTINGS
// CALLS IN THE MAIN EVENTLOOP

import (
	// "time"
	"fmt"
	alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
	"sync"
	"github.com/gorilla/websocket"
	// apibuffer "github.com/vkatari10/trading-bot/src/runtime/go-src/apibuffer"
)

var (
	us technicals.RuntimeSettings // to use its values in the rest of the package
)

// Start Entry point of the program to read all tickers and 
// feed them to the Spawner which creates independent event loops for
// each ticker
func Start() {
	tickers, err := json.GetTradeTickers("../../../" + getFileName())
	if err != nil {
		SendPayload(map[string]any{
			"msg": "ERROR CODE: 3 [See ERRORS.md]",
		}, logLink)
		fmt.Println(err)
		return
	} // if

	Spawner(tickers)
} // Start

// Spawner starts runtime loops for every ticker giving each 
// one its own independent event loop
func Spawner(tickers []string) {

	var wg sync.WaitGroup

	for i := range tickers {
		wg.Add(1)

		go func(t string, id int) {
			defer wg.Done()
			EventLoop(t, id)
		}(tickers[i], i)
	} // for

	wg.Wait()

} // Spawner

// tempTicker is how we will integrate specific tickers onto the eventloop
func EventLoop(tempTicker string, tickerID int) {
	url := fmt.Sprintf("ws://localhost:8000/results-ws/%d", tickerID)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	defer c.Close()

	// apiInputChan := make(chan map[string]any, 1)
	apiOutputChan := make(chan float64, 1)
	go websocketReader(c, apiOutputChan)
	// go websocketWriter(c, apiInputChan)

	
	// Load User JSON -> Convert to Go Struct

	// TODO create a method to just return the map 
	// and then send the map ONLY to the parsers
	userConfigFile := "../../../" + getFileName()

	// userSettings, err := json.GetRuntimeSettings(userConfigFile)
	// if err != nil {
	// 	SendPayload(map[string]any{
	// 		"msg": "ERROR CODE: 2 [See ERRORS.md]",
	// 	}, logLink)
	// 	fmt.Println(err)
	// 	return 
	// } // if 

	// us = userSettings

	thisRunTime := int(390 - us.BurnTime) // market open time - burn time
	go sendEnvironmentData() // send env variables as JSON

	userIndicators, err := json.ParseLogicJSON(userConfigFile)
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
	// if userSettings.OverrideBurnIn {
	// 	burnQuote = [5]float64{100, 95, 105, 120, 80}
	// 	burn = overrideBurnIn(userSettings.BurnTime)
	// } else {
	// 	//fmt.Printf("burn time -> %d MINUTES, cycle time -> %d SECONDS\n", userSettings.BurnTime, userSettings.CycleTime)
	// 	burn, burnQuote = BurnIn(userSettings.BurnTime, tempTicker, userSettings.CycleTime)
	// } // if-else

	// Initialize Technical Values
	LoadBurnData(&userIndicators, burn)
	UpdateOHLCVDeltas(&userIndicators, burnQuote)

	// apiBuf := apibuffer.NewAPIBuffer() // store logging info in here

	go SendPayload(map[string]any{
		"msg": fmt.Sprintf("(%s) STAGE: LIVE", tempTicker),
	}, logLink)

	// Main Runtime Loop
	i := 0
	for i < thisRunTime {

		// Pull new Quote
		newQuote, err := alpaca.GetQuote(tempTicker)
		if err != nil {
			go SendPayload(map[string]any {
				"msg" : fmt.Sprintf("(%s) ERROR: Could not get market data", tempTicker),
			}, logLink)
			/*

			Allow user to define how many retries to go for before quitting
			rather than just stopping

			*/
			break // Stop if we cannot get a quote 
		} // if

		UpdateOHLCVDeltas(&userIndicators, newQuote)

		// go apiBuf.Enqueue(
		// 	map[string]any{
		// 		"msg": fmt.Sprintf("%s: $%.2f", tempTicker, newQuote[0]),
		// 	}, logLink)
		
		UpdateTechnicals(&userIndicators, newQuote[0])  // Close values
		// go apiBuf.Enqueue(
		// 	map[string]any{
		// 		"msg": "UPDATE: Updated Technicals",
		// 	}, logLink)
		
		// DEBUG for seeing live updates of technicals
		// for j := range userIndicators.Techs {
		// 	log.Println(userIndicators.Ind[j])
		// }

		// payload, err := MakeMLPayload(&userIndicators, tempTicker)
		// if err != nil {
		// 	go SendPayload(map[string]any{
		// 		"msg": "ERROR: could not produce ML Payload",
		// 	}, logLink)
		// } // if
		// apiInputChan <- payload

		// pred := <-apiOutputChan

		// Decide Buy/Sell/Hold 
		// handlePrediction(apiBuf, pred, tempTicker)

		// go apiBuf.enqueue(map[string]any{ 
		// 	"msg": fmt.Sprintf("STAGE: WAIT (%d seconds)", userSettings.CycleTime),
		// }, logLink)

		// // Flush all messages to logLink
		// go apiBuf.flush(6, time.Duration(userSettings.LogAPIFlushTime)) // items, milliseconds buffer
		
		// sendTechnicalData(userIndicators) // send new technical data
		// time.Sleep(time.Duration(userSettings.CycleTime) * time.Second)
		// i++
		
	} // for

	SendPayload(map[string]any{
		"msg": "STAGE: STOP",
	}, logLink)
} // eventLoop