package eventloop

// Contains new v0.4.0 main event loop logic
//
// Author: Vikas Katari
// Date: 07/24/2025

import (
	"fmt"
	"sync"
	"time"
	"github.com/gorilla/websocket"
	"encoding/json"
	alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
	apibuffer "github.com/vkatari10/trading-bot/src/runtime/go-src/apibuffer"
	jsonInternal "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

var (
	log_to_stdio bool
)

// NOTE:
// IFF errors are encountered DO NOT put SendPayload() on its own goroutine
// Why? Force the Payload to be sent and blocked before returning to ensure
// the payload is delivered to its intended source


// deepCopyRuntimeData makes a deep copy of a runtimeData structure
func deepCopyRuntimeData(rd technicals.RuntimeData) technicals.RuntimeData {
	var copy technicals.RuntimeData
	b, _ := json.Marshal(rd)
	json.Unmarshal(b, &copy)
	return copy
} // deepCopyRuntimeData

// startEventLoops starts a new event loop for every 
// live trade ticker declared in the JSON file
func StartEventLoops() {
	rd, err := jsonInternal.NewRuntimeData("../../../" + getFileName())
	if err != nil {
		fmt.Printf("Exit Code 1: %v", err)
		return
	}
	log_to_stdio = rd.RuntimeSettings.LogToStdout

	var wg sync.WaitGroup

	for i := range rd.Tickers {
		wg.Add(1)

		go func(rd technicals.RuntimeData, t string, id int) {
			defer wg.Done()
			eventLoop(rd, t, id)
		}(deepCopyRuntimeData(rd), rd.Tickers[i], i)

	}

	wg.Wait()
} // startEventLoops

// Handles the entire eventloop process from initialization, trading, ML prediction
func eventLoop(rd technicals.RuntimeData, ticker string, id int) {
	// websocket connection setup
	urlML := fmt.Sprintf(mlServerLink + "/%d", id)
	
	logUrl := fmt.Sprintf(logLink + "/%d", id)
	dataUrl := fmt.Sprintf(dataLink + "/%d", id)
	_ = fmt.Sprintf(brokerLink + "/%d", id)
	c, _, err := websocket.DefaultDialer.Dial(urlML, nil)
	if err != nil {
		SendPayload(
			map[string]any{"msg": fmt.Sprintf("Exit code 2: %v", err)},
			logUrl,
		)
		return
	}
	defer c.Close()
	apiInputChan := make(chan map[string]float64, 1)
	apiOutputChan := make(chan float64, 1)
	go websocketReader(c, apiOutputChan)
	go websocketWriter(c, apiInputChan)

	// Initializing runtimeData obj
	thisRuntime := int(390 - rd.RuntimeSettings.BurnTime)  // 390 = market open time
	// go sendEnvironmentData() if u want here or something

	if rd.RuntimeSettings.OverrideBurnIn {
		OverrideBurnIn(&rd, ticker, logUrl)
	} else {
		burnIn(&rd, ticker)
	}

	err = rd.InitRelationships()
	if err != nil {
		SendPayload(
			map[string]any{"msg": "Exit code 3: JSON object failed to be initialized"},
			logUrl,
		)
		return
	}	

	apibuf := apibuffer.NewAPIBuffer()
	i := 0
	for i < int(float64(thisRuntime) / rd.RuntimeSettings.CycleTime) { // main event loop
		newBars, err := alpaca.GetAlpacaBars(ticker)
		if err != nil {
			go SendPayload(map[string]any{"msg": fmt.Sprintf("%s: Failed to fet latest quote, waiting...", ticker)}, logUrl)
			time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second) 
			i++
			continue
		}

		go apibuf.Enqueue(
			map[string]any{
				"msg": fmt.Sprintf("%s: %.2f", ticker, newBars[3]),
			}, 
			logUrl,
		)

		rd.PopLeft()
		rd.AppendNewOHLCV(*newBars)
		rd.UpdateOHLCVDeltas()

		err = rd.UpdateTALIBTechnicals()
		if err != nil {
			go SendPayload(map[string]any{"msg": fmt.Sprintf("%s: Feature computations failed, waiting...", ticker)}, logUrl)
			time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second) 
			i++
			continue
		}

		err = rd.UpdateOtherTechnicals()
		if err != nil {
			go SendPayload(map[string]any{"msg": fmt.Sprintf("%s: Failed to Update non-TA-Lib features, waiting...", ticker)}, logUrl)
			time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second) 
			i++
			continue
		}

		err = rd.UpdateRelationships()
		if err != nil {
			go SendPayload(map[string]any{"msg": fmt.Sprintf("%s: Label logic computations failed, waiting...", ticker)}, logUrl)
			time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second) 
			i++
			continue
		}

		apiInputChan <- rd.FeatureJSON // will block if these fails, update internal methods
		prediction := <-apiOutputChan

		go apibuf.Enqueue(
			map[string]any{
				"msg": fmt.Sprintf("%s: Received Prediction of %.2f", ticker, prediction),
			}, logUrl,
		)

		handlePrediction(apibuf, prediction, ticker, logUrl)

		time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second)
		go apibuf.FlushAll(time.Duration(rd.RuntimeSettings.LogAPIFlushTime), SendPayload)
		sendDataPayload(rd.FeatureJSON, dataUrl) // impl uses goroutined anon function 
		i++
	}
} // eventLoop

