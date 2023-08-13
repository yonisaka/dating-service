package presentations

type QueryProfileRequest struct {
	Action string `json:"action" url:"action" `
}

type QueryProfileResponse struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Intend    string   `json:"intend"`
	Images    []string `json:"images"`
}
