package notify

import "time"

// Notification represents a notification that can be
// published to the notify service
type Notification struct {
	NotificationID string    `json:"notification_id,omitempty" bson:"_id"`
	TenantID       string    `json:"-" bson:"tenant_id"`
	Topic          string    `json:"topic" bson:"topic"`
	Default        string    `json:"default" bson:"default"`
	Email          string    `json:"email,omitempty" bson:"email,omitempty"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
	Status         string    `json:"-" bson:"status"`
}
