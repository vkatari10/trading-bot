package engine

// this file contains methods to compute technicals in real time

// To Connect To C library in future (use import C)
/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data -lm
*/

// get user JSON
// load onto a LiveIndicator struct (
// burn in data for 30 minutes on a float32 array

// ====================== ^^^ DONE ||| vvv WIP =============================


// cycle that needs to take in a new price every 60 seconds and call update
// on each element of the LiveIndicator Array (MAKE METHOD)

// send back as JSON to ML model somehow and get prediction back from the
// Flask local server (MAKE METHODS)

// if we get buy sig utilze broker commands (just 1 share for now)

// need to do something to export this info for the bot API to expose to the frontend

// LoadBurnData loads the burn-in data into every technical 
// indicator and calls the Load() method for each indicator
func LoadBurnData(obj *UserData, burn []float64) {

	// TODO: ADD go routines since burn is always read only (no mutex needed)

	// for i := range obj.ObjectNames {

	// 	if obj.ObjectNames[i] == "SMA" {

	// 		sma, ok := obj.Objects[i].(*SMA)

	// 		if !ok {
	// 			obj.Objects[i] = nil
	// 		}

	// 		sma.Data = burn // put burn data as the SMA's data
	// 		sma.Load() // initialize SMA values based on burn data
				
	// 	} else if obj.ObjectNames[i] == "EMA" {

	// 		ema, ok := obj.Objects[i].(*EMA) 

	// 		if !ok {
	// 			obj.Objects[i] = nil
	// 		}

	// 		ema.Data = burn
	// 		ema.Load()

	// 	} else {
	// 		obj.Objects[i] = nil
	// 	}
	// } // for
} // AssertType

// // UpdateTechnicals updates the current technical indicators
// // given a new price from the market API
// func UpdateTechnicals(obj *LiveIndicator, newPrice float64) {
// 	for i := range obj.ObjectNames {
// 		obj.Ind[i].GetNew(newPrice)
// 	} // for
// } // UpdateTechnicals

//	
func UpdateOHLCVDeltas (obj *UserData, json map[string]float64) {
	return 
} // UpdateOHLCVDeltas