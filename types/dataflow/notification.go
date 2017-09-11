package dataflow

import (
	"time"

	"github.com/naveego/errors"
)

// Notification represents a scheduled notification in the data flow log
type Notification struct {
	ID        string    `json:"id" bson:"_id"`
	TenantID  string    `json:"-" bson:"tenant_id"`
	Type      string    `json:"type" bson:"type"`
	Filter    string    `json:"filter" bson:"filter"`
	Topic     string    `json:"topic" bson:"topic"`
	CreatedOn time.Time `json:"created_on" bson:"created_on"`
}

// Validate checks to make sure we have valid information on
// the notification.
func (n *Notification) Validate() error {
	if n.Filter == "" {
		return errors.NewWithCode(4000001, "Invalid notification: missing filter")
	}

	if n.Topic == "" {
		return errors.NewWithCode(4000002, "Invalid notification: missing topic")
	}

	return nil
}
