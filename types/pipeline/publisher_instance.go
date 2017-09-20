package pipeline

// PublisherInstance represents an publisher that is configured for a given repository
type PublisherInstance struct {
	ID            string                 `json:"id" bson:"_id,omitempty"`                            // The ID of the Publisher
	Name          string                 `json:"name" bson:"name"`                                   // The Name of the Publisher
	SourceName    string                 `json:"source_name" bson:"source_name"`                     // The Source name of the publisehr
	Description   string                 `json:"description,omitempty" bson:"description,omitempty"` // The Description of the Publisher
	Type          string                 `json:"type" bson:"type"`                                   // The Type of the Publisher
	IconURL       string                 `json:"icon" bson:"icon"`                                   // The Icon of the Publisher
	Schedule      string                 `json:"schedule,omitempty" bson:"schedule,omitempty"`       // The Schedule to run the publish
	LiveEndpoint  string                 `json:"live_endpoint" bson:"live_endpoint"`                 // The URL for connecting to Naveego Live
	LogEndpoint   string                 `json:"log_endpoint" bson:"log_endpoint"`
	Shapes        []string               `json:"shapes" bson:"shapes"`     // The Shapes publsihed by this Publisher
	Settings      map[string]interface{} `json:"settings" bson:"settings"` // The settings of the Publisher
	PluginVersion string                 `json:"pluginVersion"`            // The identifier for the plugin used to run this subscriber.
}
