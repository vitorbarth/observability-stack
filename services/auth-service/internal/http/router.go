package httpapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/vitormbarth/observability-stack/services/auth-service/internal/config"
	"github.com/vitormbarth/observability-stack/services/auth-service/internal/service"
	"github.com/vitormbarth/observability-stack/services/auth-service/internal/token"
	otelmw "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter(cfg config.Config) *chi.Mux {

	logger := httplog.NewLogger(cfg.ServiceName)

	tm := token.NewManager(cfg.JWTSecret)
	authSvc := service.NewAuthService(tm)

	h := NewHandler(cfg, authSvc)

	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.AllowAll().Handler)
	r.Use(otelmw.NewMiddleware(cfg.ServiceName))

	r.Post("/login", h.Login)

	return r
}
