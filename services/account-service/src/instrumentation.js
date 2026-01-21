import * as opentelemetry from '@opentelemetry/sdk-node';
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-proto';
import { OTLPMetricExporter } from '@opentelemetry/exporter-metrics-otlp-proto';
import { PeriodicExportingMetricReader } from '@opentelemetry/sdk-metrics';
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin';
import dotenv from "dotenv";
dotenv.config();

// Configure the Zipkin exporter
const zipkinExporter = new ZipkinExporter({
  url: "http://zipkin:9411/api/v2/spans",
  serviceName: process.env.OTEL_SERVICE_NAME,
});

const sdk = new opentelemetry.NodeSDK({
  traceExporter: zipkinExporter,
  instrumentations: [getNodeAutoInstrumentations()],
});
sdk.start();