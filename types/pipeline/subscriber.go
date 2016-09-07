package pipeline

type Subscriber struct {
	ID          string `json:"id"`                    // The ID
	Name        string `json:"name"`                  // The Name
	Description string `json:"description,omitempty"` // The Description
	IconURL     string `json:"icon"`
}
