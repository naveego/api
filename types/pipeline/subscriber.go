package pipeline

// Subscriber represents an invocable subscriber in the pipeline
type Subscriber struct {
	ID          string                 `json:"id"`          // The ID of the Subscriber
	Name        string                 `json:"name"`        // The Name of the Subscriber
	Description string                 `json:"description"` // The Description of the Subscriber
	Type        string                 `json:"type"`        // The Type of the Subscriber
	Shapes      map[string]Shape       `json:"shapes"`      // The Shapes accepted by this Subscriber
	Settings    map[string]interface{} `json:"settings"`    // The settings of the Subscriber
}
