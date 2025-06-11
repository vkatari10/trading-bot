package eventloop

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
)

var (
	logLink = "http://localhost:3000/api/log" 			// send logging information
	envLink = "http://localhost:3000/api/env"			// send environment variables (done once)
	brokerLink = "http://localhost:3000/api/broker" 	// send brokerage account infromation (done on sell/buy)
	dataLink = "http://localhost:3000/api/data" 		// send technical update information
)

// SendPayload should send the JSON as an Object to the frontend
func SendPayload(data map[string]any, postLink string) {
	payload, err := json.Marshal(data)
    if err != nil {
        log.Println(err)
		return
    } // if

    resp, err := http.Post(postLink, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        log.Println(err)
		return
    } // if

    defer resp.Body.Close()
} // SendPayload