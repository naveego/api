package pipeline

const (

	// PipelineActive represents the status of an active pipeline
	PipelineActive PipelineStatus = "active"

	// PipelinePaused represents the status of a paused pipeline
	PipelinePaused PipelineStatus = "paused"

	// PipelineDeleted represents the status of a deleted pipeline
	PipelineDeleted PipelineStatus = "deleted"
)

// PipelineStatus represents the status of a pipeline
type PipelineStatus string

// Pipeline is the core structure for defining how data is processed
// in Naveego pipeline.  It controls the flow of data from one activity
// to another.
type Pipeline struct {
	ID          string           `json:"id" bson:"_id,omitempty"` // The ID of the pipeline
	Name        string           `json:"name"`                    // The Name of the pipeline
	Description string           `json:"description,omitempty"`   // The description of the piepline
	Status      PipelineStatus   `json:"status"`                  // The Status of the pipeline
	Publisher   PublisherStream  `json:"publisher"`               // The Publisher for the pipeline
	Subscriber  SubscriberStream `json:"subscriber"`              // The Subscriber for the pipeline
	Activities  []Activity       `json:"activities"`              // The activities in the pipeline
	CreatedOn   string           `json:"created_on,omitempty"`    // The date the pipeline was created on
	ModifiedOn  string           `json:"modified_on,omitempty"`   // The date the pipeline was modified
}

type PublisherStream struct {
	ID           string `json:"publisher_id" bson:"publisher_id"` // The ID of the Publisher
	OutputStream string `json:"output" bson:"output"`             // The stream that the publisher will publish to
}

type SubscriberStream struct {
	ID          string `json:"subscriber_id" bson:"subscriber_id"` // The ID of the Subscriber
	InputStream string `json:"input" bson:"input"`                 // The input stream to send data to
}
