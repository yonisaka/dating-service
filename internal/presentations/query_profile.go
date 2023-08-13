package presentations

type QueryProfileResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Intend    string `json:"intend"`
}
