package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/vitormbarth/observability-stack/services/auth-service/internal/config"
	"github.com/vitormbarth/observability-stack/services/auth-service/internal/service"
)

type Handler struct {
	cfg    config.Config
	auth   *service.AuthService
}

func NewHandler(cfg config.Config, auth *service.AuthService) *Handler {
	return &Handler{cfg, auth}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	token, err := h.auth.Login(req.User, req.Pass)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
