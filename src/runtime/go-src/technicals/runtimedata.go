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
)

const (
	CapLimitMultiplier = 2
)

/*
GENERAL FLOW

1. PopLeft()
2. AppendNewOHLCV()
3. UpdateTALIBTechnicals()
4. UpdateOtherTechnicals()
5. UpdateRelationships()
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
// OHLCV arrays
func (rd *RuntimeData) AppendNewOHLCV(
	open float64,
	high float64,
	low float64,
	close float64,
	volume float64,
) {
	rd.OHLCV.Open = append(rd.OHLCV.High, open)
	rd.OHLCV.High = append(rd.OHLCV.High, high)
	rd.OHLCV.Low = append(rd.OHLCV.Low, low)
	rd.OHLCV.Close = append(rd.OHLCV.Close, close)
	rd.OHLCV.Volume = append(rd.OHLCV.Volume, volume)
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
func (rd *RuntimeData) UpdateTALIBTechnicals() error {
	for i := range rd.TALIBFeatureTechnicals {


		obj := &rd.TALIBFeatureTechnicals[i]

		// TODO: See issue #38
		res, err := talibDispatch[obj.Technical](Feature(*obj), &rd.OHLCV)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		obj.Value = res[0]
	}
	return nil
} // UpdateTALIBTechnicals

// UpdateOtherTechnicals updates the rd.OhterTechincals
// array after UpdateTALIBTechnicals() has been called
func (rd *RuntimeData) UpdateOtherTechnicals() error {
	for i := range rd.OtherFeatureTechnicals {

		obj := &rd.OtherFeatureTechnicals[i]

		if obj.Technical == "diff" {
			obj.Value = rd.TALIBFeatureTechnicals[rd.ColNames[obj.Col1]].Value - rd.TALIBFeatureTechnicals[rd.ColNames[obj.Col2]].Value
		} else { /// delta
			if obj.Col2 == "" { // single col
				obj.Value -= rd.TALIBFeatureTechnicals[rd.ColNames[obj.Col1]].Value
			} else { // delta of differences
				obj.Value -= rd.TALIBFeatureTechnicals[rd.ColNames[obj.Col1]].Value - rd.TALIBFeatureTechnicals[rd.ColNames[obj.Col2]].Value
			}
		}
	}

	return nil
} // UpdateOtherTechnicals

// UpdateOtherTechnicals updates the rd.Relationships
// array after UpdateOtherTechnicals() has been called
func (rd *RuntimeData) UpdateRelationships() error {
	return nil
} // UpdateOtherTechnicals
