package oauth2

type Token struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IssuedAt     uint64 `json:"issued_at"`
}
