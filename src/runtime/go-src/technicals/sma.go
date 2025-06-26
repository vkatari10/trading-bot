package engine

// File that contains logic for the SMA Indicator

import (
	"fmt"
)

// NewSMA makes a new SMA reference 
func NewSMA(json map[string]any) (*SMA, error) {
	window, ok := json["window"].(float64)
	if !ok {
		return nil, fmt.Errorf("window field should be an int")
	} // if

	win_int := int(window)

	return &SMA{
		Window : win_int,
		Data: nil, 
		Sum: 0.0,
	}, nil
} // NewSMA

// Load Loads initial SMA values onto the Data field of the SMA struct
func (sma *SMA) Load() (error) {

	originalLength := len(sma.Data)
	window := sma.Window

	if window > originalLength {
		sma.Data = nil
		return fmt.Errorf("window larger than array size")
	} // if

	sum := 0.0

	for i := range window {
		sum += sma.Data[i]
	} // for

	final_len := originalLength - sma.Window + 1

	var smas []float64 = make([]float64, final_len)

	for i := sma.Window; i < originalLength; i++ {
		smas[i-window] = sum / float64(window)
		sum -= sma.Data[i-window]
		sum += sma.Data[i]
	} // for

	smas[final_len - 1] = sum / float64(window) // do final

	sma.Data = smas
	sma.Sum = sum

	return nil
} // Load (SMA)

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

// GetData (EMA) Gets the Data stored in its Data field given an index
func (sma *SMA) GetData(index int) (float64, error) {
		size := len(sma.Data)
		if index >= size || index < 0 {
			return 0.0, fmt.Errorf("invalid index %d for len %d", index, size)
		} // if

		return sma.Data[index], nil
} // GetData (SMA)