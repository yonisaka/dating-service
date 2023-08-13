package presentations

type ActionHistoryResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Action    string `json:"action"`
	UpdatedAt string `json:"updated_at"`
}
