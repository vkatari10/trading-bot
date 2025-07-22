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
	// models JSON objects
	ColNames map[string]int

	// All features
	TALIBFeatureTechnicals []FeatureTechnical // technical indicators
	OtherFeatureTechnicals []FeatureTechnical // delta/diff
	Relationships []Relationship // labelling logic

	// contains all historical data
	OHLCV TALIBWrapper 

	// for FastAPI served models
	FeatureJSON map[string]float64

	// for ONNX served models
	FeatureArray []float64
	
	// live trade tickers
	Tickers []string

	RuntimeSettings RuntimeSettings
} // RuntimeData

// RuntimeSettings is a struct to represent user runtime
// settings from the "runtime_settings" section of the JSON
type RuntimeSettings struct {
	CycleTime 		int `json:"cycle_time"` // seconds 
	BurnTime 		int `json:"burn_window_time"` // minutes
	LogAPIFlushTime int `json:"log_api_flush_time"` // milliseconds
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

	// historical data
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
	Technical 	string 				`json:"tech"`
	Name 		string 				`json:"name"`
	Col1 		string				`json:"col1"`
	Col2 		string				`json:"col2"`
	Value 		float64		
	Args 		map[string]float64 	`json:"args"`
}

// Relationship represents the "labelling logic" section 
// of the config JSON for each object
type Relationship struct {
	Name 	string 	`json:"name"`
	Col1 	string 	`json:"col1"`
	Col2 	string 	`json:"col2"`
	Signal 	string 	`json:"sig"`
	Weight 	float64 `json:"weight"`
	Persist int 	`json:"persist"`
}

// declarations
func (FeatureTechnical) Tag() {}
func (Relationship) 	Tag() {}

func (FeatureTechnical) Type() (string) {return "FEATURE_TECHNICAL"}
func (Relationship) 	Type() (string) {return "RELATIONSHIP"}
