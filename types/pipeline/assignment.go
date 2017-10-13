package pipeline

import (
	"time"

	"github.com/naveego/errors"
)

// AssignmentStatus enumerates the states an Assignment can be in.
// Instances of this type are prefixed "AS"
type AssignmentStatus string

const (
	// ASPendingRun means the assigned segment has not yet been started by the agent.
	ASPendingRun AssignmentStatus = "pending-run"

	// ASRunning means the assigned segment is running on the agent.
	ASRunning AssignmentStatus = "running"

	// ASRecycling means the segment is being restarted.
	ASRecycling AssignmentStatus = "recycling"

	// ASPendingDelete means the segment has been scheduled for delete but the agent has not yet stopped it.
	ASPendingDelete AssignmentStatus = "pending-delete"

	// ASPendingDelete means the segment has been scheduled for delete but the agent has not yet stopped it.
	ASDeleted AssignmentStatus = "deleted"

	// ASError means the agent running the segment has reported that the segment has an error.
	ASError AssignmentStatus = "error"
)

// Assignment represents the assignment of a publisher or subscriber
// (or other segment) to an agent where it will run.
type Assignment struct {
	ID              string           `json:"id,omitempty" bson:"_id,omitempty"`    // The ID of the assignment.
	AgentID         string           `json:"agentID,omitempty" bson:"agentID"`     // The ID of the agent.
	SegmentID       string           `json:"segmentID,omitempty" bson:"segmentID"` // The ID of the segment.
	Type            string           `json:"type,omitempty" bson:"type"`           // The segment type, probably "publisher" or "subscriber".
	Status          AssignmentStatus `json:"status,omitempty" bson:"status"`
	CreatedAt       time.Time        `json:"created_at,omitempty" bson:"createdAt"`                        // The UTC time this assignment was created.
	StartedAt       time.Time        `json:"started_at,omitempty" bson:"startedAt,omitempty"`              // The UTC time the segment started running on the agent.
	LastHeartbeatAt time.Time        `json:"last_heartbeat_at,omitempty" bson:"lastHeartbeatAt,omitempty"` // The most recent UTC time the agent signalled that the segment was healthy.
}

// Validate returns an error if the Assigment is invalid.
func (a Assignment) Validate() error {
	if a.AgentID == "" {
		return errors.New("AgentID is required")
	}
	if a.SegmentID == "" {
		return errors.New("SegmentID is required")
	}
	if a.Type == "" {
		return errors.New("Type is required")
	}
	return nil
}
