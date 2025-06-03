package engine

// file containing JSON related methods during runtime to interact 
// with APIs and files in src/logic

import (
	"os"
	"encoding/json"
	"fmt"
)

// InitUserLogic intitializes the JSON files founds in src/logic as
// an array of technical indicators
func InitUserLogic(file string) (LiveIndicator, error) {
	var userArray LiveIndicator;

	userJSON, err := ParseLogicJSON(file)
	if err != nil {
		return userArray, fmt.Errorf("A problem occured when parsing the JSON file in src/logic")
	} // if

	userArray, err = LoadIndicators(userJSON)

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
	}

	for i := range json {

		indicator, err := decideConstructor(json[i])
		if err != nil {
			live.Ind = nil
			return live, fmt.Errorf("technical indicator does" + 
			 "not exist, or is not supported")
		}
		if indicator != nil {
			live.Ind = append(live.Ind, indicator)
		} // if
		
	} // for

	return live, nil
} // LoadIndicators

func decideConstructor(json map[string]any) (Indicator, error) {

	indicator, ok := json["name"].(string)

	if !ok {
		return nil, fmt.Errorf("name field should be a string")
	} // if

	if indicator == "EMA" {
		return NewEMA(json)
	} else if indicator == "SMA" {
		return NewSMA(json)
	} else {
		return nil, fmt.Errorf("invalid technical indicator field")
	} // if-else

} // decideConstructor
