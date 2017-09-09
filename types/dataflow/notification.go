package dataflow

import "time"

// Notification represents a scheduled notification in the data flow log
type Notification struct {
	ID        string    `json:"id" bson:"_id"`
	TenantID  string    `json:"-" bson:"tenant_id"`
	Filter    string    `json:"filter" bson:"filter"`
	Topic     string    `json:"topic" bson:"topic"`
	CreatedOn time.Time `json:"created_on" bson:"created_on"`
}
