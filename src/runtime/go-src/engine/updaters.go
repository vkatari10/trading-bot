package engine

// This file conatains methods to update technical indicators given new
// prices assumes calling associated methods in loaders.go had no errors


// CopySlice copies the Slice to reduce the capacity 
// of the underlying array
func CopySlice(slice []float64) []float64 {

	copied := make([]float64, len(slice))
	copy(copied, slice)
	return copied

} // CopySlice

// GetNew gets the new SMA value given a new price and
// appends it to the SMA data field
func (sma *SMA) GetNew(newPrice float64) {
	if cap(sma.Data) > sma.Window * 2 {
		sma.Data = CopySlice(sma.Data)
	} // if

	// minus oldest value, add back newest Price
	newSum := sma.Sum - sma.Data[0] + newPrice  
	sma.Sum = newSum

	// find newest SMA value of this window
	newSma := newSum / float64(sma.Window)

	// clear first entry, append new SMA
	sma.Data[0] = 0.0
	sma.Data = sma.Data[1:]
	sma.Data = append(sma.Data, newSma)
} // GetNew (SMA)

// GetNew gets the new EMA given a new price and 
// appends it to the EMA data Field
func (ema *EMA) GetNew(newPrice float64) {
	if cap(ema.Data) > ema.Window * 2 {
		ema.Data = CopySlice(ema.Data)
	} // if

	old_ema := ema.Data[len(ema.Data) - 1]
	
	new_ema := float64((newPrice * ema.Alpha) + 
	((1- ema.Alpha) * old_ema))

	// Drop oldest EMA value in the window to keep len(Data)
	// constant size
	ema.Data[0] = 0.0 
	ema.Data = ema.Data[1:]
	ema.Data = append(ema.Data, new_ema)

} // GetNew (EMA)

// GetNew (Delta) Gets the new 	
func (delta *Delta) GetNew(tech *Indicator) {

	

} // GetNew (Delta)