package subscriber

import (
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
