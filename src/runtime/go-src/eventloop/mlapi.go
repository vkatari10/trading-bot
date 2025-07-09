package eventloop

// This file interact with the ML model in Python to send data and get 
// predictions back to inform the broker API

import (
    "fmt"
    technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
    alpaca "github.com/vkatari10/trading-bot/src/runtime/go-src/alpaca"
    "log"
    "github.com/gorilla/websocket"
)

// Use local server right now
var (
    FixedCols = 5 // The OHLCV bars are fixed (from yfinance DF)
)

// PutPrices loads the intial close, high, low, and open prices that 
// the ML models was trained on (Yfinance includes these by deafult)
func PutPrices(data *technicals.UserData, json map[string]any, ticker string) (map[string]any) {
    bars, err := alpaca.GetQuote(ticker)
    if err != nil {
        log.Println("ERROR: Failed to get market data")
        return nil
    } // if

    json["0"] = bars[0]
    json["1"] = bars[1]
    json["2"] = bars[2]
    json["3"] = bars[3]
    json["4"] = bars[4]

    for i := range data.OHLCVDelta { // Putting OHCLV deltas 
        name := fmt.Sprintf("%d", i + FixedCols)
        json[name] = data.OHLCVDelta[i]
    } // for

    return json
} // PutPrices

// PutNewTechnicals inserts the new Technical Values after updating 
// Values of the Indicators that are not Diff, or Delta
func PutNewTechnicals(data *technicals.UserData, json map[string]any) (map[string]any) {

    for i, ind := range data.Objects {

         name := fmt.Sprintf("%d", i + FixedCols * 2) // *2 since deltas were put on method above

        switch v := ind.(type) {

        case *technicals.SMA:
            json[name] = v.Data[len(v.Data) - 1]
        case *technicals.EMA:
            json[name] = v.Data[len(v.Data) - 1]
        case *technicals.Delta:
            json[name] = v.Value
        case *technicals.Diff:
            json[name] = v.Value
        } // switch

    } // for

    return json
   
} // PutNewTechnicals

// GetLatestData returns back a JSON representation of the lastest values in 
// order as the defined JSON in src/logic/features.json
func MakeMLPayload(obj *technicals.UserData, ticker string) (res map[string]any, err error) {
    var json map[string]any = make(map[string]any)
    json = PutPrices(obj, json, ticker)  // Put OHCLV Values
    json = PutNewTechnicals(obj, json) 
    return json, nil
} // Construct

// websocketWriter will write the features to the ML API Server as a JSON
func websocketWriter(conn *websocket.Conn, payload <-chan map[string]any) {
    for {
        err := conn.WriteJSON(<-payload)
        if err != nil {
            log.Println("ERROR: Could not write JSON to ML API")
        } // if
    } // for
} // websocketWriter

// websocketReader returns the inference based on the written features
func websocketReader(conn *websocket.Conn, result chan<- float64) {
    for {
        var response map[string]any
        err := conn.ReadJSON(&response)
        if err != nil {
            log.Println("ERROR: Could not get result from ML API", err)
        } // if

        serverResult, ok := response["result"].(float64)
        if !ok {
            log.Println("ERROR: Could not read response")
        } // if

        result <- serverResult
    } // for
} // websocketReader