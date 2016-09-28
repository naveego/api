package pipeline

import (
	"testing"

	"github.com/naveego/errors"
	. "github.com/smartystreets/goconvey/convey"
)

type testData map[string]interface{}

func TestDataPointValidation(t *testing.T) {

	Convey("Given a valid dataPoint", t, func() {
		Convey("Should return error as nil", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate()
			So(err, ShouldBeNil)
		})
	})

	Convey("Given an invalid dataPoint", t, func() {

		Convey("When missing the repository, should return error code 4220001", func() {
			msg := DataPoint{Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220001)
		})

		Convey("When the repository is too short, should return error code 4220002", func() {
			msg := DataPoint{Repository: "te", Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220002)
		})

		Convey("When the repository is too long, should return error code 4220002", func() {
			msg := DataPoint{Repository: "thisisareponamethatistoolong", Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220002)
		})

		Convey("When the repository has invalid characters, should return error code 4220002", func() {
			msg := DataPoint{Repository: "tes#t", Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220002)
		})

		Convey("When missing the entity, should return error code 4220004", func() {
			msg := DataPoint{Repository: "test", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220004)
		})

		Convey("When the entity is too short, should return error code 4220005", func() {
			msg := DataPoint{Repository: "test", Entity: "u", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220005)
		})

		Convey("When the entity is too long, should return error code 4220005", func() {
			msg := DataPoint{Repository: "test", Entity: "userentitythatistoolonglkjasfoiwjeifwekrwalkafeajke", Action: "upsert", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220005)
		})

		Convey("When missing the action, should return error code 4220006", func() {
			msg := DataPoint{Repository: "test", Entity: "user", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220006)
		})

		Convey("When an invalid action is given, should return error code 4220007", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "stop", KeyNames: []string{"id"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220007)
		})

		Convey("When missing keyNames, should return error code 4220008", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220008)
		})

		Convey("When the keyNames is an empty array, should return error code 4220008", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", KeyNames: []string{}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220008)
		})

		Convey("When the data is missing one or more property listed in keyNames, should return error code 4220010", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", KeyNames: []string{"id", "time"}, Data: testData{"id": 1, "name": "Derek"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220010)
		})

		Convey("When missing data, should return error code 4220009", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", KeyNames: []string{"id"}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220009)
		})

		Convey("when given empty data, should return error code 4220009", func() {
			msg := DataPoint{Repository: "test", Entity: "user", Action: "upsert", KeyNames: []string{"id"}, Data: testData{}}
			err := msg.Validate().(errors.Error)
			So(err.Code, ShouldEqual, 4220009)
		})
	})

}

func TestDataPointIsShaped(t *testing.T) {
	Convey("Given a dataPoint without a shape", t, func() {
		Convey("Should return false", func() {
			msg := DataPoint{}
			isShaped := msg.IsShaped()
			So(isShaped, ShouldBeFalse)
		})
	})

	Convey("Given a dataPoint with a shape", t, func() {
		Convey("Should return true", func() {
			var props = make([]string, 0)
			msg := DataPoint{Shape: Shape{Properties: props, PropertyHash: 100}}
			isShaped := msg.IsShaped()
			So(isShaped, ShouldBeTrue)
		})
	})
}
