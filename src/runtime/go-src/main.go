package main

import "fmt"
import "runtime" // for GC

import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

// Main Runtime Engine should be placed here
func main() {

     runtime.GC() // force GC before starting main loop

     doubles := []float64{3.0, 5.0, 7.0}

     mean := engine.Mean(doubles)

     stdDev := engine.StdDev(doubles)

     fmt.Println(mean)
     fmt.Println(stdDev)

} // main