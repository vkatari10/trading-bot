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

// NewDelta makes a new Delta reference
func NewDelta(json map[string]any) (*Delta, error) {
	col1, ok := json["col1"].(string)
	if !ok {
		return nil, fmt.Errorf("col1 should be a string")
	} // if

	col2, ok := json["col2"].(string)
	if !ok {
		return nil, fmt.Errorf("col2 should be a string")
	} // if

	return &Delta {
		Col1: col1,
		Col2: col2,
		Value: 0,
	}, nil
} // NewDelta

// NewDiff makes a new Diff reference
func NewDiff(json map[string]any) (*Diff, error) {
	col1, ok := json["col1"].(string)
	if !ok {
		return nil, fmt.Errorf("col1 should be a string")
	} // if

	col2, ok := json["col2"].(string)
	if !ok {
		return nil, fmt.Errorf("col2 should be a string")
	} // if

	return &Diff {
		Col1: col1,
		Col2: col2,
		Value: 0,
	}, nil
} // NewDelta