package technicals

// Contains struct definitions for technical indicators

/*
Each struct will need at least 3 methods 

1. NewX() -- Constructs a new Object of a given struct
2. Load() -- Receiver method that initializes the value of a technical 
indicator given burn in data
3. GetNew() -- Receiver method that updates the value of a technical 
indicator given new data
4. GetData() -- (Optional) Gets the latest value of a technical indicator

In addition to satisfy the Indicator{} interface declare 
these two methods at the bottom of this file 
1. Tag() -- Placeholder method
2. Type() -- Return Type of Object as a String

Q: Why are the first 3 methods not declared in the interface? 

A: Not every technical indicator takes in the same data

It is up to the user to determine how a method should be defined and won't 
limit the method to param constraints. 
*/

// Indicator interface for all technical indicators
type Indicator interface{
	Tag() 			// Placeholder method
	Type() string 	// For Type Assertions
} // Indicator

// UserData contains user data needed at runtime
type UserData struct {
	ColNames 	map[string]int // Raw Col Names and index from user JSON
	Objects  	[]Indicator // feature refs from user JSON "tech"s
	OHLCVDelta	[5]float64 // Store deltas for all 5 values
	OHLCVRaw 	[5]float64 // Store raw values 
	// Might need to add historical volume and close price arrays as well???
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
	Alpha  		float64 // initialized when Load() is called
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

// MACD Represents a Moving Average Convergence Divergence
// technical indicator
type MACD struct {
	EMA1 EMA
	EMA2 EMA
	Diff float64 // may not be needed
} // MACD

// MACDsignal Represents a Moving Average Convergence Divergence
// Signal Line (To be used with the MACD Struct)
type MACDSignal struct {
	Signal EMA
} // MACDSignal

// BollingerUpper Represents the Upper Bollinger Band Technical
// indicator
type BollingerUpper struct {

} // BollingerUpper

// BollingerLower Represents the Lower Bollinger Band Technical
// indicator
type BollingerLower struct {

} // BollingerLower

// RSI Represents the relative strength index technical 
// indicator
type RSI struct {

} // RSI

// Declaration of Dummy Method
func (SMA) 		Tag() {}
func (EMA) 		Tag() {}
func (Delta) 	Tag() {}
func (Diff) 	Tag() {}

// Type() Implementations
func (SMA)		Type() (string) {return "SMA"}
func (EMA) 		Type() (string) {return "EMA"}
func (Delta)	Type() (string) {return "delta"}
func (Diff) 	Type() (string) {return "diff"}
