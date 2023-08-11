package presentations

type ErrorValidation struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}
