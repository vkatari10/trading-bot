package engine

// this file contains methods to read the technicals declared in
// src/logic to know what technicals need to be computed in real time

import (
	"fmt"
)

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

// Load Loads initial EMA values onto the Data field of the EMA struct
func (ema *EMA) Load() (error) {

	originalLength := len(ema.Data)
	window := ema.Window

	if window > originalLength {
		return fmt.Errorf("window larger than array size")
	} // if

	sum := 0.0

	for i := range window {
		sum += ema.Data[i]
	} // for

	final_len := originalLength - window + 1

	alpha := float64((ema.Smoothing) / (1 + window))
	ema.Alpha = alpha

	var emas []float64 = make([]float64, final_len)

	old_ema := sum / float64(window)

	emas = append(emas, old_ema)

	for i := 0; i <= originalLength; i++ {
		new_ema := (ema.Data[i] * alpha) + ((1 - alpha) * old_ema)
		emas[i-window+1] = new_ema
		old_ema = new_ema
	} // for

	return nil

} // Load (EMA)





