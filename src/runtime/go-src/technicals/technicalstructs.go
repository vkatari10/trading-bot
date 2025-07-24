package technicals

// Struct declarations to model user features and labelling logic

import "C"

// Represents a technical indicator 
type Technical interface {
	Tag()
	Type() string
} // Technical

// Represents technicals and price data 
type RuntimeData struct {
	// models user given name: index on TALIBFeatureTechnicals
	ColNames map[string]int

	// All features -> update these in order as listed
	
	TALIBFeatureTechnicals []FeatureTechnical // technical indicators
	OtherFeatureTechnicals []FeatureTechnical // delta/diff
	Relationships []Relationship // labelling logic

	// contains all historical data
	OHLCV TALIBWrapper 

	// for FastAPI served models
	FeatureJSON map[string]float64

	// for ONNX served models
	FeatureArray []float64

	fillFeatureIndex int
	
	// live trade tickers
	Tickers []string

	RuntimeSettings RuntimeSettings
} // RuntimeData

// RuntimeSettings is a struct to represent user runtime
// settings from the "runtime_settings" section of the JSON
type RuntimeSettings struct {
	CycleTime 		float64 `json:"cycle_time"` // seconds 
	BurnTime 		int `json:"burn_window_time"` // minutes
	LogAPIFlushTime float64 `json:"log_api_flush_time"` // milliseconds
	LogToStdout 	bool `json:"log_to_stdout"`
	RunAfterClose 	bool `json:"run_after_close"`
	OverrideBurnIn 	bool `json:"override_burn_in"`
} // RuntimeSettings

// TALIBWrapper represents a TA-Lib wrapper to interact
// with the compiled C code using cgo 
type TALIBWrapper struct {
	// slice capacity tracker	
	sliceCapCount int
	SliceMaxCap int

	// historical OHLCV
	Open 	[]float64
	High 	[]float64 
	Low 	[]float64
	Close 	[]float64
	Volume 	[]float64

	// pointers
	OpenPtr *C.double
	HighPtr *C.double
	LowPtr	*C.double
	ClosePtr *C.double
	VolumePtr *C.double

	// deltas
	OpenDelta 	float64
	HighDelta 	float64	
	LowDelta 	float64
	CloseDelta 	float64
	VolumeDelta float64
} // TALIBWrapper

// Represents all features on the object array 
type Feature interface{
	Tag()
	Type() string
} 

// Technical represents the "feature" section 
// of the config JSON for each object
type FeatureTechnical struct {
	Technical 	string 				`json:"tech"` // must match to talibDispatch
	Name 		string 				`json:"name"` // user input name
	Col1 		string				`json:"col1"` // for delta/diff objects
	Col2 		string				`json:"col2"`
	Value 		float64 // actual feature value
	Args 		map[string]float64 	`json:"args"` // TALIB kwargs
} // FeatureTechnical

// Relationship represents the "labelling logic" section 
// of the config JSON for each object
type Relationship struct {
	Name 	string 	`json:"name"`
	Col1 	string 	`json:"col1"`
	Col2 	string 	`json:"col2"`
	Signal 	string 	`json:"sig"`
	Weight 	float64 `json:"weight"`
	Persist int 	`json:"persist"`
	Threshold float64 `json:"threshold"`

	// track historical prices 
	Col1Val float64
	Col2Val float64

	// if we need to track above/below thresholds
	PersistCounter int 

	/*

	Say Col1Val and Col2Val are not default values
	and we get back new data for it 

	we can just say compare these old values to the new ones
	at that point and then for 3 cases

	1. crossover -- cross check col1, col2 vals against 
	existing values

	2. above/below -- check colX against ColY and update
	persist value if needed and replace existing values
	as well and if they fail to satisfy it again then reset
	the persist counter back down to 0

	Based on those values we can multiply the resulting value
	based on whatever values specified in the docs and multiply by 
	weight

	*/
} // Relationship

// declarations
func (FeatureTechnical) Tag() {}
func (Relationship) 	Tag() {}

func (FeatureTechnical) Type() (string) {return "FEATURE_TECHNICAL"}
func (Relationship) 	Type() (string) {return "RELATIONSHIP"}
