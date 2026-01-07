package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vitormbarth/observability-stack/services/auth-service/internal/config"
	"github.com/vitormbarth/observability-stack/services/auth-service/internal/http"
	"github.com/vitormbarth/observability-stack/services/shared/otel"
)

func main() {
	cfg := config.Load()

	// Init OpenTelemetry
	shutdown, err := otel.InitProvider(otel.Config{
		ServiceName: cfg.ServiceName,
		OtlpEndpoint: cfg.OtelEndpoint,
	})
	if err != nil {
		log.Fatalf("failed to init otel: %v", err)
	}
	defer shutdown()

	r := httpapi.NewRouter(cfg)

	log.Printf("Auth service running on %s", cfg.HttpAddr)
	if err := http.ListenAndServe(cfg.HttpAddr, r); err != nil {
		log.Fatal(err)
	}
}
