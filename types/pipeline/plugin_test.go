package pipeline

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPluginSelector(t *testing.T) {
	Convey("Given a selector", t, func() {
		sut := PluginSelector{"test", "0.1.2"}
		Convey("Should be able to string it", func() {
			actual := sut.String()
			So(actual, ShouldEqual, "test@0.1.2")
		})
		Convey("Should be able to round trip it", func() {
			actual, err := NewPluginSelector("test@0.1.2")
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, sut)
		})
	})
	Convey("Should error if string invalid", t, func() {
		_, err := NewPluginSelector("test0.1.2")
		So(err, ShouldNotBeNil)
	})
}
