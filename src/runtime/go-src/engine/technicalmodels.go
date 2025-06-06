package engine

// Contains model outlines and struct definitions

// Indicator interface for all technical indicators
type Indicator interface{
} // Inidicator

// Used at runtime to store user indicaotors
type LiveIndicator struct {
	Ind 		[]Indicator // All indicators are stored here
	Techs 		[]string // All 'tech' names are stored here (from JSON)
	ColNames 	[]string // All 'name' names are stored here 
} // LiveIndicators

// SMA Simple Moving Average indicator 
type SMA struct {
	Window 		int
	Data   		[]float64
	Sum			float64
	Delta		bool
} // SMA

// EMA Exponential Moving Average Indicator
type EMA struct {
	Window    	int
	Smoothing 	int
	Data      	[]float64
	Alpha  		float64 // intialized when Load() is called
	Delta 		bool
} // EMA

// Delta Represents JSON objects with the tech of 'delta'
type Delta struct {
	Col1		string // should store the index of the actual technical indicator in .Ind
	Col2		string
	Value 		float64
} // Delta

// Diff Represents JSON objects with the tech of 'diff
type Diff struct {
	Col1		string
	Col2		string
	Value		float64
} // Diff