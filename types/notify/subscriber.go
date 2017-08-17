package notify

// Subscriber is a target for notifications.
type Subscriber struct {
	ID       string   `json:"id" bson:"_id"`
	TenantID string   `json:"tenant_id" bson:"tenant_id"`
	Name     string   `json:"name" bson:"name"`
	Methods  []Method `json:"methods" bson:"methods"`
}

// Method is a supported method for notification.  The only two
// supported methods at this time are application and smtp
type Method struct {
	Type    string `json:"type" bson:"type"`
	Address string `json:"address" bson:"address"`
}
