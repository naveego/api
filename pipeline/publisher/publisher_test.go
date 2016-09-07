package publisher

import (
	"testing"

	"github.com/naveego/api/types/pipeline"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	expectedPublisher = &testPublisher{}
)

type testPublisher struct{}

func (t *testPublisher) Shapes(ctx Context) (map[string]pipeline.Shape, error) {
	return nil, nil
}

func (t *testPublisher) Publish(ctx Context, dataTransport DataTransport) {}

func testPublisherFactory() Publisher {
	return expectedPublisher
}

func expectPanic(f func()) {
	defer func() {
		err := recover()
		So(err, ShouldNotBeNil)
	}()
	f()
}

func TestRegisterFactory(t *testing.T) {

	Convey("Should add the publisher factory to the publishers dictionary", t, func() {
		unregisterAllFactories()
		RegisterFactory("test", testPublisherFactory)
		factory, ok := publishers["test"]
		if !ok {
			t.Fatal("Expected connector to be in map")
		}
		sub := factory()
		So(sub, ShouldEqual, expectedPublisher)
	})

	Convey("Registering a publisher factory that is nil", t, func() {

		unregisterAllFactories()

		Convey("Should panic", func() {
			expectPanic(func() { RegisterFactory("test", nil) })
		})
	})

	Convey("Registering a publisher factory more than once", t, func() {

		unregisterAllFactories()

		Convey("Should panic", func() {
			RegisterFactory("test", testPublisherFactory)
			expectPanic(func() { RegisterFactory("test", testPublisherFactory) })
		})

	})

}

func TestFactories(t *testing.T) {

	Convey("Should return the list of registered subscriber factories in alpha order", t, func() {
		unregisterAllFactories()

		RegisterFactory("test", testPublisherFactory)
		RegisterFactory("awesome", testPublisherFactory)
		RegisterFactory("doh", testPublisherFactory)

		factories := Factories()
		So(factories, ShouldResemble, []string{"awesome", "doh", "test"})
	})

}

func TestGetFactory(t *testing.T) {

	Convey("Should return the registered connector", t, func() {
		unregisterAllFactories()

		publishers["test"] = testPublisherFactory
		factory, _ := GetFactory("test")
		sub := factory()
		So(sub, ShouldEqual, expectedPublisher)

	})

	Convey("Given the name of a publisher factory that is not registered", t, func() {
		unregisterAllFactories()

		c, err := GetFactory("notregistered")
		Convey("Should return an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Should return nil for Connector", func() {
			So(c, ShouldBeNil)
		})
	})

}
