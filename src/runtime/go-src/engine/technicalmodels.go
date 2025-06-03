package engine

// Contains model outlines and struct definitions

// Indicator interface for all technical indicators
type Indicator interface {
	GetNew(newPrice float64) 	
	Load() 		error
} // Inidicator

// Used at runtime to store user indicaotors
type LiveIndicator struct {
	Ind 		[]Indicator
} // LiveIndicators

// SMA Simple Moving Average indicator 
type SMA struct {
	Window 		int
	Data   		[]float64
	Sum			float64
} // SMA

// EMA Exponential Moving Average Indicator
type EMA struct {
	Window    	int
	Smoothing 	int
	Data      	[]float64
	Alpha  		float64 // intialized when Load() is called
} // EMA
