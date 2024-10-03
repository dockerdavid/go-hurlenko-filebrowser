package auth

type Credentials struct {
	BaseURL   string `json:"baseURL"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Recaptcha string `json:"recaptcha"`
}
