package pipeline

// SubscriberInstance represents an invocable subscriber that is defined inside a specific repository
type SubscriberInstance struct {
	ID            string                 `json:"id" bson:"_id,omitempty"`    // The ID of the Subscriber
	Name          string                 `json:"name"`                       // The Name of the Subscriber
	SafeName      string                 `json:"safe_name" bson:"safe_name"` // The Safe name of the Subscriber
	Description   string                 `json:"description"`                // The Description of the Subscriber
	Type          string                 `json:"type"`                       // The Type of the Subscriber
	IconURL       string                 `json:"icon" bson:"icon"`
	Shape         ShapeDefinition        `json:"shape" bson:"shape"` // The Shape accepted by this Subscriber
	Settings      map[string]interface{} `json:"settings"`           // The settings of the Subscriber
	PluginVersion string                 `json:"pluginVersion"`      // The identifier for the plugin used to run this subscriber.
}
