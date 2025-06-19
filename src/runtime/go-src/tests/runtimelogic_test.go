package tests

import (
	"fmt"
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

	testBurnData := []float64{
	82.13, 57.84, 96.47, 51.29, 88.11, 74.25, 59.63, 92.05, 68.74, 85.97,
    60.22, 78.49, 53.38, 90.16, 71.02, 66.44, 55.87, 83.33, 72.58, 99.44,
	56.56}

	engine.LoadBurnData(&userData, testBurnData)

	expected := []float64{}

	expected = append(expected, 73.09) // SMA
	expected = append(expected, 72.68) // EMA


	for i, ind := range userData.Objects {	
		switch v := ind.(type) {
		case *engine.SMA:
			fmt.Println(v.Data)
			err := checkEquivalent(expected[i], v.Data[0])
			if err != nil {
				t.Errorf("TestLoadBurnData failed for object index %d for type %s (%v)", i, "SMA", err)
			}
		case *engine.EMA:
			err := checkEquivalent(expected[i], v.Data[0])
			if err != nil {
				t.Errorf("TestLoadBurnData fialed for object index %d for type %s (%v)", i, "EMA", err)
			}
		} // swtich

	} // for

} // TestLoadBurnData

// checkEquivalent is a helper method that helps
// determine if 
func checkEquivalent(want float64, got float64) error {

	if want != got {
		return fmt.Errorf("Unequal values (want %.2f != got %.2f)", want, got)
	} // if

	return nil 
} // checkEquivalent