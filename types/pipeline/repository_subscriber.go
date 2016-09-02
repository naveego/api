package pipeline

// RepositorySubscriber represents an invocable subscriber that is defined inside a specific repository
type RepositorySubscriber struct {
	ID          string                 `json:"id" bson:"_id,omitempty"` // The ID of the Subscriber
	Name        string                 `json:"name"`                    // The Name of the Subscriber
	Description string                 `json:"description"`             // The Description of the Subscriber
	Type        string                 `json:"type"`                    // The Type of the Subscriber
	InputStream string                 `json:"input"`                   // The input stream to consume
	Shapes      map[string]Shape       `json:"shapes"`                  // The Shapes accepted by this Subscriber
	Settings    map[string]interface{} `json:"settings"`                // The settings of the Subscriber
}
