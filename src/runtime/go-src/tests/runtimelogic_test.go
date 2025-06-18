package tests

import (
	"testing"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine" 
)

// TestLoadBurnData verifies that burn in data is intialized 
// properly based on given burn data
func TestLoadBurnData(t *testing.T) {

	file := "testdata.json"

	userData, err := engine.InitUserLogic(file)
	if err != nil {
		t.Errorf("TestLoadBurnData failed with error %v", err)
	} // if

} // TestLoadBurnData