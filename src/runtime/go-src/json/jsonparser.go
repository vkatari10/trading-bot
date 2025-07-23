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
	Labels 		[]technicals.Relationship 		`json:"label_logic_2"`
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

	// features
	res.TALIBFeatureTechnicals = make([]technicals.FeatureTechnical, 0)
	res.OtherFeatureTechnicals  = make([]technicals.FeatureTechnical, 0)
	res.Relationships = make([]technicals.Relationship, 0)

	res.ColNames = make(map[string]int, 0)

	talibCount := 0
	otherCount := 0

	for i := range data.Features {
		tech := data.Features[i].Technical
		name := data.Features[i].Name

		if tech == "delta" || tech == "diff" {
			res.OtherFeatureTechnicals = append(res.OtherFeatureTechnicals, data.Features[i])
			res.ColNames[name] = otherCount
			otherCount++
		} else {
			res.TALIBFeatureTechnicals = append(res.TALIBFeatureTechnicals, data.Features[i])
			res.ColNames[name] = talibCount
			talibCount++
		}

	} 
	for i := range data.Labels {
		res.Relationships = append(res.Relationships, data.Labels[i])
		// res.ColNames[data.Labels[i].Name] = i + len(data.Features)
	}

	res.FeatureJSON = map[string]float64{}
	res.FeatureArray = make([]float64, len(data.Labels) + len(res.TALIBFeatureTechnicals) + len(res.OtherFeatureTechnicals))

	res.OHLCV = technicals.TALIBWrapper{} // init during burn in period

	// let the max capacity have -2 from limit to prevent unwanted GC
	res.OHLCV.SliceMaxCap = res.RuntimeSettings.BurnTime * technicals.CapLimitMultiplier - 2 
	
	return &res, nil
} // NewRuntimeData