package main

import (
	"log"
	"net/http"

	"github.com/vitormbarth/observability-stack/services/account-service/internal/config"
	httpapi "github.com/vitormbarth/observability-stack/services/account-service/internal/http"
	"github.com/vitormbarth/observability-stack/services/shared/otel"
)

func main() {
	cfg := config.Load()

	shutdown, err := otel.InitProvider(otel.Config{
		ServiceName:  cfg.ServiceName,
		OtlpEndpoint: cfg.OtelEndpoint,
	})
	if err != nil {
		log.Fatalf("failed to init otel: %v", err)
	}
	defer shutdown()

	r := httpapi.NewRouter(cfg)

	log.Printf("Account service running on %s", cfg.HttpAddr)
	if err := http.ListenAndServe(cfg.HttpAddr, r); err != nil {
		log.Fatal(err)
	}
}
