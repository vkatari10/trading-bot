package engine

// file containing JSON related methods during runtime to interact
// with APIs and files in src/logic

import (
	"encoding/json"
	"fmt"
	"os"
)

// InitUserLogic intitiadatazes the JSON files founds in src/logic as
// an array of technical indicators
func InitUserLogic(file string) (UserData, error) {
	var userArray UserData;

	userJSON, err := ParseLogicJSON(file)
	if err != nil {
		return userArray, fmt.Errorf("could not read src/logic/features.json")
	} // if

	userArray, err = LoadIndicators(userJSON)
	if err != nil {
		return userArray, 
		fmt.Errorf("could not parse, check src/logic/features.json")
	}

	return userArray, nil
} // InitUserLogic

// ParseLogicJSON parses the JSONs files found in src/logic
func ParseLogicJSON(f string) ([]map[string]any, error) {

	file := "../../logic/"

	file += f

	jsonData, err := os.ReadFile(file)
	if err != nil {
		return nil, err // figure how else to handle this later another way
	} // if

	var jsonMap []map[string]any

	err = json.Unmarshal(jsonData, &jsonMap)

	if err != nil {
		return nil, err
	} // if
	return jsonMap, nil
} // ParseLogicJSON

// Loads technical indicators from the JSON onto an Indicator array
func LoadIndicators(json []map[string]any) (UserData, error) {

	data := UserData {
		ColNames: map[string]int{},
		Objects: []Indicator{},
		OHLCVDelta: [5]float64{},
	} // UserData

	decideConstructor(&data, json) // load all objects onto array

	return data, nil
} // LoadIndicators

// decideConstructor calls constructors based on each JSON objected 
// defined in the features.json file
func decideConstructor(data *UserData, json []map[string]any) (error) {

	for i := range json {

		indicator, ok := json[i]["tech"].(string)

		if !ok {
			return fmt.Errorf("tech field should be a string")
		} // if

		colName, ok := json[i]["name"].(string)
		if !ok {
			return fmt.Errorf("name field should be a string")
		} // if

		if indicator == "EMA" {
			ema, err := NewEMA(json[i])
			if err != nil {
				return fmt.Errorf("construction failed -> object index %d", i)
			} // if
			data.Objects = append(data.Objects, ema)
		} else if indicator == "SMA" {
			sma, err := NewSMA(json[i])
			if err != nil {
				return fmt.Errorf("construction failed -> object index %d", i)
			} // if
			data.Objects = append(data.Objects, sma)
		} else if indicator == "delta" {
			delt, err := NewDelta(json[i])
			if err != nil {
				return fmt.Errorf("construction failed -> object index %d", i)
			} // if
			data.Objects = append(data.Objects, delt)
		} else if indicator == "diff" {
			diff, err := NewDiff(json[i])
			if err != nil {
				return fmt.Errorf("construction failed -> object index %d", i)
			} // if
			data.Objects = append(data.Objects, diff)
		} else {
			return fmt.Errorf("\"tech\" field for object at index %d is not recognized", i)
		} // if-else

		data.ColNames[colName] = i // store index for the colName
		} // for 
		
	return nil
} // decideConstructor
