package credential

type CredentialResponse struct {
	ID        uint   `json:"-"`
	UserID    uint   `json:"user_id"`
	Secret    string `json:"secret"`
	Status    string `json:"status"`
	ExpiredAt string `json:"expired_at"`
}
