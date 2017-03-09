package subscriber

import (
	"testing"

	"github.com/naveego/api/types/pipeline"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	expectedSubscriber = &testSubscriber{}
)

type testSubscriber struct{}

func (t *testSubscriber) TestConnection(ctx Context, connSettings map[string]interface{}) (bool, string, error) {
	return false, "", nil
}

func (t *testSubscriber) Shapes(ctx Context) (pipeline.ShapeDefinitions, error) {
	return pipeline.ShapeDefinitions{}, nil
}

func testSubscriberFactory() Subscriber {
	return expectedSubscriber
}

func expectPanic(f func()) {
	defer func() {
		err := recover()
		So(err, ShouldNotBeNil)
	}()
	f()
}

func TestRegisterFactory(t *testing.T) {

	Convey("Should add the subscriber factory to the subscribers dictionary", t, func() {
		unregisterAllFactories()
		RegisterFactory("test", testSubscriberFactory)
		factory, ok := subscribers["test"]
		if !ok {
			t.Fatal("Expected connector to be in map")
		}
		sub := factory()
		So(sub, ShouldEqual, expectedSubscriber)
	})

	Convey("Registering a subscriber factory that is nil", t, func() {

		unregisterAllFactories()

		Convey("Should panic", func() {
			expectPanic(func() { RegisterFactory("test", nil) })
		})
	})

	Convey("Registering a subscriber factory more than once", t, func() {

		unregisterAllFactories()

		Convey("Should panic", func() {
			RegisterFactory("test", testSubscriberFactory)
			expectPanic(func() { RegisterFactory("test", testSubscriberFactory) })
		})

	})

}

func TestFactories(t *testing.T) {

	Convey("Should return the list of registered subscriber factories in alpha order", t, func() {
		unregisterAllFactories()

		RegisterFactory("test", testSubscriberFactory)
		RegisterFactory("awesome", testSubscriberFactory)
		RegisterFactory("doh", testSubscriberFactory)

		factories := Factories()
		So(factories, ShouldResemble, []string{"awesome", "doh", "test"})
	})

}

func TestGetFactory(t *testing.T) {

	Convey("Should return the registered connector", t, func() {
		unregisterAllFactories()

		subscribers["test"] = testSubscriberFactory
		factory, _ := GetFactory("test")
		sub := factory()
		So(sub, ShouldEqual, expectedSubscriber)

	})

	Convey("Given the name of a subscriberFactory that is not registered", t, func() {
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
