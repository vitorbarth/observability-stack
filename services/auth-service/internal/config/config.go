package config

import "os"

type Config struct {
	ServiceName string
	HttpAddr    string
	JWTSecret   string
	OtelEndpoint string
}

func Load() Config {
	return Config{
		ServiceName: getEnv("SERVICE_NAME", "auth-service"),
		HttpAddr:    getEnv("HTTP_ADDR", ":8080"),
		JWTSecret:   getEnv("JWT_SECRET", "super-secret-local"),
		OtelEndpoint: getEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317"),
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
