package pipeline

import (
	"time"
)

// Agent contains information about an agent which can run pipeline segments.
type Agent struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"` // The ID of the Agent.
	Host         string    `json:"host,omitempty"`                    // The Host the Agent is running on.
	OS           string    `json:"os,omitempty"`
	Arch         string    `json:"arch,omitempty"`
	Capabilities []string  `json:"capabilities,omitempty"`
	OnPremises   bool      `json:"onPremise,omitempty"` // True if the Agent is running within the tenant's firewall.
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
