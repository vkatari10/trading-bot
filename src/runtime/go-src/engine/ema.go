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

	smoothing, ok := json["smoothing"].(int)
	if !ok {
		return nil, fmt.Errorf("smoothing field should be an int")
	} // if

	return &EMA{
		Window: win_int, 
		Smoothing: smoothing,
		Data: nil, 
		Alpha: 0,
	}, nil
} // NewEMA

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

// Load method for Delta objects, matches their column
// indexes to the UserData.Objects slice
func (delta *Delta) Load(data *UserData) (error) {
	// Match the col index to the object array index
	delta.Col1Index = data.ColNames[delta.Col1]

	if delta.Col2 != "" { // if col2 isn't null in JSON
		delta.Col2Index = data.ColNames[delta.Col2]
	} else {  // if col2 was null in JSON
		delta.Col2Index = -1
	} // if-else

	return nil
} // Load (Delta)

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

