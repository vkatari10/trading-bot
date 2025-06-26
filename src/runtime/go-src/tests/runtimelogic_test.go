package tests

import (
	"fmt"
	"testing"
	"math"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

// TestLoadBurnData verifies that burn in data is intialized
// properly based on given burn data
func TestLoadBurnData(t *testing.T) {

	file := "future.json"

	userData, err := engine.ParseLogicJSON(file)
	if err != nil {
		t.Errorf("TestLoadBurnData failed with error %v", err)
	} // if

	testBurnData := getTestData()

	err = engine.LoadBurnData(&userData, testBurnData)
	if err != nil {
		t.Errorf("TestLoadBurnData failed with error %v", err)
	} // if

	expected := []float64{}

	expected = append(expected, 73.0919) // SMA
	expected = append(expected, 72.7296) // EMA

	for i, ind := range userData.Objects {	
		switch v := ind.(type) {
		case *technicals.SMA:
			err := checkEquivalent(expected[i], v.Data[len(v.Data) - 1])
			if err != nil {
				t.Errorf("TestLoadBurnData failed for object index %d for type %s (%v)", i, "SMA", err)
			} // if
		case *technicals.EMA:
			err := checkEquivalent(expected[i], v.Data[len(v.Data) - 1])
			if err != nil {
				t.Errorf("TestLoadBurnData fialed for object index %d for type %s (%v)", i, "EMA", err)
			} // if
		} // swtich
	} // for

} // TestLoadBurnData

// TestUpdateTechnicals tests if the updating methods for 
// updating technical indicators works as intended
func TestUpdateTechnicals(t *testing.T) {

	userData, err := engine.ParseLogicJSON("future.json")
	if err != nil {
		t.Errorf("TestUpdateTechnicals failed with error %v", err)
	} // if
	
	testBurnData := getTestData()

	err = engine.LoadBurnData(&userData, testBurnData)
	if err != nil {
		t.Errorf("TestUpdateTechnicals failed with error %v", err)
	} // if

	engine.UpdateTechnicals(&userData, 56.56)

	wants := make([]float64, 0)

	wants = append(wants, 72.20147) // SMA
	wants = append(wants, 71.13957) // EMA

	for i, ind := range userData.Objects {
		switch v := ind.(type) {
		case *technicals.EMA:
			err := checkEquivalent(wants[i], v.Data[len(v.Data) - 1])
			if err != nil {
				t.Errorf("TestUpdateTechnicals failed with %v", err)
			} // if
		case *technicals.SMA:
			err := checkEquivalent(wants[i], v.Data[len(v.Data) - 1])
			if err != nil {
				t.Errorf("TestUpdateTechnicals failed with %v", err)
			} // if
		} // switch 
	} // for

} // TestUpdateTechnicals


// checkEquivalent is a helper method that helps
// determine if 
func checkEquivalent(want float64, got float64) error {
	if floatsEqual(want, got, 1e-9) {
		return fmt.Errorf("Unequal values (want %.2f != got %.2f)", want, got)
	} // if
	return nil 
} // checkEquivalent

// floatsEqual checks if two floats are nearly equal 
func floatsEqual(a, b, epsilon float64) bool {
    return math.Abs(a-b) <= epsilon
} // floatsEqual

func getTestData() []float64  {
	return []float64{
	82.13, 57.84, 96.47, 51.29, 88.11, 74.25, 59.63, 92.05, 68.74, 85.97,
    60.22, 78.49, 53.38, 90.16, 71.02, 66.44, 55.87, 83.33, 72.58, 99.44,
	56.56}
} // getTestData()
