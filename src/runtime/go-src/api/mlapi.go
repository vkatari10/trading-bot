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
    FixedCols = 5
)

// PutPrices loads the intial close, high, low, and open prices that 
// the ML models was trained on (Yfinance includes these by deafult)
func PutPrices(json map[string]any, ticker string) (map[string]any) {
    bars := []string{"c", "h", "l", "o"} // close, high, low, open

    for i := range FixedCols {
        name := fmt.Sprintf("%d", i)
        val, err := GetQuote(ticker, bars[i])
        if err != nil {
            log.Println("ERROR: Failed to get quote for ML inference")
            return json
        } // if
        json[name] = val
    } // for

    return json
} // PutPrices

// GetLatestData returns back a JSON representation of the lastest values in 
// order as the defined JSON in src/logic/features.json
func GetLatestData(obj *engine.UserData, ticker string) (res map[string]any, err error) {

    var json map[string]any = make(map[string]any)

    json = PutPrices(json, ticker)  

    for i := range obj.Objects {
            name := fmt.Sprintf("%d", i + FixedCols)

            var insertValue float64;

            if obj.Objects[i].Type() == "SMA" {
                sma, ok := obj.Objects[i].(*engine.SMA) 
                if !ok {
                    return nil, err
                } // if

                // grab latest value
                insertValue = sma.Data[len(sma.Data) - 1]
            } else if obj.Objects[i].Type() == "EMA" {
                ema, ok := obj.Objects[i].(*engine.EMA)
                if !ok {
                    return nil, err
                } // if

                insertValue = ema.Data[len(ema.Data) - 1]
            } // if-else

            json[name] = insertValue
    } // for


    return json, nil

} // GetLatestData

// SendData sends data to the shared ML API to give updated
// Data
func SendData(obj *engine.UserData, ticker string) error {

    data, err := GetLatestData(obj, ticker)
    if err != nil {
        log.Fatal(err)
    }

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