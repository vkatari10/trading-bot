package main

import "fmt"
import "runtime" // for GC

import engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"

func main() {

     runtime.GC()

     engine.TestC()
     fmt.Println("Done")

} // main
