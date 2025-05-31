package engine

// this file contains methods to compute technicals in real time

/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data

#include "live_technicals.h"
*/
import "C"

/*
TODO implement functions
*/

import (
    "fmt"
)

func TestC() {

     fmt.Println("Before C function call from runtime/technicals.go")
     C.dummy_test()
     fmt.Println("After C function call from runtime/techincals.go")

} // TestC
