package tests

import (
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine" 
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals" 
	"testing"
	"fmt"
)


// TestParseLogicJSON verifies if the JSON defined in config files can be properly
// read 
func TestParseLogicJSON(t *testing.T) {

	get, err := engine.ParseLogicJSON(fileName)
	if err != nil {
		t.Errorf("TestParseLogicJSON faied with error %v", err)
	} // if

	for i, ind := range get.Objects {

		switch v := ind.(type) {

		case *technicals.SMA:
			window := (v.Window == 20)
			if i != 0 && !window {
				t.Error(makeErrorString("SMA"))
			} // if
		case *technicals.EMA:
			window := (v.Window == 20)
			smoothing := (v.Smoothing == 2)
			if i != 1 || !window || !smoothing {
				t.Error(makeErrorString("EMA"))
			} // if
		case *technicals.Delta:
			colChecks := (v.Col1 == "SMA_20")
			
			if !colChecks {
				t.Error(makeErrorString("Delta"))
			} // if
		case *technicals.Diff:
			col1Check := (v.Col1 == "SMA_20")
			col2Check := (v.Col2 == "EMA_20")

			if !col1Check || !col2Check {
				t.Error(makeErrorString("Diff"))
			} // if
		default:
			t.Errorf("TestParseLogicJSON failed, unknown tech field")
		}
	}

	if len(get.ColNames) != len(get.Objects) {
		t.Errorf("TestParseLogicJSON failed to read all column names")
	} // if 
} // TestParseLogicJSON

func TestGetTradeTickers(t *testing.T) {

	get, err := engine.GetTradeTickers(fileName)
	if err != nil {
		t.Errorf("TestGetTradeTickers failed with error %v", err)
	} // if 

	want := []string{"AAPL", "AMZN"}

	for i := range get {
		if get[i] != want[i] {
			t.Errorf("TestGetTradeTickers failed (ticker %s != %s)", get[i], want[i])
		} // if
	} // for
} // TestGetTradeTickers

func makeErrorString(technical string) string {
	return fmt.Sprintf("TestParseLogicJSON failed to read %s properly", technical)
} // makeErrorString