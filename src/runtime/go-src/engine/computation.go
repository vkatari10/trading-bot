package engine

// this file contains methods to compute technicals in real time

/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data -lm 

#include "live_technicals.h"
#include "live_stats.h"
*/
import "C"

/*
TODO implement functions
*/

import (
    "unsafe"
)

// Mean test function to find the mean of an array calling C code
func Mean(array []float64) float64 {
    ptr := (*C.double)(unsafe.Pointer(&array[0]))

    mean := C.mean(ptr, C.size_t(len(array)))

    return float64(mean)
} // Mean

// StdDev test function to find the standard deviation of an array 
// calling C code
func StdDev(array []float64) float64 {
     ptr := (*C.double)(unsafe.Pointer(&array[0]))

     stdDev := C.std_dev(ptr, C.size_t(len(array)))

     return float64(stdDev)
} // StdDev

/*
Implement C sma, ema functions

*/

// Returns an array of SMA values based on the given window
// func SMA(array []float64, window uint32) []float64 {

//     ptr := (*C.double)(unsafe.Pointer(&array[0]))
// } // SMA