package engine

// this file contains methods to read the technicals declared in
// src/logic to know what technicals need to be computed in real time

import (
	"encoding/json"
	"os"
)

// ParseLogicJSON parses the JSONs files found in src/logic
func ParseLogicJSON(f string) ([]map[string]any, error) {

	file := "../../logic/"

	file += f

	json_data, err := os.ReadFile(file)
	if err != nil {
		return nil, err // figure how else to handle this later another way
	} // if

	var json_map []map[string]any

	err = json.Unmarshal(json_data, &json_map)

	if err != nil {
		return nil, err
	} // if

	return json_map, nil
} // ParseLogicJSON	