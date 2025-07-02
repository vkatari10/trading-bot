package eventloop

import (
	"fmt"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)
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
func LoadBurnData(obj *technicals.UserData, burn []float64) error {

	// TODO: ADD go routines since burn is always read only (no mutex needed)

	for _, ind := range obj.Objects {
		switch v := ind.(type) {
		case *technicals.SMA:
			v.Data = burn // put burn data as the SMA's data
			err := v.Load() // initialize SMA values based on burn data
			if err != nil {
				return fmt.Errorf("%v", err)
			} // if
		case *technicals.EMA:
			v.Data = burn	
			err := v.Load()
			if err != nil {
				return fmt.Errorf("%v", err)
			} // if
		case *technicals.Diff:
			v.Load(obj)
		case *technicals.Delta:
			v.Load(obj)
		default:
			v = nil
		} // swtich
	} // for

	return nil 
} // AssertType

// UpdateTechnicals updates the current technical indicators
// given a new price from the market API
func UpdateTechnicals(obj *technicals.UserData, newPrice float64) {
	for _, ind := range obj.Objects {
		switch v := ind.(type) {
		case *technicals.SMA:
			v.GetNew(newPrice)
		case *technicals.EMA:
			v.GetNew(newPrice)
		case *technicals.Delta:
			v.GetNew(obj)
		case *technicals.Diff:
			v.GetNew(obj)
		} // switch
	} // for
} // UpdateTechnicals

// UpdateOHLCVDeltas Updates the Deltas for OHCLV bars that all 
// Dataframes at train time contain
func UpdateOHLCVDeltas(obj *technicals.UserData, json [5]float64) {
	//fmt.Printf("INCOMING PRICE JSON -> %v\n", json)

	for i := 0; i < 5; i++ {
		newVal := json[i] - obj.OHLCVRaw[i] 
		obj.OHLCVDelta[i] = newVal
		obj.OHLCVRaw[i] = json[i]
	} // for
	//fmt.Printf("OHLCV DELTAS --> %v\n", obj.OHLCVDelta)
} // UpdateOHLCVDeltas



