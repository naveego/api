package notify

import (
	"fmt"
	"time"

	"github.com/naveego/errors"
)

// Subscription represents a notification registrtion for a
// subscriber.  It is used to connect subscribers to subscribed
// topics.
type Subscription struct {
	ID           string    `json:"id" bson:"_id"`
	TenantID     string    `json:"tenant_id" bson:"tenant_id"`
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

// Validate checks a subscription to make sure it is valid.
func (sub *Subscription) Validate() error {

	if sub.TenantID == "" {
		return errors.NewWithCode(40000001, "missing tenant_id")
	}

	if sub.Topic == "" {
		return errors.NewWithCode(40000002, "missing topic")
	}

	if sub.Label == "" {
		return errors.NewWithCode(40000003, "missing label")
	}

	if len(sub.Methods) == 0 {
		return errors.NewWithCode(40000004, "must have at least one method")
	}

	for _, m := range sub.Methods {
		if m.Type != "email" {
			return errors.NewWithCode(4000005, fmt.Sprintf("'%s' is not a valid method type", m.Type))
		}

		if m.Target == "" {
			return errors.NewWithCode(4000006, "one or more methods is missing a target")
		}
	}

	return nil
}
