package presentations

type ProfileResponse struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Status    string   `json:"status"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Dob       string   `json:"dob"`
	Gender    string   `json:"gender"`
	Intend    string   `json:"intend"`
	Images    []string `json:"images"`
}
