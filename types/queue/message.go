package queue

type Message struct {
	ID   string                 `json:"id"`   // The Messages ID
	Data map[string]interface{} `json:"data"` // The data for the Message
}
