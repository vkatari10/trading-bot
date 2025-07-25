package tests

import (
	"testing"
	"time"
	"log"
	"math/rand"
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

	eventloop.OverrideBurnIn(&rd, "X")

	

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

	eventloop.OverrideBurnIn(&rd, "X")

	rd.InitRelationships()
	
	start := time.Now()
	for range 1_000_000 { // 1.7 microseconds per tick
		rd.UpdateOHLCVDeltas()

		err = rd.UpdateTALIBTechnicals()
		if err != nil {
			t.Errorf("TALIB computation failed")
		}

		err = rd.UpdateOtherTechnicals()
		if err != nil {
			t.Errorf("Delta/Diff comp failed")
		}

		err := rd.UpdateRelationships()
		if err != nil {
			t.Errorf("Relationships could not be updated")
		}

		rd.PopLeft()
		rd.TestAppend(rand.Float64() * 10) // dummy new market price 
	}
	end := time.Since(start)

	log.Println(end)
	log.Println(end / 1_000_000)
	
	if len(rd.OHLCV.High) > 31 || cap(rd.OHLCV.High) > 62 {
		t.Errorf("Cap or length of OHLCV arrays grew beyond permitted size")
	}
}

func TestInitRelationships(t *testing.T) {
	data, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("InitRelationships: Failed to construct runtime data object")
	}

	eventloop.OverrideBurnIn(&data, "X")
	data.UpdateTALIBTechnicals()
	data.UpdateOtherTechnicals()
	data.InitRelationships()

	if data.Relationships[0].Col1Val == 0 || data.Relationships[0].Col2Val == 0 {
		t.Errorf("InitRelationships: Failed to initialize relationship values")
	}

}