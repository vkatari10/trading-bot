package technicals

// Represents a technical indicator 
type Technical interface {
	Tag()
	Type() string
	LastValue() []float64 // Returns the latest value
} // Technical

type RuntimeData struct {
	ColNames map[string]int
	Objects []Technical
	OHLCV PriceData
} // RuntimeData

type PriceData struct {
	MaxLookBack int
	Open 		[]float64
	High 		[]float64
	Low 		[]float64
	Close 		[]float64
	Volume 		[]float64
	OpenDelta 	float64
	HighDelta 	float64
	LowDelta 	float64
	CloseDelta 	float64
	VolumeDelta float64
} // PriceData

