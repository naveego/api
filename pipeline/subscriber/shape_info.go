package subscriber

import (
	"strings"

	"github.com/naveego/api/types/pipeline"
)

// PropertiesAndTypes is a map that contains the property name
// and the type.  The key will be the name of the property and
// the value will contain the type
type PropertiesAndTypes map[string]string

// ShapeInfo will contain information about the current data points
// shape, with respect to the pipeline shape for the same entity.
// This information can be used by the subcriber to alter its
// storage if necessary.
type ShapeInfo struct {
	IsNew            bool
	HasKeyChanges    bool
	HasNewProperties bool
	PreviousShape    pipeline.Shape
	Shape            pipeline.Shape
	NewKeys          []string
	NewProperties    PropertiesAndTypes
}

func (si ShapeInfo) HasChanges() bool {
	return si.IsNew || si.HasKeyChanges || si.HasNewProperties
}

// GenerateShapeInfo will determine the diffferences between an existing shape and the shape of a new
// data point.  If the new shape is a subset of the current shape it is not considered a change.  This
// is due to the fact that it does not represent a change that needs to be made in the storge system.
func GenerateShapeInfo(sub pipeline.SubscriberInstance, dataPoint pipeline.DataPoint) ShapeInfo {
	shape := dataPoint.Shape

	// create the info
	info := ShapeInfo{
		Shape: shape,
	}

	// Get the shape if we already know about it
	prevShape, ok := sub.Shapes[dataPoint.Entity]

	// If this shape does not exists previously then
	// we need to treat it as brand new
	if !ok {
		info.IsNew = true
		info.NewKeys = dataPoint.KeyNames
		info.HasNewProperties = true
		info.HasKeyChanges = true
	} else {

		// If the shape is exactly the same as the previous shape, or it is a subset of the previous shape
		// then there is no change.  We can just use the previous shape.
		if shape.PropertyHash != prevShape.PropertyHash && isSubsetOf(shape.Properties, prevShape.Properties) {
			info.Shape = prevShape
		}

		// Check the key names
		if !areSame(info.Shape.KeyNames, prevShape.KeyNames) {
			info.HasKeyChanges = true

			// Load the new keys
			info.NewKeys = []string{}
			for _, key := range info.Shape.KeyNames {
				info.NewKeys = append(info.NewKeys, key)
			}
		}
	}

	// Set the previous shape on the info
	info.PreviousShape = prevShape

	// Load any new properties
	info.NewProperties = PropertiesAndTypes{}
	for _, prop := range info.Shape.Properties {
		if !contains(prevShape.Properties, prop) {
			p := strings.Split(prop, ":")
			info.NewProperties[p[0]] = p[1]
		}
	}

	info.HasNewProperties = (len(info.NewProperties) > 0)

	return info
}

// contains is a helper function to determine if a string slice
// contains a string value
func contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

// areSame is a helper function that determines if two slices are
// the same.  Two slices are considered the same if they are the same
// length and contain equal values at the same indexes.
func areSame(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// isSubsetOf is a helper function that determines if one slice
// is a subset of another
func isSubsetOf(list []string, all []string) bool {
	for _, l := range list {
		if !contains(all, l) {
			return false
		}
	}
	return true
}
