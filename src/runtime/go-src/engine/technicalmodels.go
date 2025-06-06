package engine

// Contains model outlines and struct definitions

// Indicator interface for all technical indicators
type Indicator interface{
	// For the sake of including all 
	// technicals under a common interface

	Tag() // Placeholder method
	Type() string
} // Inidicator

// UserData contains runtime data 
type UserData struct {
	ColNames 	[]string // Raw Col Names from user JSON
	Objects  	[]Indicator // feature refs from user JSON "tech"s
	OHLCVDelta	[5]float64 // Store deltas for all 5 values
} // UserData

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
	Col1Index 	int
	Col2Index 	int
	Value 		float64
} // Delta

// Diff Represents JSON objects with the tech of 'diff
type Diff struct {
	Col1		string
	Col2		string
	Value		float64
} // Diff

// // Used at runtime to store user indicaotors
// type LiveIndicator struct {
// 	Ind 		[]Indicator // All indicators are stored here
// 	Techs 		[]string // All 'tech' names are stored here (from JSON)
// 	ColNames 	[]string // All 'name' names are stored here 
// } // LiveIndicators


// Decleration of Dummy Method
func (SMA) 		Tag() {}
func (EMA) 		Tag() {}
func (Delta) 	Tag() {}
func (Diff) 	Tag() {}

// Type Declerations
func (SMA)		Type() (string) {return "SMA"}
func (EMA) 		Type() (string) {return "EMA"}
func (Delta)	Type() (string) {return "DELTA"}
func (Diff) 	Type() (string) {return "DIFF"}
