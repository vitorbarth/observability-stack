package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vitorbarth/observability-stack/services/auth-service/internal/service"
	"github.com/vitorbarth/observability-stack/services/shared/otel"
)

type Handler struct {
	svc *service.AuthService
}

func NewHandler(svc *service.AuthService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/validate-token", h.ValidateToken).Methods("GET")
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.StartSpan(r.Context(), "auth.login")
	defer span.End()

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "payload inválido", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// em ambiente real: validar senha no banco
	token, err := h.svc.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.StartSpan(r.Context(), "auth.validate-token")
	defer span.End()

	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		http.Error(w, "token não informado", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	username, err := h.svc.ValidateToken(token)
	if err != nil {
		http.Error(w, "token inválido", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"valid":    "true",
		"username": username,
	})
}
