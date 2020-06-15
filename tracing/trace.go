package tracing

import (
	"log"

	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// GRPCTracer contains the tracer for UnaryServerInterceptor
type GRPCTracer struct {
	Tracer trace.Tracer
}

// InitTraceProvider initiates the tracing provider for a given service
func InitTraceProvider(service string) func() {
	// Create and install Jaeger export pipeline
	_, flush, err := jaeger.NewExportPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: service,
			Tags: []kv.KeyValue{
				kv.Key("exporter").String("jaeger"),
			},
		}),
		jaeger.RegisterAsGlobal(),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}
