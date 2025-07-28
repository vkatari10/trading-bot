package eventloop

// This file contains methods to handle object creation
// and constructing the initial OHLCV arrays
//
// Author: Vikas Katari
// Date: 07/18/2025

import "C"
import (
	"fmt"
	"log"
	"math/rand"
	"time"
	alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

// All methods should receive a technicals.RuntimeData object and
//
// 1. Initialize its OHLCV field arrays (technicals.TALIBWrapper)
// 2. Initialize C pointers to all 5 arrays
// 3. Return the delta of each OHLCV as well from the last to second to last bars

// BurnIn burns in the OHLCV arrays with real time data via the market data API,
// modifies the RuntimeData object in place 
func burnIn(rd *technicals.RuntimeData, ticker string) (error) {
	log.Printf("STAGE: BURN IN (%d minutes)", rd.RuntimeSettings.BurnTime)

	len := rd.RuntimeSettings.BurnTime

	open := make([]float64, 0, len * technicals.CapLimitMultiplier)
	high := make([]float64, 0, len * technicals.CapLimitMultiplier)
	low := make([]float64, 0, len * technicals.CapLimitMultiplier)
	close := make([]float64, 0, len * technicals.CapLimitMultiplier)
	vol := make([]float64, 0, len * technicals.CapLimitMultiplier)

	for i := range len {
		newQuote, err := alpaca.GetAlpacaBars(ticker)
		if err != nil {
			return fmt.Errorf("burnIn failed with error %v", err)
		} 

		open = append(open, newQuote[0])
		high = append(high, newQuote[1])
		low = append(low, newQuote[2])
		close = append(close, newQuote[3])
		vol = append(vol, newQuote[4])

		log.Printf("%s: $%.2f (%d / %d) burned in", ticker, newQuote[3], i + 1, len)

		time.Sleep(time.Duration(rd.RuntimeSettings.CycleTime) * time.Second)
	}

	rd.OHLCV.Open = open
	rd.OHLCV.High = high
	rd.OHLCV.Low = low
	rd.OHLCV.Close = close
	rd.OHLCV.Volume = vol

	// initialize pointers
	rd.SetAllPointersToIndex(0)

	// find deltas
	rd.UpdateOHLCVDeltas()
	
	return nil
} // BurnIn

// overrideBurnIn creates fake data to skip the burn in process
// immediately, modifies the RuntimeData object in place
func OverrideBurnIn(rd *technicals.RuntimeData, ticker string, postLink string) { // RENAME METHOD TO BE UNEXPORTED AFTER REMOVING DEPRECATED VERSION
	go SendPayload(
		map[string]any{
			"msg": fmt.Sprintf("%s Overriding burn in", ticker),
		}, postLink,
	)

	len := rd.RuntimeSettings.BurnTime

	open := make([]float64, 0,len * technicals.CapLimitMultiplier)
	high := make([]float64, 0, len * technicals.CapLimitMultiplier)
	low := make([]float64, 0, len * technicals.CapLimitMultiplier)
	close := make([]float64, 0, len * technicals.CapLimitMultiplier)
	vol := make([]float64, 0, len * technicals.CapLimitMultiplier)

	for range len {
		open = append(open, rand.Float64() * 10)
		high = append(high, rand.Float64() * 10)
		low = append(low, rand.Float64() * 10)
		close = append(close, rand.Float64() * 10)
		vol = append(vol, rand.Float64() * 10)
	}

	rd.OHLCV.Open = open
	rd.OHLCV.High = high
	rd.OHLCV.Low = low
	rd.OHLCV.Close = close
	rd.OHLCV.Volume = vol

	// initialize pointers
	rd.SetAllPointersToIndex(0)

	// find deltas
	rd.UpdateOHLCVDeltas()
} // OverrideBurnIn