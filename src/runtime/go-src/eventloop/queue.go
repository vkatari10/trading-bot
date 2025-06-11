package eventloop

// Queue Designed to handle outbound API JSONs in an orderly way

var (
	initialSize = 30
)

// DataStream contains a queue implmentation to send API JSONs
// to the frontend
type APIStream struct {
	Data	[]map[string]any
} // DataStream	

func NewAPIStream() *APIStream {
	return &APIStream{
		Data: make([]map[string]any, 30),
	}
} // NewAPIStream

// enqueue adds data onto the APIStream Object
func (queue *APIStream) enqueue(data map[string]any) {
	if cap(queue.Data) > 2 * initialSize {
		queue.Data = copySlice(queue.Data)
	}
	queue.Data = append(queue.Data, data)
} // enqueue 

// dequeue removes the first item form the APIStream
func (queue *APIStream) dequeue() map[string]any {
	if len(queue.Data) == 0 {
		return nil 
	} // if

	dqItem := queue.Data[0]
	queue.Data[0] = nil

	queue.Data = queue.Data[1:]
	return dqItem
} // dequeue

// CopySlice copies the Slice to reduce the capacity 
// of the underlying array
func copySlice(slice []map[string]any) []map[string]any {
	copied := make([]map[string]any, len(slice))
	copy(copied, slice)
	return copied
} // CopySlice
