package notify

import "time"

// Subscription represents a notification registrtion for a
// subscriber.  It is used to connect subscribers to subscribed
// topics.
type Subscription struct {
	ID           string    `json:"id" bson:"_id"`
	TenantID     string    `json:"tenant_id" bson:"tenant_id"`
	Name         string    `json:"name" bson:"name"`
	Topic        string    `json:"topic" bson:"topic"`
	Subscriber   string    `json:"subscriber" bson:"subscriber"`
	RegisteredOn time.Time `json:"registered_on" bson:"registered_on"`
	Tags         []string  `json:"tags" bson:"tags"`
}
