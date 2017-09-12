package notify

import "time"

// Subscription represents a notification registrtion for a
// subscriber.  It is used to connect subscribers to subscribed
// topics.
type Subscription struct {
	ID           string    `json:"id" bson:"_id"`
	TenantID     string    `json:"-" bson:"tenant_id"`
	SubscriberID string    `json:"subscriber_id" bson:"subscriber_id"`
	Topic        string    `json:"topic" bson:"topic"`
	Label        string    `json:"label" bson:"label"`
	Methods      []Method  `json:"methods" bson:"methods"`
	RegisteredOn time.Time `json:"registered_on" bson:"registered_on"`
	Tags         []string  `json:"tags" bson:"tags"`
}

// Method is a supported method for notification.  The only two
// supported methods at this time are application and smtp
type Method struct {
	Type   string `json:"type" bson:"type"`
	Target string `json:"target" bson:"target"`
}
