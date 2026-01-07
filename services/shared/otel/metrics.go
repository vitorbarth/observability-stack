package otel

import (
    "context"
    "log"

    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
    "go.opentelemetry.io/otel/sdk/metric"
    "google.golang.org/grpc"
)

func InitMeter(ctx context.Context) (*metric.MeterProvider, error) {
    conn, err := grpc.DialContext(ctx, "localhost:4317",
        grpc.WithInsecure(),
        grpc.WithBlock(),
    )
    if err != nil {
        return nil, err
    }

    exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
    if err != nil {
        return nil, err
    }

    mp := metric.NewMeterProvider(
        metric.WithReader(metric.NewPeriodicReader(exporter)),
    )

    log.Println("Metrics provider initialized")
    return mp, nil
}
