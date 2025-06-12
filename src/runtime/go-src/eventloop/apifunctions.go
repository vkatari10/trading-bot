package eventloop

// this file contains methods to send specific non logging
// information to the front end

import (
	//api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	// "fmt"

	"github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

// sendEnvironmentData will send the environment variable
// to the front end once before any other calls happen
func sendEnvironmentData() {
	go func() {
	SendPayload(map[string]any {
		"refresh_rate": refreshRate,
		"ticker": ticker,
		"burn_time": burnWindow,
	},envLink)
	}()
} // sendEnvironmentData

func sendBrokerData() {
	go func() {

	qty, avgCost, marketVal, err := api.GetPositions()
	if err != nil {
		qty = 0
		avgCost = 0
		marketVal = 0
	} // if

	cash, accountValue, err := api.GetCashValue()
	if err != nil {
		cash = 0
		accountValue = 0
	} // if

	SendPayload(map[string]any{
		"cash": cash,
		"account_value": accountValue,
		"stock_qty": qty,
		"stock_cost": avgCost,
		"market_value": marketVal,
	}, brokerLink)
	}()
} // sendBrokerData

func sendTechnicalData(data engine.UserData) {
	go func(engine.UserData) {
		
		columnNames := data.ColNames // this is constant 
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
			"technicals": technicals,
			"col_names": columnNames,
			"quotes": prices,
			"quotes_delta": priceDeltas,
		}, dataLink)


	}(data)
} // sendTechnicalData


