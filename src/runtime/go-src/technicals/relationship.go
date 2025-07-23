package technicals

import (
	"fmt"
	"strings"
	"math"
)

type relationshipUpdater func(r *Relationship, rd *RuntimeData, compValues []float64) (float64)

var (
	relationshipDispatch map[string]relationshipUpdater // relationship method dispatch table
)

// init loads the relationship dispatch table
func init() {
	relationshipDispatch = map[string]relationshipUpdater{
		"crossover": CrossoverRelationship,
		"above": AboveRelationship,
		"below": BelowRelationship,
	}
} // init

// getColVal gets the value given a column name and its prefix value
func getColVal(rd *RuntimeData, col string, arr rune) (float64, error) {
	if arr == 'T' {
		return rd.TALIBFeatureTechnicals[rd.ColNames[col]].Value, nil
	} else if arr == 'D' {
		return rd.OtherFeatureTechnicals[rd.ColNames[col]].Value, nil
	}
	return 0, nil
} // getColVal

// getPrefix gets the prefix value of a given relationship
// and its ColName
func getPrefix(colName string) rune {
	if strings.HasPrefix(colName, "T_") { // TALIB
		return 'T'
	} else if strings.HasPrefix(colName, "D_") { // Other
		return 'D'
	}
	return '_'
} // getPrefix

// getRelationshipValues gets the relationship values for both col1
// and the col2 fields in the struct and returns the newest values given
// the values that they map to.
// returns values as [newCol1Val, newCol2Val]
func (r *Relationship) getRelationshipValues(rd *RuntimeData) ([]float64, error) {

	newCol1, err := getColVal(rd, r.Col1, getPrefix(r.Col1))
	if err != nil {
		return nil, 
		fmt.Errorf("getRelationshipValues failed for %s col1", r.Name)
	} 

	newCol2, err := getColVal(rd, r.Col2, getPrefix(r.Col2))
	if err != nil {
		return nil,
		fmt.Errorf("getRelationshipValues failed for %s col2", r.Name)
	}

	return []float64{newCol1, newCol2}, nil
} // getRelationshipValues

// setNewColValues sets new col values inside a relationship once a signal has been 
// detected
func (r *Relationship) setNewColValues(newCol1Val float64, newCol2Val float64) {
	r.Col1Val = newCol1Val
	r.Col2Val = newCol2Val
} // setNewValue

// CrossoverRelationship checks if the new compValues have changed
// to where the new compValues indicate a cross up where col1 > oldCol2 && col2 < oldCol1
// returning 1, else -1 if the opposite is true, else 0
func CrossoverRelationship(r *Relationship, rd *RuntimeData, compValues []float64) (float64) {
	e := 1e-9 // epsilon

	oldDelta := r.Col1Val - r.Col2Val
	newDelta := compValues[0] - compValues[1]

	var ans float64

	if oldDelta < -e && newDelta > e { 
		ans = 1.0
	} else if oldDelta > -e && newDelta < e {
		ans = -1	
	} else {
		ans =  0
	}

	r.setNewColValues(compValues[0], compValues[1])
	return ans
} // CrossoverRelationship

// AboveRelationship checks if the new compValues have changed
// to where col1 > col2; if true returns 1, else if the values
// are equivalent or the persistence counter is below required
// 0, else -1
func AboveRelationship(r *Relationship, rd *RuntimeData, compValues []float64) (float64) {
	e := 1e-9 

	var ans float64

	diff := compValues[0] - compValues[1]

	if math.Abs(diff) < e {
		ans = 0
		r.PersistCounter = 0
	} else if diff > 0 {

		r.PersistCounter++

		// persist tracker
		if r.PersistCounter >=r.Persist {
			ans = 1
			r.PersistCounter = 0
		} else {
			ans = 0
		}

	} else {
		ans = -1
		r.PersistCounter = 0
	}

	r.setNewColValues(compValues[0], compValues[1])
	return ans
} // AboveRelationship

// BelowRelationship checks if the new compValues have changed
// to where col1 < col2; if true returns 1, else if the values
// are equivalent 0, else -1
func BelowRelationship(r *Relationship, rd *RuntimeData, compValues []float64) (float64) {
		e := 1e-9 

	var ans float64

	diff := compValues[0] - compValues[1]

	if math.Abs(diff) < e {
		ans = 0
		r.PersistCounter = 0
	} else if diff < 0 {

		r.PersistCounter++

		// persist tracker
		if r.PersistCounter >= r.Persist {
			ans = 1
			r.PersistCounter = 0
		} else {
			ans = 0
		}

	} else {
		ans = -1
		r.PersistCounter = 0
	}

	r.setNewColValues(compValues[0], compValues[1])
	return ans
} // BelowRelationship