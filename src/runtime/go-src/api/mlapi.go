package api

// This file interact with the ML model in Python to send data and get 
// predictions back to inform the broker API

import (
    "fmt"
    engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine" 
    "encoding/json"
    "bytes"
    "net/http" 
    "log"
    "io"
)

// Use local server right now
var (
    ServerLink string = "http://127.0.0.1:5000/api/prediction"
    FixedCols = 5 // The OHLCV bars are fixed (from yfinance DF)
)

// PutPrices loads the intial close, high, low, and open prices that 
// the ML models was trained on (Yfinance includes these by deafult)
func PutPrices(data *engine.UserData, json map[string]any, ticker string) (map[string]any) {
    bars, err := GetQuote(ticker)
    if err != nil {
        log.Println("ERROR: Failed to get market data")
        return nil
    } // if

    json["0"] = bars[0]
    json["1"] = bars[1]
    json["2"] = bars[2]
    json["3"] = bars[3]
    json["4"] = bars[4]

    

    for i := range data.OHLCVDelta {
        name := fmt.Sprintf("%d", i + FixedCols)

        json[name] = data.OHLCVDelta[i]
    } // for

    return json
} // PutPrices

// PutNewTechnicals inserts the new Technical Values after updating 
// Values of the Indicators that are not Diff, or Delta
func PutNewTechnicals(data *engine.UserData, json map[string]any) (map[string]any) {

    for i, ind := range data.Objects {

         name := fmt.Sprintf("%d", i + FixedCols * 2)

        switch v := ind.(type) {

        case *engine.SMA:
            json[name] = v.Data[len(v.Data) - 1]
        case *engine.EMA:
            json[name] = v.Data[len(v.Data) - 1]
        case *engine.Delta:
            json[name] = v.Value
        case *engine.Diff:
            json[name] = v.Value
        } // switch

    } // for

    return json
   
} // PutNewTechnicals


// GetLatestData returns back a JSON representation of the lastest values in 
// order as the defined JSON in src/logic/features.json
func MakeMLPayload(obj *engine.UserData, ticker string) (res map[string]any, err error) {
    var json map[string]any = make(map[string]any)
    json = PutPrices(obj, json, ticker)  // Put OHCLV Values
    json = PutNewTechnicals(obj, json) 
    return json, nil
} // Construct

// SendData sends data to the shared ML API to give updated
// Data
func SendData(obj *engine.UserData, ticker string) error {

    data, err := MakeMLPayload(obj, ticker)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(data)

    json, err := json.Marshal(data)
    if err != nil {
        log.Fatal(err)
    } // if

    resp, err := http.Post(ServerLink, "application/json", bytes.NewBuffer(json))
    if err != nil {
        log.Fatal(err)
    } // if

    defer resp.Body.Close()

    if resp.Status == "200" {
        return nil
    } else {
        return err
    }
} // SendData

// GetPrediction Gets the prediction back from the ML API   
// to determine the decision
func GetPrediction() (float64) {

    var result map[string]any = make(map[string]any)

    resp, err := http.Get(ServerLink)
    if err != nil {
        log.Println("ERROR: Could not get ML prediction")
    } // if
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("ERROR: Could not get ML prediction")
    } // if

    json.Unmarshal(body, &result)

    prediction, ok := result["prediction"].(float64)
    if !ok {
        log.Println("ERROR: Could not get ML prediction")
    } // ok


    return prediction

} // GetPrediction