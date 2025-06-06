package engine

// file containing JSON related methods during runtime to interact
// with APIs and files in src/logic

import (
	"encoding/json"
	"fmt"
	"os"
)

// InitUserLogic intitializes the JSON files founds in src/logic as
// an array of technical indicators
func InitUserLogic(file string) (LiveIndicator, error) {
	var userArray LiveIndicator;

	userJSON, err := ParseLogicJSON(file)
	if err != nil {
		return userArray, fmt.Errorf("a problem occured when parsing the JSON file in src/logic")
	} // if

	userArray, err = LoadIndicators(userJSON)
	if err != nil {
		return userArray, fmt.Errorf("a problem occured when parsing the JSON file")
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
func LoadIndicators(json []map[string]any) (LiveIndicator, error) {

	live := LiveIndicator {
		Ind: []Indicator{},
		Techs: []string{},
		ColNames: []string{},
	}

	for i := range json {
		
		indicator, err := decideConstructor(&live, json[i])
		if err != nil {
			return live, fmt.Errorf("constructor failed at index %d", i)
		}      
		if indicator != nil {
			live.Ind = append(live.Ind, indicator)
		} 
		
	} // for

	return live, nil
} // LoadIndicators
func decideConstructor(li *LiveIndicator, json map[string]any) (Indicator, error) {

	indicator, ok := json["tech"].(string)

	if !ok {
		return nil, fmt.Errorf("tech field should be a string")
	} // if

	colName, ok := json["name"].(string)
	if !ok {
		return nil, fmt.Errorf("name field should be a string")
	}

	if indicator == "EMA" {
		li.Techs = append(li.Techs, "EMA")
		li.ColNames = append(li.ColNames, colName)
		return NewEMA(json)
	} else if indicator == "SMA" {
		li.Techs = append(li.Techs, "SMA")
		li.ColNames = append(li.ColNames, colName)
		return NewSMA(json)
	} else if indicator == "delta" {
		li.Techs = append(li.Techs, "DELTA")
		li.ColNames = append(li.ColNames, colName)
		return NewDelta(json)
	} else if indicator == "diff" {
		li.Techs = append(li.Techs, "DIFF")
		li.ColNames = append(li.ColNames, colName)
		return NewDiff(json)
	} else {
		return nil, fmt.Errorf("invalid technical indicator field")
	} // if-else

	/*


		Determine how we will decide delta and based on what ? (use col names)

		OK HERES WHAT WE DO MAKE THE TECHS JUST BE UP THE COLUMN NAME NOT THE TECH 
		OR JUST ADD A NEW STRING ARRAY THAT STORES RAW COL NAMES AND WE CAN USE 
		THE DELTAS THERE AS THE INDEX VALUE AS TECHS TO THE ACTUAL OBJECTS


	*/

} // decideConstructor
