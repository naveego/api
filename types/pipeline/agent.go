package pipeline

import (
	"time"
)

// Agent contains information about an agent which can run pipeline segments.
type Agent struct {
	ID            string      `json:"id,omitempty" bson:"_id,omitempty"` // The ID of the Agent
	Host          string      `json:"host,omitempty"`
	SubscriberIDs []string    `json:"subscribers,omitempty"` // The IDs of the subscribers this agent should run
	PublisherIDs  []string    `json:"publishers,omitempty"`  // The IDs of the subscribers this agent should run
	Status        AgentStatus `json:"status,omitempty"`
}

// Agent status contains the current run status of the agent.
type AgentStatus struct {
	Host        string    `json:"host,omitempty"`
	LastStarted time.Time `json:"last_started,omitempty"`
}
