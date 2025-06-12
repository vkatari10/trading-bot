package eventloop

// this file contains methods to send specific non logging 
// information to the front end

import (
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

// sendEnvironmentData will send the environment variable
// to the front end once before any other calls happen
func sendEnvironmentData() {
	go func() {
	SendPayload(map[string]any {
		"REFRESH_RATE": refreshRate,
		"TICKER": ticker,
		"BURN_TIME": burnWindow,
	},envLink)
	}()
} // sendEnvironmentData

func sendBrokerData() {
	api.Acct("hello")
	// get cash and other stuff from here every time a 
	// a trade is place to get new position data, account 
	// and cash values
} // sendBrokerData

func sendTechnicalData(data engine.UserData) {
	go func(engine.UserData) {
		
		columnNames := data.ColNames
		prices := data.OHLCVRaw
		priceDeltas := data.OHLCVDelta

		technicals :=  make([]float64, 0, 10)	

		for _, ind := range data.Objects {
			switch v := ind.(type) {
			case *engine.SMA:
				technicals = append(technicals, v.Data[len(v.Data) - 1])
			case *engine.EMA:
				technicals = append(technicals, v.Data[len(v.Data) - 1])
			case *engine.Delta:
				technicals = append(technicals, v.Value)
			case *engine.Diff:
				technicals = append(technicals, v.Value)
			} // swtich
		} // for

		SendPayload(map[string]any{
			"TECHNICALS": technicals,
			"COL_NAMES": columnNames,
			"PRICE": prices,
			"PRICE_DELTAS": priceDeltas,
		}, dataLink)

	}(engine.UserData{})
}