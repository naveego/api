package types

import "time"

type User struct {
	ID            string    `json:"id" bson:"_id"`                       // The ID of the user
	DisplayName   string    `json:"display_name" bson:"displayname"`     // The Name of the user
	FirstName     string    `json:"first_name" bson:"firstname"`         // The First Name of the user
	LastName      string    `json:"last_name" bson:"lastname"`           // The Last Nazme of the user
	EmailAddress  string    `json:"email_address" bson:"email"`          // The email address of the user
	EmailVerified bool      `json:"email_verified" bson:"emailverified"` // Whether or not the email has been verified
	Flags         []string  `json:"flags,omitempty" bson:"flags"`
	LastLogin     time.Time `json:"last_login,omitempty" bson:"lastlogin"` // THe last time the user logged in
	Roles         []string  `json:"roles,omitempty" bson:"roles"`          // The roles the user belongs too
}
