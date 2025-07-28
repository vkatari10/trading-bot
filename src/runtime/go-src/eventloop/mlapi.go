package eventloop

// This file interact with the ML model in Python to send data and get 
// predictions back to inform the broker API
//
// Author: Vikas Katari
// Date: 05/30/2025

import (
    "log"
    "github.com/gorilla/websocket"
)

// websocketWriter will write the features to the ML API Server as a JSON
// given the channel with the payload to send
func websocketWriter(conn *websocket.Conn, payload <-chan map[string]float64) {
    for {
        err := conn.WriteJSON(<-payload)
        if err != nil {
            log.Println("ERROR: Could not write JSON to ML API")
        } // if
    } // for
} // websocketWriter

// websocketReader returns the inference to the result channel 
// based on the written features
func websocketReader(conn *websocket.Conn, result chan<- float64) {
    for {
        var response map[string]any
        err := conn.ReadJSON(&response)
        if err != nil {
            log.Println("ERROR: Could not get result from ML API", err)
        } // if

        serverResult, ok := response["result"].(float64)
        if !ok {    
            log.Println("ERROR: Could not read response")
        } // if

        result <- serverResult
    } // for
} // websocketReader