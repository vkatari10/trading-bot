package engine

import "fmt"

// this file contains methods to compute technicals in real time

// To Connect To C library in future (use import C)
/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data -lm
*/
// import "C"

//import "fmt"

// LoadBurnData loads the burn-in data into every technical
// indicator and calls the Load() method for each indicator
func LoadBurnData(obj *UserData, burn []float64) error {

	// TODO: ADD go routines since burn is always read only (no mutex needed)

	for _, ind := range obj.Objects {
		switch v := ind.(type) {
		case *SMA:
			v.Data = burn // put burn data as the SMA's data
			err := v.Load() // initialize SMA values based on burn data
			if err != nil {
				return fmt.Errorf("%v", err)
			} // if
		case *EMA:
			v.Data = burn	
			err := v.Load()
			if err != nil {
				return fmt.Errorf("%v", err)
			} // if
		case *Diff:
			v.Load(obj)
		case *Delta:
			v.Load(obj)
		default:
			v = nil
		} // swtich
	} // for

	return nil 
} // AssertType

// UpdateTechnicals updates the current technical indicators
// given a new price from the market API
func UpdateTechnicals(obj *UserData, newPrice float64) {
	for _, ind := range obj.Objects {
		switch v := ind.(type) {
		case *SMA:
			v.GetNew(newPrice)
		case *EMA:
			v.GetNew(newPrice)
		case *Delta:
			v.GetNew(obj)
		case *Diff:
			v.GetNew(obj)
		} // switch
	} // for
} // UpdateTechnicals

// UpdateOHLCVDeltas Updates the Deltas for OHCLV bars that all 
// Dataframes at train time contain
func UpdateOHLCVDeltas(obj *UserData, json [5]float64) {
	//fmt.Printf("INCOMING PRICE JSON -> %v\n", json)

	for i := 0; i < 5; i++ {
		newVal := json[i] - obj.OHLCVRaw[i] 
		obj.OHLCVDelta[i] = newVal
		obj.OHLCVRaw[i] = json[i]
	} // for
	//fmt.Printf("OHLCV DELTAS --> %v\n", obj.OHLCVDelta)
} // UpdateOHLCVDeltas



