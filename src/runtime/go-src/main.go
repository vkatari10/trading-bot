package main

// Execution file for the runtime environemnt

import (
	//eventLoop "github.com/vkatari10/trading-bot/src/runtime/go-src/eventloop" // eventloop
	"fmt"

	//api "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
	engine "github.com/vkatari10/trading-bot/src/runtime/go-src/engine"
	
)

//Main Runtime Engine
func main() {
	//eventLoop.Run()
	


	userData, err := engine.InitUserLogic("features.json")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(userData)

	for i := range userData.Objects {
		fmt.Println(userData.Objects[i].Type())
	}

	

} // main