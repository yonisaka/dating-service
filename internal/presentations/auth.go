package presentations

type AuthInfo struct {
	Authorized        bool  `json:"authorized"`
	Dtid              int   `json:"dtid"`
	ExpiryTime        int   `json:"expiry_time"`
	RefreshExpiryTime int   `json:"refresh_expiry_time"`
	UserID            int64 `json:"user_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    string `json:"expired_at"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
	Intend    string `json:"intend"`
}

type RegisterResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
}
