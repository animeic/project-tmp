package response

type Auth struct {
	UserInfo UserInfo `json:"user_info"`
	Token    Token    `json:"token"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
