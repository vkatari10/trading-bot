package technicals

// This file conatains methods that all Updater (GetNew())
// methods for all technical indicators may need

// #cgo CFLAGS: -I/usr/include/ta-lib/
// #cgo LDFLAGS: -L/usr/lib -lta-lib -lm
// #include <ta_libc.h>
import "C"


// CopySlice copies the Slice to reduce the capacity 
// of the underlying array
func CopySlice(slice []float64) []float64 {

	copied := make([]float64, len(slice))
	copy(copied, slice)
	return copied

} // CopySlice



