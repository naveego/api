package dataflow

import (
	"testing"

	"github.com/naveego/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Notification_Validate(t *testing.T) {

	Convey("Given a notification with a blank tenant_id", t, func() {
		note := &Notification{
			TenantID: "",
			Label:    "My Notification",
			Filter:   "level:info",
			Topic:    "nrn:test",
		}
		err := note.Validate()

		Convey("should return a validation error", func() {
			So(err, ShouldNotBeNil)

			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, 4000001)
		})
	})

	Convey("Given a notification with a blank label", t, func() {
		note := &Notification{
			TenantID: "vandelay",
			Label:    "",
			Filter:   "level:info",
			Topic:    "nrn:test",
		}
		err := note.Validate()

		Convey("should return a validation error", func() {
			So(err, ShouldNotBeNil)

			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, 4000002)
		})
	})

	Convey("Given a notification with a blank filter", t, func() {
		note := &Notification{
			TenantID: "vandelay",
			Label:    "My Notification",
			Filter:   "",
			Topic:    "nrn:test",
		}
		err := note.Validate()

		Convey("should return a validation error", func() {
			So(err, ShouldNotBeNil)

			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, 4000003)
		})
	})

	Convey("Given a notification with a blank topic", t, func() {
		note := &Notification{
			TenantID: "vandelay",
			Label:    "My Notification",
			Filter:   "level:info",
			Topic:    "",
		}
		err := note.Validate()

		Convey("should return a validation error", func() {
			So(err, ShouldNotBeNil)

			e, _ := err.(errors.Error)
			So(e.Code, ShouldEqual, 4000004)
		})
	})

}
