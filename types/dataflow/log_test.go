package dataflow

import (
	"testing"

	"github.com/naveego/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogValidate(t *testing.T) {

	Convey("Given a log with no tenant id", t, func() {
		err := (&Log{}).Validate()
		Convey("Should return an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Should return ErrorMissingTenant", func() {
			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, ErrorMissingTenant.Code)
			So(e.Message, ShouldEqual, ErrorMissingTenant.Message)
		})
	})

	Convey("Given a log with no message id", t, func() {
		err := (&Log{TenantID: "vandelay"}).Validate()
		Convey("Should return an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Should return ErrorMissingMessage", func() {
			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, ErrorMissingMessage.Code)
			So(e.Message, ShouldEqual, ErrorMissingMessage.Message)
		})
	})
}
