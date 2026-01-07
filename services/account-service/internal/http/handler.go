package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/vitormbarth/observability-stack/services/account-service/internal/config"
	"github.com/vitormbarth/observability-stack/services/account-service/internal/service"
)

type Handler struct {
	cfg     config.Config
	account *service.AccountService
}

func NewHandler(cfg config.Config, svc *service.AccountService) *Handler {
	return &Handler{cfg, svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	acc, err := h.account.Create(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.account.GetAll())
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	acc, err := h.account.Get(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	json.NewEncoder(w).Encode(acc)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.account.Delete(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.WriteHeader(204)
}
