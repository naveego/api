package tenant

// Tenant represents a Naveego customer
type Tenant struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedOn string `json:"createdOn"`
}
