package main

/*
#cgo CFLAGS: -I../../live_data/csrc/
#cgo LDFLAGS: -L../../live_data/csrc -llive_data

#include "live_technicals.h"

*/
import "C"

import "runtime"
import "fmt"

func main() {
    
     /*
     this will be important later 
     as this file serves as the runtime
     environment file in the future so 
     we want to GC before starting
     */
      // forces GC
	runtime.GC()
     
     fmt.Println("Before C call")
     C.dummy_test()
     fmt.Println("After C call")
} // main