package types

// Principal represents a entity that can interact with the
// Naveego platform.  Principals are used to secure access to
// platform resources.
type Principal struct {
	ID         string   `json:"id"`         // A unique identifier for the principal
	Name       string   `json:"name"`       // The name of the princpal
	Repository string   `json:"repository"` // The repository the principal belongs bool
	Roles      []string `json:"roles"`      // The roles the prinicpal fulfills
}