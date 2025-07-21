package eventloop

import (
	"testing"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
)

// Burn In Methods


// TestOverrideBurnIn checks if the arrays are initialized 
// when calling in a burn in method
func TestOverrideBurnIn(t *testing.T) {
	gotConfig, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("%v", err)
	}

	OverrideBurnIn(gotConfig)

	if len(gotConfig.OHLCV.High) < gotConfig.RuntimeSettings.BurnTime {
		t.Errorf("TestOverrideBurnIn: arrays were not initialized")
	}
}

// TestRuntimeDataStress repeatedly pops and append data
// to the OHLCV arrays to check if slice capacity, length, 
// and pointers are not leading to mem leaks or segfaults
func TestRuntimeDataStressed(t *testing.T) {
	gotConfig, err := json.NewRuntimeData("../test.json")
	if err != nil {	
		t.Errorf("%v", err)
	}

	OverrideBurnIn(gotConfig)

	

	for range 1_000_000 { // 0.05s @ CapLimitMultiplier = 2
		gotConfig.PopLeft()
		gotConfig.TestAppend(5.0)
		gotConfig.UpdateDeltas()
	}
	
	checkArr := gotConfig.OHLCV.Close

	if len(checkArr) > 31 && cap(checkArr) > 62 {
		t.Errorf("cap or len grew beyond permitted size")
	}
}

