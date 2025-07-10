package technicals

// Contains logic for the Delta Object 

import (
	"fmt"
)

// NewDelta makes a new Delta reference
func NewDelta(json map[string]any) (Indicator, error) {
	col1, ok := json["col1"].(string)
	if !ok {
		return nil, fmt.Errorf("col1 should be a string")
	} // if

	col2, ok := json["col2"].(string)
	if !ok {
		 // col2 is optional for deltas if we dont want diff of deltas
		col2 = ""
	} 	

	return &Delta {
		Col1: col1,
		Col2: col2,
		Value: 0,
	}, nil
} // NewDelta

// GetNew (Delta) Gets the new Delta Value based on the specificed
// col names (NOT DESIGNED TO UPDATE THE OCHLV VALUES)
// should be called after updating underlying data
func (delta *Delta) GetNew(data *UserData) {	
	results := [2]float64{}
	var cols [2]Indicator;

	if delta.Col2Index == -1 { // one col
		cols = [2]Indicator{data.Objects[delta.Col1Index], nil}
	} else {
		cols = [2]Indicator{data.Objects[delta.Col1Index], data.Objects[delta.Col2Index]}
	} // if-else

	for i, ind := range cols {
		if cols[i] == nil {
			break
		} // if
		switch v := ind.(type) {
		case *SMA:
			results[i] = v.Data[len(v.Data) - 1] - v.Data[len(v.Data) - 2]
		case *EMA:
			results[i] = v.Data[len(v.Data) - 1] - v.Data[len(v.Data) - 2]
		} // swtich
	} // for

	if delta.Col2Index == -1 { 	// one col delta
		delta.Value = results[0]
	} else { 					// delta of differences
		delta.Value = results[0] - results[1]
	} // if-else

} // GetNew (Delta)

// Load method for Delta objects, matches their column
// indexes to the UserData.Objects slice
func (delta *Delta) Load(data *UserData) (error) {
	// Match the col index to the object array index
	delta.Col1Index = data.ColNames[delta.Col1]

	if delta.Col2 != "" { // if col2 isn't null in JSON
		delta.Col2Index = data.ColNames[delta.Col2]
	} else {  // if col2 was null in JSON
		delta.Col2Index = -1
	} // if-else

	return nil
} // Load (Delta)

// LastValue Gets the latest value of the delta
func (delta *Delta) LastValue() []float64 {
	return []float64{delta.Value}
} // LastValue (Delta)