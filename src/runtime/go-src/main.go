package main

import "fmt"
import "runtime" // for GC
import "sync"

import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

// Main Runtime Engine should be placed here
func main() {

	runtime.GC() // force GC before starting main loop

	thing, err := engine.ParseLogicJSON("features.json")

	if err != nil {
		fmt.Println("Idk")
	} // if

	for i := range thing {
		go fmt.Printf("tech: %s, window %f\n", thing[i]["name"], thing[i]["window"])
	} // for

} // main
