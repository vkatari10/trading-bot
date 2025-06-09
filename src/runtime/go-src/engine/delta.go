package engine


import (
	"fmt"
)

// NewDelta makes a new Delta reference
func NewDelta(json map[string]any) (*Delta, error) {
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
			fmt.Println(v.Data)
		
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