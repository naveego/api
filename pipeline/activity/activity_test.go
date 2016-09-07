package activity

import (
	"testing"

	"github.com/naveego/api/types/pipeline"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	expectedActivity = &testActivity{}
)

type testActivity struct{}

func (t *testActivity) Execute(ctx Context, dataPoint pipeline.DataPoint) error {
	return nil
}

func testActivityFactory() Activity {
	return expectedActivity
}

func expectPanic(f func()) {
	defer func() {
		err := recover()
		So(err, ShouldNotBeNil)
	}()
	f()
}

func TestRegisterActivityFactory(t *testing.T) {

	Convey("Should add the activity factory to the publishers dictionary", t, func() {
		unregisterAllActivityFactories()
		RegisterActivityFactory("test", testActivityFactory)
		factory, ok := activities["test"]
		if !ok {
			t.Fatal("Expected activity to be in map")
		}
		act := factory()
		So(act, ShouldEqual, expectedActivity)
	})

	Convey("Registering a activity factory that is nil", t, func() {

		unregisterAllActivityFactories()

		Convey("Should panic", func() {
			expectPanic(func() { RegisterActivityFactory("test", nil) })
		})
	})

	Convey("Registering a activity factory more than once", t, func() {

		unregisterAllActivityFactories()

		Convey("Should panic", func() {
			RegisterActivityFactory("test", testActivityFactory)
			expectPanic(func() { RegisterActivityFactory("test", testActivityFactory) })
		})

	})

}

func TestActivityFactories(t *testing.T) {

	Convey("Should return the list of registered activity factories in alpha order", t, func() {
		unregisterAllActivityFactories()

		RegisterActivityFactory("test", testActivityFactory)
		RegisterActivityFactory("awesome", testActivityFactory)
		RegisterActivityFactory("doh", testActivityFactory)

		factories := ActivityFactories()
		So(factories, ShouldResemble, []string{"awesome", "doh", "test"})
	})

}

func TestGetActivityFactory(t *testing.T) {

	Convey("Should return the registered activity factory", t, func() {
		unregisterAllActivityFactories()

		activities["test"] = testActivityFactory
		factory, _ := GetActivityFactory("test")
		act := factory()
		So(act, ShouldEqual, expectedActivity)

	})

	Convey("Given the name of a activity factory that is not registered", t, func() {
		unregisterAllActivityFactories()

		c, err := GetActivityFactory("notregistered")
		Convey("Should return an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Should return nil for activity factory", func() {
			So(c, ShouldBeNil)
		})
	})

}
