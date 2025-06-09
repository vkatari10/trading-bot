package engine

// this file contains methods to compute technicals in real time

// To Connect To C library in future (use import C)
/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data -lm
*/

// need to do something to export this info for the bot API to expose to the frontend

// LoadBurnData loads the burn-in data into every technical 
// indicator and calls the Load() method for each indicator
func LoadBurnData(obj *UserData, burn []float64) {

	// TODO: ADD go routines since burn is always read only (no mutex needed)

	for _, ind := range obj.Objects {
		switch v := ind.(type) {
		case *SMA:
			v.Data = burn // put burn data as the SMA's data
			v.Load() // initialize SMA values based on burn data
		case *EMA:
			v.Data = burn	
			v.Load()
		default:
			v = nil
		} // swtich
	} // for
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
	for i := range 5 {
		obj.OHLCVDelta[i] = json[i] - obj.OHLCVDelta[i] 
	} // for
} // UpdateOHLCVDeltas



