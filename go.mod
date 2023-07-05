module distributed-tracing-otel

go 1.13

require (
	github.com/golang/protobuf v1.5.2
	go.opentelemetry.io/otel v0.6.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.6.0
	google.golang.org/grpc v1.53.0
)
