package tests

import (
	"testing"
	"fmt"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
)

// TestNewRuntimeData checks if given a json file the object representation 
// matches what was in the JSON file 
func TestNewRuntimeData(t *testing.T) {

	data, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("TestNewRuntimeData failed to initialize")
	}

	// wantColNames := []string{"SMA_10", "SMA_30", "SMA_10_delta", 
	// "SMA_30_10_diff", "SMA_10_30_delta_diff", "SMA_CROSS"}

	
	
	wantTickerNames := []string{"AAPL", "AMZN", "BA"}

	wantCycleTime := 15.0
	wantOverride := true



	// for i := range 6 {
	// 	if data.ColNames[wantColNames[i]] != i {
	// 		t.Errorf("TestNewRuntimeData object name at index %d failed", i)
	// 	}
	// }

	for i := range 3 {
		if data.Tickers[i] != wantTickerNames[i] {
			t.Errorf("TestNewRuntimeData ticker does not match expected at index %d", i)
		}
	}

	if data.RuntimeSettings.CycleTime != wantCycleTime || data.RuntimeSettings.OverrideBurnIn != wantOverride {
		t.Errorf("TestNewRuntimeData runtime settings do not match expected values")
	}

	fmt.Println(data.Relationships)
	if len(data.Relationships) == 0 {
		t.Errorf("TestNewRuntimeData relationships array was not initialized properly")
	}

	if len(data.OtherFeatureTechnicals) == 0 {
		t.Errorf("TestNewRuntimeData other feature technicals array was not initialized properly")
	}

	if len(data.TALIBFeatureTechnicals) == 0 {
		t.Errorf("TestNewRuntimeData TALIB features technicals were not initialized properly")
	}
}
