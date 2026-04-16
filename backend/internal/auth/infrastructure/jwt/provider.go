package jwt

import (
	"errors"
	"time"

	"reeltv/backend/pkg/password"
	"github.com/golang-jwt/jwt/v5"
)

// Provider implements JWT token generation and validation
type Provider struct {
	secret            []byte
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

// Claims represents JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// NewProvider creates a new JWT provider
func NewProvider(secret string, accessExpiration, refreshExpiration time.Duration) *Provider {
	return &Provider{
		secret:            []byte(secret),
		accessExpiration:  accessExpiration,
		refreshExpiration: refreshExpiration,
	}
}

// GenerateAccessToken generates an access token
func (p *Provider) GenerateAccessToken(userID uint, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.accessExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.secret)
}

// GenerateRefreshToken generates a refresh token
func (p *Provider) GenerateRefreshToken(userID uint, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.refreshExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.secret)
}

// GenerateTokens generates both access and refresh tokens
func (p *Provider) GenerateTokens(userID uint, role string) (string, string, error) {
	accessToken, err := p.GenerateAccessToken(userID, role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := p.GenerateRefreshToken(userID, role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ValidateAccessToken validates an access token and returns claims
func (p *Provider) ValidateAccessToken(tokenString string) (uint, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return p.secret, nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, claims.Role, nil
	}

	return 0, "", errors.New("invalid token")
}

// ValidateRefreshToken validates a refresh token and returns user ID
func (p *Provider) ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return p.secret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}

// HashPassword hashes a password using bcrypt
func (p *Provider) HashPassword(pwd string) (string, error) {
	return password.Hash(pwd)
}

// VerifyPassword verifies a password against a hash
func (p *Provider) VerifyPassword(pwd, hash string) bool {
	return password.Verify(pwd, hash)
}

// ValidateCredentials validates user credentials (placeholder - will be implemented with user repository)
func (p *Provider) ValidateCredentials(email, phone, password string) (interface{}, error) {
	// This will be implemented when we have the user repository
	// For now, this is a placeholder
	return nil, errors.New("not implemented")
}
