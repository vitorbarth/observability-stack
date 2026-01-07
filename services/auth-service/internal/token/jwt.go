package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret string
}

func NewManager(secret string) *Manager {
	return &Manager{secret}
}

func (m *Manager) Generate(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *Manager) Validate(t string) (string, error) {
	token, err := jwt.Parse(t, func(tok *jwt.Token) (interface{}, error) {
		return []byte(m.secret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrTokenInvalidClaims
	}

	return claims["sub"].(string), nil
}
