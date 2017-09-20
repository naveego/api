package notify

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Subscription_Validate(t *testing.T) {
	tests := []struct {
		label         string
		subscription  Subscription
		expectedError string
	}{
		{
			"with no tenant_id",
			Subscription{},
			"missing tenant_id",
		},
		{
			"with no topic",
			Subscription{TenantID: "vandelay"},
			"missing topic",
		},
		{
			"with no label",
			Subscription{TenantID: "vandelay", Topic: "my:topic"},
			"missing label",
		},
		{
			"with no methods",
			Subscription{TenantID: "vanelday", Topic: "my:topic", Label: "Test"},
			"must have at least one method",
		},
		{
			"with invalid method",
			Subscription{
				TenantID: "vandelay",
				Topic:    "my:topic",
				Label:    "Test",
				Methods: []Method{
					Method{"email", "test@test.com"},
					Method{"notsupported", "test"},
				},
			},
			"'notsupported' is not a valid method type",
		},
		{
			"with blank method target",
			Subscription{
				TenantID: "vandelay",
				Topic:    "my:topic",
				Label:    "Test",
				Methods: []Method{
					Method{"email", ""},
				},
			},
			"one or more methods is missing a target",
		},
	}

	for _, test := range tests {
		Convey(fmt.Sprintf("Given a log %s", test.label), t, func() {
			err := test.subscription.Validate()
			Convey("Should return error with expected message", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, test.expectedError)
			})
		})
	}
}
