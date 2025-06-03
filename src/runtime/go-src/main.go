package main

import "fmt"
import "runtime" // for GC
//import "sync"

import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

// Main Runtime Engine should be placed here
func main() {

	runtime.GC() // force GC before starting main loop

	inds, err := engine.InitUserLogic("features.json")
	if err != nil {
		fmt.Errorf(err.Error())
	} // if

	for i := range inds.Ind {
		print(inds.Ind[i].Data)
	} // for
	


} // main
