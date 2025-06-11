package eventloop

// APIBuffer is just a Queue implementaion to buffer outbound API
// Posts from the log

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
	"sync"
	"time"
)

var (
	initialSize = 30
	logLink = "http://localhost:3000/api/log" 			// send logging information
	envLink = "http://localhost:3000/api/env"			// send environment variables (done once)
	brokerLink = "http://localhost:3000/api/broker" 	// send brokerage account infromation (done on sell/buy)
	dataLink = "http://localhost:3000/api/data" 		// send technical update information
)

// DataStream contains a queue implmentation to send API JSONs
// to the frontend
type APIBuffer struct {
	Data	[]map[string]any
	Links	[]string
	mu 		sync.Mutex
} // DataStream	

func newAPIBuffer() *APIBuffer {
	return &APIBuffer{
		Data: make([]map[string]any, 0, initialSize),
		Links: make([]string, 0, initialSize),
	}
} // NewAPIBuffer

// enqueue adds data onto the APIBuffer Object
func (queue *APIBuffer) enqueue(json map[string]any, link string) {	
	queue.mu.Lock()
	defer queue.mu.Unlock()
	if cap(queue.Data) > 2 * initialSize {
		queue.Data = copySlice(queue.Data)
		queue.Links = copyLinks(queue.Links)
	}
	queue.Data = append(queue.Data, json)
	queue.Links = append(queue.Links, link)	
} // enqueue 

// dequeue removes the first item form the APIBuffer
func (queue *APIBuffer) dequeue() (json map[string]any, link string) {
	queue.mu.Lock()
	defer queue.mu.Unlock()
	if len(queue.Data) == 0 {
		return nil, "failure"
	} // if

	data := queue.Data[0] // grab first val
	queue.Data[0] = nil // set back to intial val (to help GC)
	queue.Data = queue.Data[1:] // adjust field attribute

	dest := queue.Links[0]
	queue.Links[0] = "" 
	queue.Links = queue.Links[1:]

	return data, dest
} // dequeue

// offload will send the number of items to their destination given a 
// valid wait time
func (queue *APIBuffer) offload(times int, wait time.Duration) {
	for i := 0; i < times; i++ {
		data, dest := queue.dequeue() // NOTE: this uses mutex
		go SendPayload(data, dest)
		time.Sleep(wait * time.Millisecond)
	} // for
} // offload

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

// CopySlice copies the Slice to reduce the capacity 
// of the underlying array
func copySlice(slice []map[string]any) []map[string]any {
	copied := make([]map[string]any, len(slice))
	copy(copied, slice)
	return copied
} // CopySlice

// Combine with above method to make generic in the future
// Same function as above 
func copyLinks(slice []string) []string {
	copied := make([]string, len(slice))
	copy(copied, slice)
	return copied
}