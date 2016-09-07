package activity

import (
	"sort"
	"sync"

	pipeerrors "github.com/naveego/api/pipeline/errors"
	"github.com/naveego/api/types/pipeline"
	"github.com/naveego/errors"
)

var (
	activitiesMU sync.RWMutex
	activities   = make(map[string]ActivityFactory)
)

// ActivityFactory is a function that is used to create
// new instances of activities.
type ActivityFactory func() Activity

type OutputCollector interface {
	Emit(dataPoint pipeline.DataPoint) error
}

// Activity represents a core interface for building pipelines.  Activities
// can will receive data from an input, and can optionally export data using
// output channels
type Activity interface {

	// Execute executes the logic in the activity
	Execute(context Context, dataPoint pipeline.DataPoint) error
}

// ActivityIniter can be implemented by activities that need to be
// initialized before they are executed.  Init will be called one time
// per istantiation of the activity.  This is typically done durring startup
// time.
type ActivityIniter interface {
	Init(settings map[string]interface{}) error
}

// RegisterActivityFactory registers a node with the system so it can be used
// by a pipeline.
func RegisterActivityFactory(name string, factory ActivityFactory) {
	activitiesMU.Lock()
	defer activitiesMU.Unlock()

	if factory == nil {
		panic("pipeline: activity factory is nil")
	}

	if _, dup := activities[name]; dup {
		panic("pipeline: there is already an activity factory registered with name " + name)
	}

	activities[name] = factory
}

// GetActivityFactory returns the ActivityNode registered with the given name.
func GetActivityFactory(name string) (ActivityFactory, error) {
	activitiesMU.RLock()
	defer activitiesMU.RUnlock()

	node, ok := activities[name]
	if !ok {
		return nil, errors.NewWithCode(pipeerrors.GetActivityNodeError, "pipeline: could not find activity factory with name "+name)
	}

	return node, nil
}

// ActivityFactories returns a sorted list of the names of the registered activity nodes.
func ActivityFactories() []string {
	activitiesMU.RLock()
	defer activitiesMU.RUnlock()

	var list []string
	for name := range activities {
		list = append(list, name)
	}

	sort.Strings(list)
	return list
}

func unregisterAllActivityFactories() {
	activitiesMU.Lock()
	defer activitiesMU.Unlock()
	activities = make(map[string]ActivityFactory)
}
