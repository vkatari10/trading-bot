package json

// This file contains a method to obtain a RuntimeData object
// by parsing the user JSON file
//
// Author: Vikas Katari
// Date: 07/17/2025

import (
	"encoding/json"
	"fmt"
	"os"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

// ConfigSettings represents the main config settings 
type configSettings struct {
	LiveTickers []string 						`json:"live_trade_stocks"`
	Features 	[]technicals.FeatureTechnical 	`json:"features"`
	Labels 		[]technicals.Relationship 		`json:"label_logic"`
	Runtime 	technicals.RuntimeSettings 		`json:"runtime_settings"`
} // ConfigSettings

// printErrorMsg returns an error message to any method that fails
func printErrorMsg(method string, err error) error {
	return fmt.Errorf("%s failed with error: %w", method, err)
} // printErrorMsg

// GetConfigJSON converts the config file into Go usable objects
func getConfigJSON(file string) (*configSettings, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, printErrorMsg("GetConfigJSON", err)
	}

	var result configSettings
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, printErrorMsg("GetConfigJSON", err)
	}

	return &result, nil
} // GetConfigJSON

// NewRuntimeData gives a new RuntimeData object that represents all user
// defined config values as a Go object
func NewRuntimeData(file string) (*technicals.RuntimeData, error) {

	data, err := getConfigJSON(file)
	if err != nil {
		return nil, err
	} 

	var res technicals.RuntimeData 
	res.Tickers = data.LiveTickers // live tickers
	res.RuntimeSettings = data.Runtime // runtime Settings

	res.Objects = make([]technicals.Feature, 0) // Feature Objects
	res.ColNames = make(map[string]int, 0)
	for i := range data.Features {
		res.Objects = append(res.Objects, data.Features[i])
		res.ColNames[data.Features[i].Name] = i
	} 
	for i := range data.Labels {
		res.Objects = append(res.Objects, data.Labels[i])
		res.ColNames[data.Labels[i].Name] = i + len(data.Features)
	}

	res.OHLCV = technicals.TALIBWrapper{} // init during burn in period
	res.OHLCV.SliceMaxCap = res.RuntimeSettings.BurnTime * 2 - 2
	
	return &res, nil
} 

// GetRuntimeData initializes and returns the RuntimeData object 
func GetRuntimeData(json map[string]any, maxLookback int) (*technicals.RuntimeData) {

	// Init col names
	colNames := make(map[string]int, 0)
	


	// init feature payloads 
	featureJSON := make(map[string]float64, 0)
	featureArray := make([]float64, 0)

	// init object array of features
	features := make([]technicals.Feature, 0)

	return &technicals.RuntimeData{
		ColNames: colNames,
		Objects: features,
		FeatureJSON: featureJSON,
		FeatureArray: featureArray,
	}
} // GetRuntimeData	

