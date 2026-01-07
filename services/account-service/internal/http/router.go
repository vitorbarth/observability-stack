package httpapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/vitormbarth/observability-stack/services/account-service/internal/config"
	"github.com/vitormbarth/observability-stack/services/account-service/internal/repository"
	"github.com/vitormbarth/observability-stack/services/account-service/internal/service"
)

func NewRouter(cfg config.Config) *chi.Mux {
	logger := httplog.NewLogger(cfg.ServiceName)

	repo := repository.NewMemoryRepository()
	svc := service.NewAccountService(repo)
	h := NewHandler(cfg, svc)

	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.AllowAll().Handler)
	r.Use(otelhttp.NewMiddleware(cfg.ServiceName))

	r.Route("/accounts", func(rt chi.Router) {
		rt.Get("/", h.GetAll)
		rt.Post("/", h.Create)
		rt.Get("/{id}", h.Get)
		rt.Delete("/{id}", h.Delete)
	})

	return r
}
