package eventloop

// this file contains methods to send specific non logging
// information to the front end

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
	api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

// SendPayload should send the JSON as an Object to the frontend
func SendPayload(data map[string]any, postLink string) {
	if postLink == logLink && us.LogToStdout { // dev debug mode
		log.Println(data["msg"])
	} // if 

	payload, err := json.Marshal(data)
    if err != nil {
		return
    } // if
		// handle these errors in the future somehow
    resp, err := http.Post(postLink, "application/json", bytes.NewBuffer(payload))
    if err != nil {
		return
    } // if

    defer resp.Body.Close()
} // SendPayload

// sendEnvironmentData will send the environment variable
// to the front end once before any other calls happen
func sendEnvironmentData() {
	go func() {
	SendPayload(map[string]any {
		"refresh_rate": refreshRate,
		"ticker": thisTicker,
		"burn_time": burnWindow,
		"bufer_flush_time": thisBufferFlushTime,
		"always_run": thisAlwaysRun,
		"override_burn_in": thisBurnInOverride,
	},envLink)
	}()
} // sendEnvironmentData

// sendBrokerData will send broker account data everytime
// the eventloop gets back a prediction
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

// sendTechnicalData sendsTechnical data as a JSON to the frontend
func sendTechnicalData(data technicals.UserData) { // copy by value to prevent races
	go func(technicals.UserData) {
		colNameArray := make([]string, len(data.ColNames))
		for key, val := range data.ColNames {
			colNameArray[val] = key
		}
		
		prices := data.OHLCVRaw
		priceDeltas := data.OHLCVDelta

		technicals :=  make([]float64, 0, 10)	

		// for _, ind := range data.Objects {
		// 	switch v := ind.(type) {
		// 	case *technicals.SMA:
		// 		technicals = append(technicals, v.Data[len(v.Data) - 1])
		// 	case *technicals.EMA:
		// 		technicals = append(technicals, v.Data[len(v.Data) - 1])
		// 	case *technicals.Delta:
		// 		technicals = append(technicals, v.Value)
		// 	case *technicals.Diff:
		// 		technicals = append(technicals, v.Value)
		// 	} // swtich
		// } // for

		SendPayload(map[string]any{
			"technicals": technicals,
			"col_names": colNameArray,
			"quotes": prices,
			"quotes_delta": priceDeltas,
		}, dataLink)
	}(data)
} // sendTechnicalData


