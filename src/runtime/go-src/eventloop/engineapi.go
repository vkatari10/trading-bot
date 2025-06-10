package eventloop

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
)

var (
	logLink = "http://locahost:3000/api/log" 			// send logging information
	quoteLink = "http://localhost:3000/api/quote" 		// send quote update information
	envLink = "http://localhost:3000/api/env"			// send environment variables (done once)
	brokerLink = "http://localhost:3000/api/broker" 	// send brokerage account infromation (done on sell/buy)
	technicalLink = "http://localhost:3000/api/technical" // send technical update information
)

// SendPayload should send the JSON as an Object to the frontend
func SendPayload(data map[string]any, postLink string) {
	json, err := json.Marshal(data)
    if err != nil {
        log.Println(err)
		return
    } // if

    resp, err := http.Post(postLink, "application/json", bytes.NewBuffer(json))
    if err != nil {
        log.Println(err)
		return
    } // if

    defer resp.Body.Close()
} // SendPayload