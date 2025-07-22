package tests

import (
	"testing"

	eventloop "github.com/vkatari10/trading-bot/src/runtime/go-src/eventloop"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	//technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

// Tests if technicals are updated on calls
func TestRuntimeDataUpdateTechnicals(t *testing.T) {

	rd, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("Test RuntimeData Update Technicals failed to init runtime data obj")
	}

	eventloop.OverrideBurnIn(rd)

	err = rd.UpdateTALIBTechnicals()
	if err != nil {
		t.Errorf("RuntimeData.UpdateTechnicals() failed to run: %v", err)
	}

	err = rd.UpdateOtherTechnicals()
	if err != nil {
		t.Errorf("RuntimeData.UpdateOtherTechnicals() failed")
	}

}

func TestRuntimeDataStressTests(t *testing.T) {

	rd, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("Test RuntimeData Update Technicals failed to init runtime data obj")
	}

	eventloop.OverrideBurnIn(rd)
	
	for range 1_000_000 { // 65 microseconds per tick
		rd.UpdateOHLCVDeltas()
		err = rd.UpdateTALIBTechnicals()
		if err != nil {
			t.Errorf("TALIB computation failed")
		}

		err = rd.UpdateOtherTechnicals()
		if err != nil {
			t.Errorf("Delta/Diff comp failed")
		}

		rd.PopLeft()
		rd.TestAppend(5.0)
	}
	
	if len(rd.OHLCV.High) > 31 || cap(rd.OHLCV.High) > 62 {
		t.Errorf("Cap or length of OHLCV arrays grew beyond permitted size")
	}
}
