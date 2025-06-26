package engine

// File Containing Methods for the EMA Indicator

import (
	"fmt"
)

// NewEMA makes a new EMA reference	
func NewEMA(json map[string]any) (*EMA, error) {
	window, ok := json["window"].(float64)
	if !ok {
		return nil, fmt.Errorf("window field should be an int")
	} // if

	win_int := int(window)

	smoothing, ok := json["smoothing"].(float64)
	if !ok {
		return nil, fmt.Errorf("smoothing field should be an int")
	} // if

	smoothInt := int(smoothing)

	return &EMA{
		Window: win_int, 
		Smoothing: smoothInt,
		Data: nil, 
		Alpha: 0,
	}, nil
} // NewEMA

func (ema *EMA) Load() error {
	originalLength := len(ema.Data)
	window := ema.Window

	if window > originalLength {
		return fmt.Errorf("window larger than array size")
	} // for

	firstEMA := findSMA(ema.Data, window)

	alpha := float64(ema.Smoothing) / float64(1+window)
	ema.Alpha = alpha

	emas := make([]float64, 0, originalLength-window+1)
	emas = append(emas, firstEMA)

	previousEMA := firstEMA
	for i := window; i < originalLength; i++ {
		price := ema.Data[i]
		newEMA := (price * alpha) + (previousEMA * (1 - alpha))
		emas = append(emas, newEMA)
		previousEMA = newEMA
	} // for

	ema.Data = emas
	return nil
} // Load (EMA)


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

// GetData (EMA) Gets the Data stored in its Data field given an index
func (ema *EMA) GetData(index int) (float64, error) {
		size := len(ema.Data) 
		if index >= size || index < 0 {
			return 0.0, fmt.Errorf("invalid index %d for len %d", index, size)
		} // if

		return ema.Data[index], nil
} // GetData (EMA)

// findSMA intializes the EMA calculation by finding the first SMA value
func findSMA(prices []float64, window int) float64 {

	var sum float64

	for i := range window {
		sum += prices[i]
	} // for

	return float64(sum) / float64(window)
} // findSMA