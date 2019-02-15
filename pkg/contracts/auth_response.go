package contracts

type AuthResponse struct {
	AuthToken string `json:"auth_token"`
	TokenType string `json:"token_type"`
}
