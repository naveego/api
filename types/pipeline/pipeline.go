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
	ID              string         `json:"id" bson:"_id,omitempty"` // The ID of the pipeline
	Name            string         `json:"name"`                    // The Name of the pipeline
	Description     string         `json:"description,omitempty"`   // The description of the piepline
	Status          PipelineStatus `json:"status"`                  // The Status of the pipeline
	Schedule        string         `json:"schedule"`
	PublisherID     string         `json:"publisher" bson:"publisher"`
	PublisherType   string         `json:"publisher_type" bson:"publisher_type"`
	SubscriberID    string         `json:"subscriber" bson:"subscriber"`
	SubscriberType  string         `json:"subscriber_type" bson:"subscriber_type"`
	PublishedShape  string         `json:"published_shape" bson:"published_shape"`
	SubscribedShape string         `json:"subscribed_shape" bson:"subscribed_shape"`
	Mappings        []ShapeMapping `json:"mappings" bson:"mappings"`
	CreatedBy       string         `json:"created_by" bson:"created_by"`
	CreatedOn       string         `json:"created_on,omitempty" bson:"created_on"`   // The date the pipeline was created on
	ModifiedOn      string         `json:"modified_on,omitempty" bson:"modified_on"` // The date the pipeline was modified
	ModifiedBy      string         `json:"modified_by" bson:"modified_by"`
}

type ShapeMapping struct {
	From     string `json:"from" bson:"from"`
	FromType string `json:"from_type" bson:"from_type"`
	To       string `json:"to" bson:"to"`
	ToType   string `json:"to_type" bson:"to_type"`
}
