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

// Publisher provides the core API for publishing data to the pipeline.
type Publisher interface {

	// TestConnection tests the connection to the publisher
	TestConnection(ctx Context, connSettings map[string]interface{}) (bool, string, error)

	// Shapes returns the shapes that a publisher can send to the pipeline.
	// Shapes are a core component of the Pipline and represent self describing
	// data.
	Shapes(ctx Context) (pipeline.ShapeDefinitions, error)

	// Publish triggers the send data operation on the data source.
	// When Publish is called it will be provided with an execution
	// context, a transport for sending data points.
	Publish(ctx Context, dataTransport DataTransport)
}

type Context struct {
	PublisherInstance pipeline.PublisherInstance // Reference to the publisher data from the expectedPublisher
	APIToken          string                     // The API token to use for authentication
	Logger            *logrus.Entry
}

// GetStringSetting is a helper function that will read a setting
// as a string, and let the caller know if it was valid or not.
func (c *Context) GetStringSetting(path string) (string, bool) {
	rawValue, ok := c.PublisherInstance.Settings[path]
	if !ok {
		return "", false
	}

	val, ok := rawValue.(string)
	if !ok || val == "" {
		return "", false
	}

	return val, true
}

func (c *Context) NewDataPoint(entity string, keyNames []string, data map[string]interface{}) pipeline.DataPoint {
	return pipeline.DataPoint{
		Repository: "",
		Entity:     entity,
		Source:     c.PublisherInstance.SourceName,
		Action:     "upsert",
		KeyNames:   keyNames,
		Data:       data,
	}
}
