package pipeline

import (
	"hash/crc32"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/naveego/errors"
)

var castagnoliTable = crc32.MakeTable(crc32.Castagnoli) // see http://golang.org/pkg/hash/crc32/#pkg-constants

// Shape is used to maintain type information about the data contained in the dataPoint.  Shape information
// may be provided from the producer, but it is not required.  The pipeline will generate type information
// automatically based on the data itself.
type Shape struct {
	Properties   []string `json:"properties,omitempty"`   // An array of properties including type, the form of [name]:[type]
	PropertyHash uint32   `json:"propertyHash,omitempty"` // A hash used to determine if the properties have changed
}

func NewShapeFromProperties(properties []string) (Shape, error) {

	shape := Shape{}

	sort.Sort(sortByPropName(properties))

	// We are using a CRC check sum because it is very
	// efficient.  We are simply looking for a change,
	// we are not giving an identity.  Therefore, we don't
	// have to be concered about collisions.
	crcStr := ""
	propLen := len(properties)

	// We are using a lower case value for the properties
	// in order to allow for case in-sensitivity.
	for i, prop := range properties {
		crcStr = crcStr + strings.ToLower(prop)

		if i < (propLen - 1) {
			crcStr = crcStr + ","
		}
	}

	crc := crc32.New(castagnoliTable)

	if _, err := crc.Write([]byte(crcStr)); err != nil {
		return shape, err
	}

	shape.Properties = properties
	shape.PropertyHash = crc.Sum32()

	return shape, nil

}

// Shaper determines the schema of a given data point.  It will read through all the properties
// and return a Shape. This shape can be used to determine if the set of properties has changed
// between data points.
type Shaper interface {
	GetShape(data map[string]interface{}) (Shape, error) // Gets the shape of a given data structure
}

type shaper struct {
}

type sortByPropName []string

func (s sortByPropName) Len() int {
	return len(s)
}

func (s sortByPropName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByPropName) Less(i, j int) bool {
	s1Parts := strings.Split(s[i], ":")
	s2Parts := strings.Split(s[j], ":")

	s1Prop := strings.ToLower(s1Parts[0])
	s2Prop := strings.ToLower(s2Parts[0])

	return strings.Compare(s1Prop, s2Prop) < 0
}

// NewShaper creates a new instance of the default shaper.
func NewShaper() Shaper {
	return &shaper{}
}

func (s *shaper) GetShape(data map[string]interface{}) (shape Shape, err error) {

	var properties []string

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown error")
			}
			shape = Shape{}
		}
	}()

	getShapeRecursive(&properties, "", data)

	shape, err = NewShapeFromProperties(properties)

	return shape, err
}

func getShapeRecursive(properties *[]string, prefix string, data map[string]interface{}) {

	for key, val := range data {

		if strings.Contains(key, ":") || strings.Contains(key, ",") {
			panic("Invalid character found in property '" + key + "'.")
		}

		propName := getPropertyName(key, prefix)

		switch x := val.(type) {
		case string:
			if len(x) > 0 && unicode.IsDigit(rune(x[0])) && isDate(x) {
				*properties = append(*properties, propName+":date")
			} else {
				*properties = append(*properties, propName+":string")
			}
		case int, int8, int16, int32, int64, float32, float64:
			*properties = append(*properties, propName+":number")
		case bool:
			*properties = append(*properties, propName+":bool")
		case map[string]interface{}:
			*properties = append(*properties, propName+":object")
			getShapeRecursive(properties, propName, val.(map[string]interface{}))
		default:
			*properties = append(*properties, propName+":unknown")
		}
	}

}

func getPropertyName(name string, prefix string) string {
	propName := name
	if prefix != "" {
		propName = prefix + "." + propName
	}

	return propName
}

func isDate(val string) bool {

	if _, err := time.Parse(time.RFC3339, val); err == nil {
		return true
	}

	if _, err := time.Parse(time.RFC3339Nano, val); err == nil {
		return true
	}

	if _, err := time.Parse(time.RFC822, val); err == nil {
		return true
	}

	if _, err := time.Parse(time.RFC822Z, val); err == nil {
		return true
	}

	return false
}