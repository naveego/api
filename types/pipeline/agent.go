package pipeline

import (
	"time"
)

// Agent contains information about an agent which can run pipeline segments.
type Agent struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"` // The ID of the Agent.
	Host      string    `json:"host,omitempty"`                    // The Host the Agent is running on.
	Name      string    `json:"name,omitempty"`                    // The Name of the host
	OS        string    `json:"os,omitempty"`
	Arch      string    `json:"arch,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
}
