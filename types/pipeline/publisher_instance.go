package pipeline

// PublisherInstance represents an publisher that is configured for a given repository
type PublisherInstance struct {
	ID              string                 `json:"id" bson:"_id,omitempty"`                            // The ID of the Publisher
	Name            string                 `json:"name" bson:"name"`                                   // The Name of the Publisher
	SafeName        string                 `json:"safe_name" bson:"safe_name"`                         // The Safe name of the publisehr
	Repository      string                 `json:"repository" bson:"repository"`                       // The Repository the publisher belongs too
	Description     string                 `json:"description,omitempty" bson:"description,omitempty"` // The Description of the Publisher
	Type            string                 `json:"type" bson:"type"`                                   // The Type of the Publisher
	IconURL         string                 `json:"icon" bson:"icon"`                                   // The Icon of the Publisher
	Schedule        string                 `json:"schedule,omitempty" bson:"schedule,omitempty"`       // The Schedule to run the publish
	PublishEndpoint string                 `json:"publish_endpoint" bson:"publish_endpoint"`           // The endpoint used for publishing messages
	LiveEndpoint    string                 `json:"live_endpoint" bson:"live_endpoint"`                 // The URL for connecting to Naveego Live
	Shapes          map[string]Shape       `json:"shapes" bson:"shapes"`                               // The Shapes publsihed by this Publisher
	Settings        map[string]interface{} `json:"settings" bson:"settings"`                           // The settings of the Publisher
}
