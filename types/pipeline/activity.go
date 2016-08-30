package pipeline

// Activity represents a activity in a pipeline
type Activity struct {
	ID            string                 `json:"id"`       // The ID of the node
	Type          string                 `json:"type"`     // The nodes type
	InputStreams  []string               `json:"inputs"`   // The input stream
	OutputStreams []string               `json:"outputs"`  // The outputs from this node
	Settings      map[string]interface{} `json:"settings"` // Any settings for this node
}
