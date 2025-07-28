package eventloop

// Methods to send Payloads for logging/data information
// 
// Author: Vikas Katari
// Date: 07/28/2025

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
)

// for post links like /api/data/{id}
// those id's should be embedded into the postLink variables
// when calling these methods, i.e. these methods are not 
// responsible for assigning IDs
// All logging information should be handled with the 
// APIBuffer and enqueue and dequeue methods 

// sendBrokeragePayload sends brokerage data to a given 
// endpoint as a JSON
func sendBrokeragePayload(payload map[string]any, postLink string) {

	// data should be wrapped as an array of tuples?

} // sendBrokeragePayload

// sendBrokeragePayload sends technical data to a given 
// endpoint as a JSON with a given id
func sendDataPayload(payload map[string]float64, postLink string) {
	copyPayload := make(map[string]float64, len(payload))
	for k, v := range payload {
		copyPayload[k] = v
	}

	go func(payload map[string]float64, postLink string){
		JSONpayload, err := json.Marshal(payload)
		if err != nil {
			log.Println("sendDataPayload: Could not marshal data")
			return
		}

		resp, err := http.Post(postLink, "application/json", bytes.NewBuffer(JSONpayload))
		if err != nil {
			log.Printf("sendDataPayload error POST to %s failed -> %v\n", postLink, err)
			return
		}
		defer resp.Body.Close()
	}(copyPayload, postLink)
} // sendDataPayload


/*
1. On new quote payload?

Send New Quote:
{
		"ticker": <ticker>,
		"last_quote": <quote>,
		"change": calculate...,
		"action": <action>
}

On Buy/Sell

Send New Account info:
{
		"ticker": <ticker>,
		"qty": <qty>,
		"avg_cost": <avg_cost>,
		"P/L": <avg_cost>,
		"Net Liq": calculate?,
}
*/