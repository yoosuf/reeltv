package domain

// AuthService defines domain-level authentication business logic
type AuthService interface {
	ValidateCredentials(email, phone, password string) (interface{}, error)
	GenerateTokens(userID uint, role string) (accessToken, refreshToken string, err error)
	ValidateAccessToken(token string) (userID uint, role string, err error)
	ValidateRefreshToken(token string) (userID uint, err error)
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

// Claims represents JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}

// TokenPair represents access and refresh token pair
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64 // seconds
}
