package pipeline

import (
	"regexp"
	"strings"

	"github.com/naveego/errors"
)

// DataPointAction represents the actions that can be taken on a pipeline dataPoint.
type DataPointAction string

const (
	// DataPointUpsert respresnts an upsert action
	DataPointUpsert DataPointAction = "upsert"
	// DataPointDelete represents a delete action
	DataPointDelete DataPointAction = "delete"
)

func (d *DataPointAction) UnmarshalJSON(bytes []byte) error {
	*d = DataPointAction(strings.ToLower(string(bytes)))
	return nil
}

var (
	validRepositoryRegex *regexp.Regexp
	validEntityRegex     *regexp.Regexp

	// Error codes start with the HTTP status codes they represent
	// followed by a more specific code
	DecodeDataPointError      = 4000001
	EncodeDataPointError      = 4000002
	NoRepositoryError         = 4220001
	InvalidRepositoryError    = 4220002
	MultipleRepositoriesError = 4220003
	NoEntityError             = 4220004
	InvalidEntityError        = 4220005
	NoActionError             = 4220006
	InvalidActionError        = 4220007
	NoKeyNamesError           = 4220008
	NoDataError               = 4220009
	DataMissingKeysError      = 4220010
)

func init() {
	// Get the name regex ready to go
	validRepositoryRegex = regexp.MustCompile("^[a-zA-Z0-9_]{3,15}$")
	validEntityRegex = regexp.MustCompile("^[a-zA-Z0-9_]{3,30}$")
}

// DataPoint represents a pipeline dataPoint that can flow through the
// system.  DataPoints
type DataPoint struct {
	Repository string                 `json:"repository"`      // The repository the dataPoint belongs to
	Entity     string                 `json:"entity"`          // The entity the data represents
	Source     string                 `json:"source"`          // optional: The source identifier of where the dataPoint is coming from
	Shape      Shape                  `json:"shape,omitempty"` // optional: The shape of the data
	Action     DataPointAction        `json:"action"`          // The action for the dataPoint
	KeyNames   []string               `json:"keyNames"`        // The list of data properties that uniquely identify the dataPoint
	Meta       map[string]string      `json:"meta,omitempty"`  // An optional map of strings for sending metadata
	Data       map[string]interface{} `json:"data"`            // The data being sent through the pipe
}

// Validate ensures that the data dataPoints is valid for processing
func (d *DataPoint) Validate() error {

	if d.Repository == "" {
		return errors.Error{Code: NoRepositoryError, Message: "no repository was defined"}
	}

	if validRepositoryRegex.MatchString(d.Repository) == false {
		return errors.Error{Code: InvalidRepositoryError, Message: "repository does not meet naming requirements"}
	}

	if d.Entity == "" {
		return errors.Error{Code: NoEntityError, Message: "no entity was defined"}
	}

	if validEntityRegex.MatchString(d.Entity) == false {
		return errors.Error{Code: InvalidEntityError, Message: "entity does not meet naming requirements"}
	}

	if d.Action == "" {
		return errors.Error{Code: NoActionError, Message: "no action was defined"}
	}

	if d.KeyNames == nil || len(d.KeyNames) == 0 {
		return errors.Error{Code: NoKeyNamesError, Message: "keyNames was either not provided or is empty"}
	}

	if d.Action != DataPointUpsert && d.Action != DataPointDelete {
		return errors.Error{Code: InvalidActionError, Message: "invalid action"}
	}

	if d.Data == nil || len(d.Data) == 0 {
		return errors.Error{Code: NoDataError, Message: "data was either not provided or is empty"}
	}

	if hasProperties(d.KeyNames, d.Data) == false {
		return errors.Error{Code: DataMissingKeysError, Message: "one or more of the keyNames was not provided in the data"}
	}

	return nil
}

// IsShaped returns whether or not this dataPoint has shape information
func (d *DataPoint) IsShaped() bool {
	return d.Shape.Properties != nil && d.Shape.PropertyHash != 0
}

// verify the the specified keys exist in the data as properties
func hasProperties(keyNames []string, data map[string]interface{}) bool {
	for _, key := range keyNames {
		if _, ok := data[key]; ok == false {
			return false
		}
	}
	return true
}
