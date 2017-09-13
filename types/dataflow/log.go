package dataflow

import (
	"time"

	"github.com/naveego/errors"
)

var (
	ErrorEmptyArray     = errors.Error{Code: 4001000, Message: "empty array not allowed"}
	ErrorMissingTenant  = errors.Error{Code: 4001001, Message: "missing tenant_id"}
	ErrorMissingMessage = errors.Error{Code: 4001002, Message: "missing message"}
)

// Log represents a data flow log that contains important information
// about a data flow process.
type Log struct {
	CorrelationID string       `json:"correlation_id"`
	TenantID      string       `json:"tenant_id"`
	Timestamp     time.Time    `json:"ts"`
	Resource      string       `json:"resource,omitempty"`
	ResourceID    string       `json:"resource_id,omitempty"`
	ResourceType  string       `json:"resource_type,omitempty"`
	Object        string       `json:"object,omitempty"`
	ObjectID      string       `json:"object_id,omitempty"`
	Property      string       `json:"property,omitempty"`
	Action        string       `json:"action,omitemtpy"`
	Source        string       `json:"source,omitempty"`
	SourceID      string       `json:"source_id,omitempty"`
	Level         string       `json:"level"`
	Message       string       `json:"message"`
	Job           string       `json:"job,omitempty"`
	JobID         string       `json:"job_id,omitempty"`
	Module        string       `json:"module,omitempty"`
	Host          string       `json:"host,omitempty"`
	DataPoint     *DataPoint   `json:"data_point,omitempty"`
	Data          interface{}  `json:"data,omitempty"`
	Shape         []string     `json:"shape,omitempty"`
	Error         *Error       `json:"error,omitempty"`
	Writeback     *Writeback   `json:"write_back,omitempty"`
	Merge         *Merge       `json:"merge,omitempty"`
	DataQuality   *DataQuality `json:"data_quality,omitempty"`
}

// DataPoint information about data points
type DataPoint struct {
	Type        string `json:"type,omitempty"`
	SizeInBytes int64  `json:"size_in_bytes,omitempty"`
	Key         string `json:"key,omitempty"`
}

// Error information about errors
type Error struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

// Writeback information about writebacks
type Writeback struct {
	Type        string `json:"type,omitempty"`
	Result      string `json:"result,omitempty"`
	ExecutionMS int64  `json:"execution_ms,omitempty"`
}

type Merge struct {
	PriorData interface{} `json:"prior_data,omitempty"`
}

type DataQuality struct {
	RuleName        string `json:"rule_name,omitempty"`
	RuleID          string `json:"rule_id,omitemtpy"`
	CheckName       string `json:"check_name,omitempty"`
	CheckID         string `json:"check_id,omitempty"`
	RunId           string `json:"run_id,omitempty"`
	ExceptionCount  int64  `json:"exception_count,omitempty"`
	PopulationCount int64  `json:"population_count,omitempty"`
	ExecutionMS     int64  `json:"execution_ms,omitempty"`
}

func (l *Log) Validate() error {

	if l.TenantID == "" {
		return ErrorMissingTenant
	}

	if l.Message == "" {
		return ErrorMissingMessage
	}

	return nil
}
