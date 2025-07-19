package technicals

// This file contains the TA-Lib wrapper FFI to call
// the compiled C library code 

// MODIFY FLAGS BELOW AND PATHS IF YOU ENCOUNTER 
// LINKER ERRORS BELOW

// DO NOT MODIFY THE INCLUDE STATEMENTS

// DO NOT MODIFY THE -lta-lib, -lm ARGUMENTS

// #cgo CFLAGS: -I/usr/include/ta-lib/
// #cgo LDFLAGS: -L/usr/lib -lta-lib -lm
// #include <ta_libc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
	"fmt"
)

// CopySlice copies the Slice to reduce the capacity 
// of the underlying array
func CopySlice(slice []float64) []float64 {
	copied := make([]float64, len(slice))
	copy(copied, slice)
	return copied
} // CopySlice

// CDoublePointerToSlice returns a C double pointer to a given index
// in a Go []float64 slice
func CDoublePointerToSlice(arr []float64, index int) (*C.double) {
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
	rd.OHLCV.Open = CopyArray(rd.OHLCV.Open, rd.RuntimeSettings.BurnTime * 2)
	rd.OHLCV.High = CopyArray(rd.OHLCV.High, rd.RuntimeSettings.BurnTime * 2)
	rd.OHLCV.Low = CopyArray(rd.OHLCV.Low, rd.RuntimeSettings.BurnTime * 2)
	rd.OHLCV.Close = CopyArray(rd.OHLCV.Close, rd.RuntimeSettings.BurnTime * 2)
	rd.OHLCV.Volume = CopyArray(rd.OHLCV.Volume, rd.RuntimeSettings.BurnTime * 2)
	rd.OHLCV.sliceCapCount = 0
} // CopyArrays

// UpdateDeltas updates the delta values for every OHLCV array 
func (rd *RuntimeData) UpdateDeltas() {
	len := rd.RuntimeSettings.BurnTime
	rd.OHLCV.OpenDelta = rd.OHLCV.Open[len - 1] - rd.OHLCV.Open[len - 2]
	rd.OHLCV.HighDelta = rd.OHLCV.High[len - 1] - rd.OHLCV.High[len - 2]
	rd.OHLCV.LowDelta = rd.OHLCV.Low[len - 1] - rd.OHLCV.Low[len - 2]
	rd.OHLCV.CloseDelta = rd.OHLCV.Close[len - 1] - rd.OHLCV.Close[len - 2]
	rd.OHLCV.VolumeDelta = rd.OHLCV.Volume[len - 1] - rd.OHLCV.Volume[len - 2]
} // UpdateDeltas()

// DropFirstArrayValues slices off the first value for all OHLCV arrays 
func (rd *RuntimeData) DropFirstArrayValues() {
	if rd.OHLCV.sliceCapCount > rd.OHLCV.SliceMaxCap {
		rd.NullifyAllPointers()
		rd.CopyArrays()
	}

	rd.SetAllPointersToIndex(1)
	rd.cleanUpArrays()

	rd.OHLCV.Open = rd.OHLCV.Open[1:]
	rd.OHLCV.High = rd.OHLCV.High[1:]
	rd.OHLCV.Low = rd.OHLCV.Low[1:]
	rd.OHLCV.Close = rd.OHLCV.Close[1:]
	rd.OHLCV.Volume = rd.OHLCV.Volume[1:]

	rd.OHLCV.sliceCapCount += 1
} // DropFirstArrayValues

// cleanUpArrays just resets the most recent dropped value
// in a slice to its default to flag GC its unused
func (rd *RuntimeData) cleanUpArrays() {
	rd.OHLCV.Open[0] = 0.0
	rd.OHLCV.High[0] = 0.0
	rd.OHLCV.Low[0] = 0.0
	rd.OHLCV.Close[0] = 0.0
	rd.OHLCV.Volume[0] = 0.0
} // cleanUpArrays()

func (rd *RuntimeData) InitializeTechnicals() {

}

func (rd *RuntimeData) UpdateTechnicals() {

}

func (rd *RuntimeData) computeTechnicalValues() {

	// for every object in objects

	// run the dispatcher for TA-Lib

	// put value to said index in objects

	// every call on its own go routine

	// for i := range rd.Objects {
	// 	// based on type of object call a 
	// 	// dispatch table for TA lib

	// 	// each object should contain a value or something?
		
	// }

}


func TASMA() {
    // Initialize TA-Lib
    if err := C.TA_Initialize(); err != C.TA_SUCCESS {
        panic("TA_Initialize failed") 
    }
    defer C.TA_Shutdown()

    // Input data (Go array)
    in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    period := C.int(3)

    // Allocate memory for input and output
    input := (*C.double)(C.malloc(C.size_t(len(in)) * C.size_t(unsafe.Sizeof(C.double(0)))))
    defer C.free(unsafe.Pointer(input))

    for i := 0; i < len(in); i++ {
        *(*C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(input)) + uintptr(i)*unsafe.Sizeof(C.double(0)))) = C.double(in[i])
    }

    out := (*C.double)(C.malloc(C.size_t(len(in)) * C.size_t(unsafe.Sizeof(C.double(0)))))
    defer C.free(unsafe.Pointer(out))

    var outBeg, outNb C.int

    // Call TA_SMA
    retCode := C.TA_SMA(0, C.int(len(in)-1), input, period, &outBeg, &outNb, out)
    if retCode != C.TA_SUCCESS {
        panic(fmt.Sprintf("TA_SMA failed with code %d", retCode))
    }

    // Print output
    for i := 0; i < int(outNb); i++ {
        val := *(*C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(i)*unsafe.Sizeof(C.double(0))))
        fmt.Printf("SMA[%d] = %.2f\n", i+int(outBeg), val)
    }
}





