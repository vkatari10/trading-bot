package technicals

// This file contains the TA-Lib wrapper FFI to call
// the compiled C library code 

// MODIFY FLAGS BELOW AND PATHS IF YOU ENCOUNTER 
// LINKER ERRORS BELOW

// DO NOT MODIFY THE INCLUDE STATEMENTS

// DO NOT REMOVE THE -I or -L BEFORE THE PATHS

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





