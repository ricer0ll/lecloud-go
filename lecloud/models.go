package lecloud

type GetJWTRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTResponse struct {
	Token string `json:"token"`
}

type GetSecretResponse struct {
	SecretValue string `json:"secret_value"`
}
