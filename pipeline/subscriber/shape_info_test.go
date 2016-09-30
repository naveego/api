package subscriber

import (
	"testing"

	"github.com/naveego/api/types/pipeline"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	knownShapes = map[string]pipeline.Shape{}

	testShape = pipeline.Shape{
		KeyNames:     []string{"id"},
		KeyNamesHash: 1,
		Properties:   []string{"age:number", "id:number", "name:string"},
		PropertyHash: 2,
	}

	testShapeNoAge = pipeline.Shape{
		KeyNames:     []string{"id"},
		KeyNamesHash: 1,
		Properties:   []string{"id:number", "name:string"},
		PropertyHash: 3,
	}

	testDataPoint = pipeline.DataPoint{
		Repository: "test",
		Source:     "pub-1",
		Entity:     "user",
		KeyNames:   []string{"id"},
		Action:     "upsert",
		Shape:      testShape,
		Data: map[string]interface{}{
			"id":   1,
			"name": "Test user",
			"age":  23,
		},
	}

	testDataPointNoAge = pipeline.DataPoint{
		Repository: "test",
		Source:     "pub-1",
		Entity:     "user",
		KeyNames:   []string{"id"},
		Action:     "upsert",
		Shape:      testShapeNoAge,
		Data: map[string]interface{}{
			"id":   1,
			"name": "Test user",
		},
	}
)

func TestGenerateShapeInfo(t *testing.T) {

	Convey("Given a data point with a shape that does not exists in Subscriber.Shapes", t, func() {

		shapeInfo := GenerateShapeInfo(knownShapes, testDataPoint)

		Convey("Should return a shape info", func() {
			Convey("with IsNew = true", func() {
				So(shapeInfo.IsNew, ShouldBeTrue)
			})
			Convey("where HasChanges() returns true", func() {
				So(shapeInfo.HasChanges(), ShouldBeTrue)
			})
			Convey("with HasKeyChanges = true", func() {
				So(shapeInfo.HasKeyChanges, ShouldBeTrue)
			})
			Convey("with HasNewProperties = true", func() {
				So(shapeInfo.HasNewProperties, ShouldBeTrue)
			})
			Convey("with PreviousShape set to empty shape", func() {
				So(shapeInfo.PreviousShape.KeyNames, ShouldBeEmpty)
				So(shapeInfo.PreviousShape.KeyNamesHash, ShouldEqual, 0)
				So(shapeInfo.PreviousShape.Properties, ShouldBeEmpty)
				So(shapeInfo.PreviousShape.PropertyHash, ShouldEqual, 0)
			})
			Convey("with Shape set to the dataPoints shape", func() {
				So(shapeInfo.Shape, ShouldResemble, testShape)
			})
			Convey("with NewKeys = ['id']", func() {
				So(shapeInfo.NewKeys, ShouldResemble, []string{"id"})
			})
			Convey("with NewProperties = ['age':'number','id':'number','name':'string']", func() {
				So(shapeInfo.NewProperties, ShouldResemble, PropertiesAndTypes{
					"age":  "number",
					"id":   "number",
					"name": "string",
				})
			})
		})

	})

	Convey("Given a data point with a shape that is exactly the same as an existing shape in Subscriber.Shapes", t, func() {

		knownShapes[testDataPoint.Entity] = testShape
		shapeInfo := GenerateShapeInfo(knownShapes, testDataPoint)

		Reset(func() {
			knownShapes = map[string]pipeline.Shape{}
		})

		Convey("Should return a shape info", func() {
			Convey("with IsNew = false", func() {
				So(shapeInfo.IsNew, ShouldBeFalse)
			})
			Convey("where HasChanges() returns false", func() {
				So(shapeInfo.HasChanges(), ShouldBeFalse)
			})
			Convey("with HasKeyChanges = false", func() {
				So(shapeInfo.HasKeyChanges, ShouldBeFalse)
			})
			Convey("with HasNewProperties = false", func() {
				So(shapeInfo.HasNewProperties, ShouldBeFalse)
			})
			Convey("with PreviousShape set to existing shape", func() {
				So(shapeInfo.PreviousShape, ShouldResemble, testShape)
			})
			Convey("with Shape set to data points shape", func() {
				So(shapeInfo.Shape, ShouldResemble, testShape)
			})
			Convey("with NewKeys set to empty array", func() {
				So(shapeInfo.NewKeys, ShouldBeEmpty)
			})
			Convey("with NewProperties set to empty array", func() {
				So(shapeInfo.NewProperties, ShouldBeEmpty)
			})
		})

	})

	Convey("Given a data point with a shape that has new properties", t, func() {
		knownShapes[testDataPoint.Entity] = testShapeNoAge
		shapeInfo := GenerateShapeInfo(knownShapes, testDataPoint)

		Reset(func() {
			knownShapes = map[string]pipeline.Shape{}
		})

		Convey("Should return a shape info", func() {
			Convey("with IsNew = false", func() {
				So(shapeInfo.IsNew, ShouldBeFalse)
			})
			Convey("where HasChanges() returns true", func() {
				So(shapeInfo.HasChanges(), ShouldBeTrue)
			})
			Convey("with HasKeyChanges = false", func() {
				So(shapeInfo.HasKeyChanges, ShouldBeFalse)
			})
			Convey("with HasNewProperties = true", func() {
				So(shapeInfo.HasNewProperties, ShouldBeTrue)
			})
			Convey("with PreviousShape set to exising shape", func() {
				So(shapeInfo.PreviousShape, ShouldResemble, testShapeNoAge)
			})
			Convey("with Shape set to data points shape", func() {
				So(shapeInfo.Shape, ShouldResemble, testShape)
			})
			Convey("with NewKeys set to empty array", func() {
				So(shapeInfo.NewKeys, ShouldBeEmpty)
			})
			Convey("with NewProperties = ['age':'number']", func() {
				So(shapeInfo.NewProperties, ShouldResemble, PropertiesAndTypes{
					"age": "number",
				})
			})
		})
	})

	Convey("Given a data point with a shape that has fewer properties than existing shape", t, func() {
		knownShapes[testDataPointNoAge.Entity] = testShape
		shapeInfo := GenerateShapeInfo(knownShapes, testDataPointNoAge)

		Reset(func() {
			knownShapes = map[string]pipeline.Shape{}
		})

		Convey("Should return a shape info", func() {
			Convey("with IsNew = false", func() {
				So(shapeInfo.IsNew, ShouldBeFalse)
			})
			Convey("where HasChanges() returns false", func() {
				So(shapeInfo.HasChanges(), ShouldBeFalse)
			})
			Convey("with HasKeyChanges = false", func() {
				So(shapeInfo.HasKeyChanges, ShouldBeFalse)
			})
			Convey("with HasNewProperties = false", func() {
				So(shapeInfo.HasNewProperties, ShouldBeFalse)
			})
			Convey("with PreviousShape set to existing shape", func() {
				So(shapeInfo.PreviousShape, ShouldResemble, testShape)
			})
			Convey("with Shape set to data points shape", func() {
				So(shapeInfo.Shape, ShouldResemble, testShape)
			})
			Convey("with NewKeys set to empty array", func() {
				So(shapeInfo.NewKeys, ShouldBeEmpty)
			})
			Convey("with NewProperties set to empty array", func() {
				So(shapeInfo.NewProperties, ShouldBeEmpty)
			})
		})
	})

	Convey("Given a data point with different keys", t, func() {
		knownShapes[testDataPoint.Entity] = testShape
		testDataPoint.Shape = *(&testShape)
		testDataPoint.Shape.KeyNames = []string{"name"}
		shapeInfo := GenerateShapeInfo(knownShapes, testDataPoint)

		Reset(func() {
			knownShapes = map[string]pipeline.Shape{}
		})

		Convey("Should return a shape info", func() {
			Convey("with IsNew = false", func() {
				So(shapeInfo.IsNew, ShouldBeFalse)
			})
			Convey("where HasChanges() returns true", func() {
				So(shapeInfo.HasChanges(), ShouldBeTrue)
			})
			Convey("with HasKeyChanges = true", func() {
				So(shapeInfo.HasKeyChanges, ShouldBeTrue)
			})
			Convey("with HasNewProperties = false", func() {
				So(shapeInfo.HasNewProperties, ShouldBeFalse)
			})
			Convey("with PreviousShape set to existing shape", func() {
				So(shapeInfo.PreviousShape, ShouldResemble, testShape)
			})
			Convey("with Shape set to data points shape", func() {
				So(shapeInfo.Shape, ShouldResemble, testDataPoint.Shape)
			})
			Convey("with NewKeys = 'name'", func() {
				So(shapeInfo.NewKeys, ShouldResemble, []string{"name"})
			})
			Convey("with NewProperties set to empty array", func() {
				So(shapeInfo.NewProperties, ShouldBeEmpty)
			})
		})
	})

}

func buildShape(dataPoint pipeline.DataPoint) pipeline.Shape {
	s, _ := pipeline.NewShaper().GetShape([]string{"id"}, dataPoint.Data)
	return s
}
