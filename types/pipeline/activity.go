package types

// ActivityFactory is a function that is used to create
// new instances of activities.
type ActivityFactory func() Activity

// Activity represents a core interface for building pipelines.  Activities
// can will receive data from an input, and can optionally export data using
// output channels
type Activity interface {

	// Execute executes the logic in the activity
	Execute(context ActivityContext, dataPoint DataPoint) error
}

// ActivityIniter can be implemented by activities that need to be
// initialized before they are executed.  Init will be called one time
// per istantiation of the activity.  This is typically done durring startup
// time.
type ActivityIniter interface {
	Init(settings map[string]interface{}) error
}

// ActivityReference represents a activity in a pipeline
type ActivityReference struct {
	ID            string                 `json:"id"`       // The ID of the node
	Type          string                 `json:"type"`     // The nodes type
	InputStreams  []string               `json:"inputs"`   // The input stream
	OutputStreams []string               `json:"outputs"`  // The outputs from this node
	Settings      map[string]interface{} `json:"settings"` // Any settings for this node
}

type ActivityContext struct {
	Pipeline        Pipeline          // The Pipeline being executed
	Activity        ActivityReference // The currently executing Activity
	OutputCollector OutputCollector   // The output collector for the activity
}

type OutputCollector interface {
	Emit(dataPoint DataPoint) error
}
