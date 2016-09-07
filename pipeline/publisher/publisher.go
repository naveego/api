package publisher

import (
	"sort"
	"sync"

	"github.com/Sirupsen/logrus"
	pipeerrors "github.com/naveego/api/pipeline/errors"
	"github.com/naveego/api/types/pipeline"
	"github.com/naveego/errors"
)

var (
	publishersMu sync.RWMutex
	publishers   = make(map[string]Factory)
)

// RegisterFactory makes a data source available by the provided name.
// If RegisterFactory is called more than one time with the same name,
// or if the connector is nil, it panics.
func RegisterFactory(name string, factory Factory) {
	publishersMu.Lock()
	defer publishersMu.Unlock()

	if factory == nil {
		panic("connector: factory is nil")
	}
	if _, dup := publishers[name]; dup {
		panic("connector: factory already registered with name " + name)
	}
	publishers[name] = factory
}

// GetFactory returns a Factory registered with the given name
func GetFactory(name string) (Factory, error) {
	publishersMu.RLock()
	defer publishersMu.RUnlock()

	c, ok := publishers[name]
	if !ok {
		return nil, errors.NewWithCode(pipeerrors.GetConnectorError, "connector: Could not find Connector with name "+name)
	}

	return c, nil
}

// Factories returns a sorted list of the names of the registered publishers.
func Factories() []string {
	publishersMu.RLock()
	defer publishersMu.RUnlock()

	var list []string
	for name := range publishers {
		list = append(list, name)
	}

	sort.Strings(list)
	return list
}

func unregisterAllFactories() {
	publishersMu.Lock()
	defer publishersMu.Unlock()
	publishers = make(map[string]Factory)
}

type Factory func() Publisher

type Context struct {
	PublisherInstance pipeline.PublisherInstance // Reference to the publisher data from the expectedPublisher
	APIToken          string                     // The API token to use for authentication
	Logger            *logrus.Entry
}

// Publisher provides the core API for publishing data to the pipeline.
type Publisher interface {

	// Shapes returns the shapes that a publisher can send to the pipeline.
	// Shapes are a core component of the Pipline and represent self describing
	// data.
	Shapes(ctx Context) (map[string]pipeline.Shape, error)

	// Publish triggers the send data operation on the data source.
	// When Publish is called it will be provided with an execution
	// context, a transport for sending data points.
	Publish(ctx Context, dataTransport DataTransport)
}
