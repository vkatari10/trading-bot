package engine

// Contains model outlines and struct definitions


/*
Each technical Indicator requires these methods 

1. Load() -- Initalizes Values based on burn in data
2. GetNew() -- Updates Value based on a new Price
3. GetData() -- (Optiona) If struct contains an array to access data values

These should be reciever methods
*/

// Indicator interface for all technical indicators
type Indicator interface{
	Tag() 			// Placeholder method
	Type() string 	// For Type Assertions
} // Inidicator

// UserData contains user data needed at runtime
type UserData struct {
	ColNames 	map[string]int // Raw Col Names and index from user JSON
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
	Col1Index 	int
	Col2Index 	int
	Value		float64
} // Diff

// Decleration of Dummy Method
func (SMA) 		Tag() {}
func (EMA) 		Tag() {}
func (Delta) 	Tag() {}
func (Diff) 	Tag() {}

// Type() Implementations
func (SMA)		Type() (string) {return "SMA"}
func (EMA) 		Type() (string) {return "EMA"}
func (Delta)	Type() (string) {return "DELTA"}
func (Diff) 	Type() (string) {return "DIFF"}
