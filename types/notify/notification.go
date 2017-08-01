package notify

import "time"

// Notification represents a notification that can be
// published to the notify service
type Notification struct {
	NotificationID string    `json:"notification_id,omitempty" bson:"_id"`
	TenantID       string    `json:"tenant_id" bson:"tenant_id"`
	Topic          string    `json:"topic" bson:"topic"`
	Subject        string    `json:"subject" bson:"subject"`
	Message        string    `json:"message" bson:"message"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
}
