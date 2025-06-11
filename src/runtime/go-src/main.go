package main

// Execution file for the runtime environemnt

import (
	eventLoop "github.com/vkatari10/trading-bot/src/runtime/go-src/eventloop"
	// "fmt"
	// api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	//engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
)

func main() {
	eventLoop.Run() 
} // main
