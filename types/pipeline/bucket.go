package pipeline

// Bucket is a default location to send all data flowing
// into the pipeline.  A repository can only have one Bucket
// defined.  A bucket is also optional.
type Bucket struct {
	ID          string                 `json:"-" bson:"_id,omitempty"`
	Status      PipelineStatus         `json:"status" bson:"status"` // The Status of the pipeline
	InputStream string                 `json:"input" bson:"input"`   // The input stream
	Type        string                 `json:"type" bson:"type"`     // The type of subscriber to use
	Shapes      map[string]Shape       `json:"shapes" bson:"shapes"`
	Settings    map[string]interface{} `json:"settings" bson:"settings"`
}
