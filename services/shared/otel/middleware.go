package otel

import (
    "net/http"
    "time"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
)

// MiddlewareHTTP adds tracing to any HTTP handler.
func MiddlewareHTTP(next http.Handler) http.Handler {
    tracer := otel.Tracer("http-middleware")

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, span := tracer.Start(r.Context(), r.Method+" "+r.URL.Path)
        defer span.End()

        start := time.Now()

        // inject trace-context into headers
        r = r.WithContext(ctx)

        // custom response writer to capture status code
        rw := &responseWriter{ResponseWriter: w, status: 200}

        next.ServeHTTP(rw, r)

        duration := time.Since(start)

        span.SetAttributes(
            attribute.String("http.method", r.Method),
            attribute.String("http.path", r.URL.Path),
            attribute.Int("http.status_code", rw.status),
            attribute.String("http.client_ip", r.RemoteAddr),
            attribute.Float64("http.duration_ms", float64(duration.Milliseconds())),
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func (w *responseWriter) WriteHeader(code int) {
    w.status = code
    w.ResponseWriter.WriteHeader(code)
}
