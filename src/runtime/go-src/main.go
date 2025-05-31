package main

import "fmt"
import "runtime" // for GC

import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

// Main Runtime Engine should be placed here
func main() {

     runtime.GC()

     engine.TestC()
     fmt.Println("Done")

     fmt.Println(engine.ParseLogicJSON("features.json"))

} // main
