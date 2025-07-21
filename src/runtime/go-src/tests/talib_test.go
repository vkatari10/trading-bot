package tests 

import (
	"testing"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	eventloop "github.com/vkatari10/trading-bot/src/runtime/go-src/eventloop"
)

func TestTalibSMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

	res, err := technicals.TA_SMA( // ~10-20 microseconds w/ len 31 window
		runtimeData.Objects[0],
		&runtimeData.OHLCV,
	)

	if err != nil {
		t.Errorf("TA SMA failed")
	} 

	if res[0] <= 0 {
		t.Errorf("unexpected SMA value returned")
	}

}