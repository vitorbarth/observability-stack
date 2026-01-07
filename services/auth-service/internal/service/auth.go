package service

import (
	"errors"

	"github.com/vitormbarth/observability-stack/services/auth-service/internal/token"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService struct {
	tokens *token.Manager
}

func NewAuthService(m *token.Manager) *AuthService {
	return &AuthService{m}
}

func (a *AuthService) Login(username, password string) (string, error) {
	if username == "admin" && password == "admin" {
		return a.tokens.Generate(username)
	}
	return "", ErrInvalidCredentials
}
