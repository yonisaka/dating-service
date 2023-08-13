package presentations

type SubscribeResponse struct {
	SubscriptionCode string `json:"subscription_code"`
	ExpiredAt        string `json:"expired_at"`
}
