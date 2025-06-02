package engine

// file containing JSON related methods during runtime to interact 
// with APIs and files in src/logic

import (
	"os"
	"encoding/json"
	"fmt"
)

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
func LoadIndicators(json []map[string]any) ([]Indicator, error) {
	var indicators []Indicator = make([]Indicator, 0)

	for i := range json {

		indicator, err := decideConstructor(json[i])
		if err != nil {
			return nil, fmt.Errorf("technical indicator does" + 
			 "not exist, or is not supported")
		}
		if indicator != nil {
			indicators = append(indicators, indicator)
		} // if
		
	} // for

	return indicators, nil
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
