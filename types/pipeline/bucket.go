package pipeline

// Bucket is a default location to send all data flowing
// into the pipeline.  A repository can only have one Bucket
// defined.  A bucket is also optional.
type Bucket struct {
	ID          string           `json:"-" bson:"_id,omitempty"`
	Status      PipelineStatus   `json:"status" bson:"status"`         // The Status of the pipeline
	InputStream string           `json:"input" bson:"input"`           // The input stream
	Subscriber  BucketSubscriber `json:"subscriber" bson:"subscriber"` // The subcriber bucket
}

type BucketSubscriber struct {
	Type           string                 `json:"type" bson:"type"`
	StreamEndpoint string                 `json:"stream_endpoint" bson:"stream_endpoint"`
	Shapes         map[string]Shape       `json:"shapes" bson:"shapes"`
	Settings       map[string]interface{} `json:"settings" bson:"settings"`
}
