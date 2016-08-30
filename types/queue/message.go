package queue

type Message struct {
	ID      int32                  `json:"id"`    // The ID of the messages
	QueueID string                 `json:"queue"` // The QueueID that the message belongs to
	Data    map[string]interface{} `json:"data"`  // The Data for the message
}
