package technicals

// This file contains method to call on the RuntimeData
// object
//
// Author: Vikas Katari
// Date: 07/18/2025

import "C"
import (
	"fmt"
	"unsafe"
	"strconv"
)

const (
	CapLimitMultiplier = 2
)

/*
GENERAL FLOW

1. PopLeft() Remove oldest value, reset cap if necessary
2. AppendNewOHLCV()  Add new bars to OHLCV arrays
3. UpdateTALIBTechnicals() Updates TALIB object values base for next method
4. UpdateOtherTechnicals() Updates Delta/Diff objects relies on prev call
5. UpdateRelationships() Updates relationship (labelling_logic) relies on prev 2 calls

INIT FLOW
1. Burn in OHLCV arrays
2. Initialize TALIB Technicals (UpdateTALIBTechnicals)
3. Initialize other user objects (UpdateOtherTechnicals)
4. Initialize technicals (InitRelationships) 
	Notice: this is not the same as UpdateRelationships 
	because we need to provide some values to the actual
	relationship value
*/

// CopySlice copies the Slice to reduce the capacity
// of the underlying array
func CopySlice(slice []float64) []float64 {
	copied := make([]float64, len(slice))
	copy(copied, slice)
	return copied
} // CopySlice

// CDoublePointerToSlice returns a C double pointer to a given index
// in a Go []float64 slice
func CDoublePointerToSlice(arr []float64, index int) *C.double {
	return (*C.double)(unsafe.Pointer(&arr[index]))
} // CDoublePointerToSlice

// InitializeAllPointers will initialize all pointers to point to every
// OHLCV array to a given index, this method assumes that all OHLCV arrays exist
func (rd *RuntimeData) SetAllPointersToIndex(index int) {
	rd.OHLCV.OpenPtr = CDoublePointerToSlice(rd.OHLCV.Open, index)
	rd.OHLCV.HighPtr = CDoublePointerToSlice(rd.OHLCV.High, index)
	rd.OHLCV.LowPtr = CDoublePointerToSlice(rd.OHLCV.Low, index)
	rd.OHLCV.ClosePtr = CDoublePointerToSlice(rd.OHLCV.Close, index)
	rd.OHLCV.VolumePtr = CDoublePointerToSlice(rd.OHLCV.Volume, index)
} // InitializeAllPointers

// NullifyAllPointers set all C pointers in the RuntimeData object
// to null, used during slice copying
func (rd *RuntimeData) NullifyAllPointers() {
	rd.OHLCV.OpenPtr = nil
	rd.OHLCV.HighPtr = nil
	rd.OHLCV.LowPtr = nil
	rd.OHLCV.ClosePtr = nil
	rd.OHLCV.VolumePtr = nil
} // NullifyAllPointers

// CopyArray resets the array back down to the original burn time * 2
// to prevent dangling C pointers
func CopyArray(arr []float64, newCap int) []float64 {
	new := make([]float64, len(arr), newCap)
	copy(new, arr)
	return new
} // CopyArray

// CopyArrays resets the capacity of all the OHLCV arrays in a RuntimeData
// object
func (rd *RuntimeData) CopyArrays() {
	rd.OHLCV.Open = CopyArray(rd.OHLCV.Open, rd.RuntimeSettings.BurnTime*CapLimitMultiplier)
	rd.OHLCV.High = CopyArray(rd.OHLCV.High, rd.RuntimeSettings.BurnTime*CapLimitMultiplier)
	rd.OHLCV.Low = CopyArray(rd.OHLCV.Low, rd.RuntimeSettings.BurnTime*CapLimitMultiplier)
	rd.OHLCV.Close = CopyArray(rd.OHLCV.Close, rd.RuntimeSettings.BurnTime*CapLimitMultiplier)
	rd.OHLCV.Volume = CopyArray(rd.OHLCV.Volume, rd.RuntimeSettings.BurnTime*CapLimitMultiplier)
	rd.OHLCV.sliceCapCount = 0
} // CopyArrays

// UpdateDeltas updates the delta values for every OHLCV array
func (rd *RuntimeData) UpdateOHLCVDeltas() {
	len := rd.RuntimeSettings.BurnTime
	rd.OHLCV.OpenDelta = rd.OHLCV.Open[len-1] - rd.OHLCV.Open[len-2]
	rd.OHLCV.HighDelta = rd.OHLCV.High[len-1] - rd.OHLCV.High[len-2]
	rd.OHLCV.LowDelta = rd.OHLCV.Low[len-1] - rd.OHLCV.Low[len-2]
	rd.OHLCV.CloseDelta = rd.OHLCV.Close[len-1] - rd.OHLCV.Close[len-2]
	rd.OHLCV.VolumeDelta = rd.OHLCV.Volume[len-1] - rd.OHLCV.Volume[len-2]
} // UpdateDeltas()

// PopLeft slices (pops) off the first value for all OHLCV arrays
func (rd *RuntimeData) PopLeft() {
	if rd.OHLCV.sliceCapCount > rd.OHLCV.SliceMaxCap {
		rd.NullifyAllPointers()
		rd.CopyArrays()
	}

	rd.SetAllPointersToIndex(1)

	rd.OHLCV.Open = rd.OHLCV.Open[1:]
	rd.OHLCV.High = rd.OHLCV.High[1:]
	rd.OHLCV.Low = rd.OHLCV.Low[1:]
	rd.OHLCV.Close = rd.OHLCV.Close[1:]
	rd.OHLCV.Volume = rd.OHLCV.Volume[1:]

	rd.OHLCV.sliceCapCount += 1
} // DropFirstArrayValues

// AppendNewOHLCV appends all new OHLCV values to all
// OHLCV arrays, this will Also add the first default
// OHLCV values to the feature payloads in YFinance order
// (CHLOV)
func (rd *RuntimeData) AppendNewOHLCV(
	bars [5]float64,
) {
	rd.OHLCV.Open = append(rd.OHLCV.Open, bars[0])
	rd.OHLCV.High = append(rd.OHLCV.High, bars[1])
	rd.OHLCV.Low = append(rd.OHLCV.Low, bars[2])
	rd.OHLCV.Close = append(rd.OHLCV.Close, bars[3])
	rd.OHLCV.Volume = append(rd.OHLCV.Volume, bars[4])

	// 3 1 2 0 4
	// YFinance DF == CHLOV
	rd.FeatureArray[0] = bars[3] // Close MUST GO first to match Training data
	rd.FeatureArray[1] = bars[1]
	rd.FeatureArray[2] = bars[2]
	rd.FeatureArray[3] = bars[0] // Open goes here
	rd.FeatureArray[4] = bars[4]

	rd.FeatureJSON["0"] = bars[3]
	rd.FeatureJSON["1"] = bars[1]
	rd.FeatureJSON["2"] = bars[2]
	rd.FeatureJSON["3"] = bars[0]
	rd.FeatureJSON["4"] = bars[4]

	rd.fillFeatureIndex = 5
} // AppendNewOHLCV

// TestAppend is a testing method that is used to test
// OHLCV array lifetime management
// NOT FOR PRODUCTION CODE
func (rd *RuntimeData) TestAppend(val float64) {
	rd.OHLCV.Open = append(rd.OHLCV.Open, val)
	rd.OHLCV.High = append(rd.OHLCV.High, val)
	rd.OHLCV.Low = append(rd.OHLCV.Low, val)
	rd.OHLCV.Close = append(rd.OHLCV.Close, val)
	rd.OHLCV.Volume = append(rd.OHLCV.Volume, val)
} // TestAppend

// UpdateTALIBTechnicals updates the rd.TALIBTechnicals
// object after the rd.OHLCV arrays have been updated
// i.e. Burn in has been complete
func (rd *RuntimeData) UpdateTALIBTechnicals() error {
	for i := range rd.TALIBFeatureTechnicals {

		obj := &rd.TALIBFeatureTechnicals[i]

		// TODO: See issue #38
		res, err := talibDispatch[obj.Technical](Feature(*obj), &rd.OHLCV)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		obj.Value = res[0] // related issue with #41
		
		rd.FeatureArray[rd.fillFeatureIndex] = res[0]
		rd.FeatureJSON[strconv.Itoa(rd.fillFeatureIndex)] = res[0]
		rd.fillFeatureIndex++
	}

	return nil
} // UpdateTALIBTechnicals

// getOtherFeatureMappedValues returns the values
// that a non TALIB object maps to in the TALIB array 
// returns as [col1Val, col2Val]
func getOtherFeatureMappedValues(rd *RuntimeData, ft *FeatureTechnical) []float64 {
	col1Val := rd.TALIBFeatureTechnicals[rd.ColNames[ft.Col1]].Value
	col2Val := rd.TALIBFeatureTechnicals[rd.ColNames[ft.Col2]].Value
	return []float64{col1Val, col2Val}
} // getOtherTechnicalValue

// UpdateOtherTechnicals updates the rd.OtherTechnicals
// array after UpdateTALIBTechnicals() has been called
func (rd *RuntimeData) UpdateOtherTechnicals() error {
	for i := range rd.OtherFeatureTechnicals {

		obj := &rd.OtherFeatureTechnicals[i]

		matchingValues := getOtherFeatureMappedValues(
			rd,
			obj,
		)

		if obj.Technical == "diff" {
			obj.Value = matchingValues[0] - matchingValues[1]
		} else { /// delta
			if obj.Col2 == "" { // single col
				obj.Value -= matchingValues[0]
			} else { // delta of differences
				obj.Value -= matchingValues[0] - matchingValues[1]
			}
		}

		rd.FeatureArray[rd.fillFeatureIndex] = obj.Value
		rd.FeatureJSON[strconv.Itoa(rd.fillFeatureIndex)] = obj.Value
		rd.fillFeatureIndex++
	}
	return nil
} // UpdateOtherTechnicals

// UpdateRelationships updates the rd.Relationships
// array after UpdateOtherTechnicals() has been called
func (rd *RuntimeData) UpdateRelationships() error {
	for i := range rd.Relationships {

		obj := &rd.Relationships[i]
		
		compValues, err := obj.getRelationshipValues(rd)
		if err != nil {
			return fmt.Errorf("technicals.UpdateRelationships(): Could not process relationship object at index %d", i)
		}

		featureVal := relationshipDispatch[obj.Signal](
			obj,
			rd,
			compValues,
		)

		weightedFeatureVal := obj.Weight * featureVal

	 	if weightedFeatureVal < obj.Threshold {
			weightedFeatureVal = 0
		}

		rd.FeatureArray[rd.fillFeatureIndex] = weightedFeatureVal
		rd.FeatureJSON[strconv.Itoa(rd.fillFeatureIndex)] = weightedFeatureVal
		rd.fillFeatureIndex++
	}

	rd.fillFeatureIndex = 0

	return nil
} // UpdateOtherTechnicals

// InitRelationships initializes the relationship values
// once the TALIB and Other features have been initialized as 
// well, SHOULD ONLY EVER BE CALLED ONCE 
func (rd *RuntimeData) InitRelationships() error {

	for i := range rd.Relationships {

		obj := &rd.Relationships[i]

		col1Val, err := getColVal(
			rd, 
			rd.Relationships[i].Col1, 
			getPrefix(rd.Relationships[i].Col1),
		)
		if err != nil {
			return fmt.Errorf("col1Val failed to return for relationship index %d", i)
		}

		col2Val, err := getColVal(
			rd, 
			rd.Relationships[i].Col2, 
			getPrefix(rd.Relationships[i].Col2),
		)
		if err != nil {
			return fmt.Errorf("col1Val failed to return for relationship index %d", i)
		}

		obj.Col1Val = col1Val
		obj.Col2Val = col2Val
	}			

	return nil
} // InitRelationships	