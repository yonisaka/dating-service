package presentations

type UserPreferenceRequest struct {
	MinAge    int64 `json:"min_age"`
	MaxAge    int64 `json:"max_age"`
	UseIntend bool  `json:"use_intend"`
}
