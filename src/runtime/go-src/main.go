package main

/*
#cgo CFLAGS: -I../c-src/include
#cgo LDFLAGS: -L../c-src -llive_data

#include "live_technicals.h"

*/
import "C"

import "fmt"
import "runtime"

func main() {
     runtime.GC() // force garbage collector

     fmt.Println("Before C function Call")
     C.dummy_test()
     fmt.Println("After C function Call")
} // main
