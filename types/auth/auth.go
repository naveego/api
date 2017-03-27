package auth

type Configuration struct {
	ID                  string `json:"id" bson:"_id"`
	Name                string `json:"name" bson:"name"`
	Provider            string `json:"provider" bson:"provider"`
	ProviderSettingsURL string `json:"provider_settings" bson:"provider_settings"`
}

type UserStatus string

const (
	UserStatusSeen                UserStatus = "SEEN"
	UserStatusRequesting          UserStatus = "REQUESTING"
	UserStatusUnconfirmed         UserStatus = "UNCONFIRMED"
	UserStatusConfirmed           UserStatus = "CONFIRMED"
	UserStatusArchived            UserStatus = "ARCHIVED"
	UserStatusComprimised         UserStatus = "COMPRIMISED"
	UserStatusForceChangePassword UserStatus = "FORCECHANGEPASSWORD"
	UserStatusResetRequired       UserStatus = "RESETREQUIRED"
)

type UserAttributes map[string]interface{}

type User struct {
	ID            string         `json:"id" bson:"_id"`
	Username      string         `json:"username" bson:"username"`
	TenantID      string         `json:"tenant_id" bson:"tenant_id"`
	Status        UserStatus     `json:"status" bson:"status"`
	Attributes    UserAttributes `json:"attributes" bson:"attributes"`
	CreatedOn     string         `json:"created_on" bson:"created_on"`
	ModifiedOn    string         `json:"modified_on" bson:"modified_on"`
	LastTouchedOn string         `json:"last_touch" bson:"last_touch"`
}
