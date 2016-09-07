package pipeline

// PublisherInstance represents an publisher that is configured for a given repository
type PublisherInstance struct {
	ID          string                 `json:"id" bson:"_id,omitempty"`    // The ID of the Publisher
	Name        string                 `json:"name"`                       // The Name of the Publisher
	SafeName    string                 `json:"safe_name" bson:"safe_name"` // The Safe name of the publisehr
	Repository  string                 `json:"repository"`                 // The Repository the publisher belongs too
	Description string                 `json:"description,omitempty"`      // The Description of the Publisher
	Type        string                 `json:"type"`                       // The Type of the Publisher
	IconURL     string                 `json:"icon"`                       // The Icon of the Publisher
	Schedule    string                 `json:"schedule,omitempty"`         // The Schedule to run the publish
	LiveURL     string                 `json:"live_url" bson:"live_url"`   // The URL for connecting to Naveego Live
	Shapes      map[string]Shape       `json:"shapes"`                     // The Shapes publsihed by this Publisher
	Settings    map[string]interface{} `json:"settings"`                   // The settings of the Publisher
}
