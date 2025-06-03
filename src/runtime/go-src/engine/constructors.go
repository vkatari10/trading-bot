package engine

// Contains methods to construct new technical indicators objects

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