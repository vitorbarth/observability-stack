module github.com/vitormbarth/observability-stack/services/auth-service

go 1.25

require (
    github.com/go-chi/chi/v5 v5.2.1
    github.com/go-chi/cors v1.2.1
    github.com/go-chi/httplog v0.2.1
    github.com/golang-jwt/jwt/v5 v5.2.1

    go.opentelemetry.io/otel v1.28.0
    go.opentelemetry.io/otel/sdk v1.28.0
    go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0
)

replace github.com/vitormbarth/observability-stack/services/shared => ../shared