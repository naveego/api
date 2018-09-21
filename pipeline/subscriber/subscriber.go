package subscriber

import (
	"sort"
	"sync"

	pipeerrors "github.com/naveego/api/pipeline/errors"
	"github.com/naveego/api/types/pipeline"
	"github.com/naveego/errors"
	"github.com/sirupsen/logrus"
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
	APIToken   string                      // The Token to use for API calls
	Logger     *logrus.Entry
}

// Subscriber represents the API for receiving data from the pipeline.
type Subscriber interface {
	Receive(ctx Context, shapeInfo ShapeInfo, dataPoint pipeline.DataPoint)
}

type Initer interface {
	Init(ctx Context) error
}
