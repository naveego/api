package types

type User struct {
	ID   string `json:"id" bson:"_id"`    // The ID of the user
	Name string `json:"name" bson:"name"` // The Name of the user

}
