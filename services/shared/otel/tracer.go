package otel

import (
    "context"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "google.golang.org/grpc"
)

func InitTracer(ctx context.Context, service string) (*sdktrace.TracerProvider, error) {
    conn, err := grpc.DialContext(ctx, "localhost:4317",
        grpc.WithInsecure(),
        grpc.WithBlock(),
    )
    if err != nil {
        return nil, err
    }

    exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
    if err != nil {
        return nil, err
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.Default()),
    )

    otel.SetTracerProvider(tp)

    log.Println("Tracer initialized for", service)
    return tp, nil
}
