package engine

// Parses User Config Features To Construct the technicals.UserData Object

import (
	"encoding/json"
	"fmt"
	"os"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

type constructor func(json map[string]any) (technicals.Indicator, error)

var (
	constructorDispatcher map[string]constructor 
)

// init declares the dispatch table for decideConstructor can call
func init() {
	constructorDispatcher = map[string]constructor{
		"EMA": technicals.NewEMA,
		"SMA": technicals.NewSMA,
		"delta": technicals.NewDelta,
		"diff": technicals.NewDiff,
	}
}

// ParseLogicJSON parses the JSONs files found in src/logic
func ParseLogicJSON(file string) (technicals.UserData, error) {

	jsonData, err := os.ReadFile(file)
	if err != nil {
		return technicals.UserData{}, fmt.Errorf("%v", err) // figure how else to handle this later another way
	} // if

	var jsonMap map[string]any

	err = json.Unmarshal(jsonData, &jsonMap)

	if err != nil {
		return technicals.UserData{}, fmt.Errorf("%v", err)
	} // if

	//fmt.Printf("json map -> %v\n", jsonMap)

	features, err := extractFeatures(jsonMap)
	if err != nil {
		return technicals.UserData{}, fmt.Errorf("%v", err)
	} // if

	return features, nil
} // ParseLogicJSON

// GetTradeTickers Gets the Tickers to trade on during live execution 
// from the user JSON config file
func GetTradeTickers(file string) ([]string, error){
	jsonData, err := os.ReadFile(file)
	if err != nil {
		return []string{}, fmt.Errorf("%v", err) // figure how else to handle this later another way
	} // if

	var jsonMap map[string]any

	err = json.Unmarshal(jsonData, &jsonMap)

	if err != nil {
		return []string{}, fmt.Errorf("%v", err)
	} // if

	tickers, ok := jsonMap["live_trade_stocks"].([]any)
	if !ok {
		return []string{}, fmt.Errorf("%v", ok)
	} // if

	tickerStrings := make([]string, 0) 

	for i := range tickers {

		tickerStr, ok := tickers[i].(string) 
		if !ok {
			return []string{}, fmt.Errorf("%v", ok) 
		} // if

		tickerStrings = append(tickerStrings, tickerStr)
	} // for

	return tickerStrings, nil
} // GetTradeTickers



func extractFeatures(config map[string]any) (technicals.UserData, error) {

	features, ok := config["features"].([]any)
	if !ok {
		return technicals.UserData{}, fmt.Errorf("%v", ok)
	} // if

	// User Data Object
	data := technicals.UserData {
		ColNames: map[string]int{},
		Objects: []technicals.Indicator{},
		OHLCVDelta: [5]float64{},
		OHLCVRaw: [5]float64{},
	} // UserData

	for i := range features {
		
		feature, ok := features[i].(map[string]any)
		if !ok {
			return technicals.UserData{}, fmt.Errorf("%v", ok)
		} // if 

		decideConstructor(&data, feature, i)	
	} // for 

	return data, nil
} // extractFeatures



// decideConstructor calls constructors based on each JSON objected 
// defined in the features.json file
func decideConstructor(data *technicals.UserData, json map[string]any, i int) (error) {

	indicator, ok := json["tech"].(string)
	if !ok {
		return fmt.Errorf("tech field should be a string")
	} // if

	colName, ok := json["name"].(string)
	if !ok {
		return fmt.Errorf("name field should be a string")
	} // if

	obj, err := constructorDispatcher[indicator](json)
	if err != nil {
		return fmt.Errorf("failed to construct object index %d (JSON feature not recognized) (%v)", i, err)
	} // if

	data.Objects = append(data.Objects, obj)

	/*

	This is for the diff and delta objects 

	Why is it str:int i have no idea 

	TODO: FIGURE OUT WHY IT LIKE THIS
	*/	
	data.ColNames[colName] = i // store index for the colName 
	
	return nil
} // decideConstructor
