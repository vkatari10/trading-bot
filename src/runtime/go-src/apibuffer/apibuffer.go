package apibuffer

// APIBuffer is just a Queue implementaion to buffer outbound API
// Posts from the log

import (
	"sync"
	"time"
	"fmt"
)

const (
	initialSize = 30
)

// APIBuffer contains a queue implmentation to send API JSONs
// to the frontend
type APIBuffer struct {
	Data	[]map[string]any
	Links	[]string
	mu 		sync.Mutex
} // APIBuffer

func NewAPIBuffer() *APIBuffer {
	return &APIBuffer{
		Data: make([]map[string]any, 0, initialSize),
		Links: make([]string, 0, initialSize),
	}
} // newAPIBuffer

// enqueue adds data onto the APIBuffer Object
func (queue *APIBuffer) Enqueue(json map[string]any, link string) {	
	queue.mu.Lock()
	defer queue.mu.Unlock()
	if cap(queue.Data) > 2 * initialSize {
		queue.Data = copySlice(queue.Data)
		queue.Links = copyLinks(queue.Links)
	} // if
	queue.Data = append(queue.Data, json)
	queue.Links = append(queue.Links, link)	
} // enqueue 

// dequeue removes the first item form the APIBuffer
func (queue *APIBuffer) Dequeue() (json map[string]any, link string) {
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
func (queue *APIBuffer) Flush(times int, wait time.Duration, action func(map[string]any, string)) {
	for i := 0; i < times; i++ {
		data, dest := queue.Dequeue() // NOTE: this uses mutex
		go action(data, dest)
		time.Sleep(wait * time.Millisecond)
	} // for
} // offload

// flushAll flushes all items inside a APIBuffer and waits
// for wait amount of time until flushing the next item
// to its destination
func (q *APIBuffer) FlushAll(wait time.Duration, action func(map[string]any, string)) {
	fmt.Println(q.Data)
	for _ = range q.Data {
		data, dest := q.Dequeue()
		go action(data, dest)
		time.Sleep(wait * time.Millisecond)
	}
} // flushAll

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
} // copyLinks

