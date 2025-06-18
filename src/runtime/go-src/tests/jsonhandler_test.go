package tests

import (
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine" 
	"testing"
	"encoding/json"
	"os"
	"reflect"
	//"fmt"
)


// TestParseLogicJSON verifies if the JSON defined in config files can be properly
// read 
func TestParseLogicJSON(t *testing.T) {
	file := "testdata.json"

	get, err := engine.ParseLogicJSON(file)
	if err != nil {
		t.Errorf("TestParseLogicJSON faied with error %v", err)
	} // if

	want := []map[string]any{}
	want = append(want, map[string]any{
		"name": "SMA_25",
		"tech": "SMA",
		"window": 25,
	})
	want = append(want, map[string]any{
		"name": "EMA_20",
		"tech": "EMA",
		"window": 20,
	})

	// check if names and tech are equal which are the most important 
	for i := range want {
		check_name := assertString(want[i], "name") == assertString(get[i], "name")
		check_tech := assertString(want[i], "tech") == assertString(get[i], "tech")
		if check_name && check_tech {
			continue
		} else {
			t.Errorf("TestParseLogicJSON failed")
		}
	} // for
} // TestParseLogicJSON

// TestLoadIndicators verifies if LoadIndicators in engine.jsonhandler correctly
// initlaize a engine.UserData object
func TestLoadIndicators(t *testing.T) {

	testJSON, err := os.ReadFile("testdata.json")
	if err != nil {
		t.Fatalf("TestLoadIndicators failed with error %v", err)
	} // if 

	var jsonMap []map[string]any
	err = json.Unmarshal(testJSON, &jsonMap)
	if err != nil {
		t.Fatalf("TestLoadInidicators failed with error %v", err)
	} // if

	get, err := engine.LoadIndicators(jsonMap)
	if err != nil {
		t.Errorf("TestLoadIndicators failed with error %v", err)
	} // if 

	want_colNames := map[string]int{"SMA_25": 0, "EMA_20": 1}
	
	if !reflect.DeepEqual(get.ColNames, want_colNames) {
		t.Error("TestLoadIndicators colnames are incorrect")
	} // if
	
} // TestLoadInidicators





func assertString(json map[string]any, key string) string {
	assert, ok := json["key"].(string)
	
	if !ok {
		return ""
	} // if 

	return assert
} // parseString
 