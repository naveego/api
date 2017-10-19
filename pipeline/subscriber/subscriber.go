package subscriber

import (
	"sort"
	"sync"

	"github.com/sirupsen/logrus"
	pipeerrors "github.com/naveego/api/pipeline/errors"
	"github.com/naveego/api/types/pipeline"
	"github.com/naveego/errors"
)

var (
	subscribersMu sync.RWMutex
	subscribers   = make(map[string]Factory)
)

// RegisterFactory makes a data source available by the provided name.
// If RegisterFactory is called more than one time with the same name,
// or if the connector is nil, it panics.
func RegisterFactory(name string, factory Factory) {
	subscribersMu.Lock()
	defer subscribersMu.Unlock()

	if factory == nil {
		panic("connector: factory is nil")
	}
	if _, dup := subscribers[name]; dup {
		panic("connector: factory already registered with name " + name)
	}
	subscribers[name] = factory
}

// GetFactory returns a connector.Factory registered with the given name
func GetFactory(name string) (Factory, error) {
	subscribersMu.RLock()
	defer subscribersMu.RUnlock()

	c, ok := subscribers[name]
	if !ok {
		return nil, errors.NewWithCode(pipeerrors.GetConnectorError, "connector: Could not find Connector with name "+name)
	}

	return c, nil
}

// Factories returns a sorted list of the names of the registered publishers.
func Factories() []string {
	subscribersMu.RLock()
	defer subscribersMu.RUnlock()

	var list []string
	for name := range subscribers {
		list = append(list, name)
	}

	sort.Strings(list)
	return list
}

func unregisterAllFactories() {
	subscribersMu.Lock()
	defer subscribersMu.Unlock()
	subscribers = make(map[string]Factory)
}

type Factory func() Subscriber

type Context struct {
	Subscriber pipeline.SubscriberInstance // Reference to the subscriber for this operation
	Pipeline   pipeline.Pipeline
	APIToken   string // The Token to use for API calls
	Logger     *logrus.Entry
}

type Subscriber interface {

	// Init gives the subscriber an opportunity to get setup before use
	Init(ctx Context, settings map[string]interface{}) error

	// TestConnection tests the connection to the publisher
	TestConnection(ctx Context, connSettings map[string]interface{}) (bool, string, error)

	// Shapes returns the shapes that a subscribers can receive from the pipeline.
	// Shapes are a core component of the Pipline and represent self describing
	// data.
	Shapes(ctx Context) (pipeline.ShapeDefinitions, error)

	Receive(ctx Context, shape pipeline.ShapeDefinition, dataPoint pipeline.DataPoint) error

	// Dispose gives the subscriber an opportunity to cleanup after use
	Dispose(ctx Context) error
}

type Initer interface {
	Init(ctx Context) error
}
