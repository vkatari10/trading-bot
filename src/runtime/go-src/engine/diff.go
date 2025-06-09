package engine

// Contains Logic For Diff(erence) Objects

import (
	"fmt"
	"log"
)

// NewDiff makes a new Diff reference
func NewDiff(json map[string]any) (*Diff, error) {
	col1, ok := json["col1"].(string)
	if !ok {
		return nil, fmt.Errorf("col1 should be a string")
	} // if

	col2, ok := json["col2"].(string)
	if !ok {
		return nil, fmt.Errorf("col2 should be a string")
	} // if

	return &Diff {
		Col1: col1,
		Col2: col2,
		Value: 0,
	}, nil
} // NewDelta


// Load method for Diff objects should be called 
// after burn-in time
func (diff *Diff) Load(data *UserData) (error) {
	// Match the col index to the object array index
	diff.Col1Index = data.ColNames[diff.Col1]

	if diff.Col2 == "" {
		return fmt.Errorf("col2 cannot be empty")
	}
	diff.Col2Index = data.ColNames[diff.Col2]

	return nil
} // Load (Diff)

// GetNew (Diff) Gets the new Diff Value based on the specified
// col names (NOT DESIGNED TO UPDATE THE OCHLV VALUES)
func (diff *Diff) GetNew(data *UserData) {
	if diff.Col1Index < 0 || diff.Col2Index < 0 {
		log.Fatal("One, or both Diff col names are not valid\n")
		return
	} // if

	cols := [2]Indicator{data.Objects[diff.Col1Index], data.Objects[diff.Col2Index]}
	results := [2]float64{}

	for i, ind := range cols {
		switch v := ind.(type) {
		case *SMA:
			results[i] = v.Data[len(v.Data) - 1]
		case *EMA:
			results[i] = v.Data[len(v.Data) - 1]
		} // swtich
	} // for

	diff.Value = results[0] - results[1]
} // GetNew (Diff)





