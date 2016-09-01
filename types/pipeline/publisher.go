package pipeline

// Publisher represents an invocable publisher in the pipeline
type Publisher struct {
	ID          string                 `json:"id"`          // The ID of the Publisher
	Name        string                 `json:"name"`        // The Name of the Publisher
	SafeName    string                 `json:"safe_name"`   // The Safe name of the publisehr
	Repository  string                 `json:"repository"`  // The Repository the publisher belongs too
	Description string                 `json:"description"` // The Description of the Publisher
	Type        string                 `json:"type"`        // The Type of the Publisher
	IconURL     string                 `json:"icon"`        // The Icon of the Publisher
	Shapes      map[string]Shape       `json:"shapes"`      // The Shapes publsihed by this Publisher
	Settings    map[string]interface{} `json:"settings"`    // The settings of the Publisher
}
