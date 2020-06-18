package main

import (
	"context"
	"distributed-tracing-otel/tracing"
	"distributed-tracing-otel/weatherpb"
	"log"
	"net"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/plugin/grpctrace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	locations map[string]string
}

func (s *server) GetCurrentWeather(ctx context.Context, in *weatherpb.WeatherRequest) (*weatherpb.WeatherResponse, error) {
	span := trace.SpanFromContext(ctx)
	l, ok := s.locations[in.Location]
	if !ok {
		err := status.Error(codes.NotFound, "Location not found")
		span.RecordError(ctx, err, trace.WithErrorStatus(codes.NotFound))
		return nil, err
	}

	span.AddEvent(ctx, "Selected condition",
		kv.String("condition", l),
		kv.String("location", in.Location),
	)

	t, err := getTemperature(ctx)

	if err != nil {
		err := status.Error(codes.Unknown, err.Error())
		span.RecordError(ctx, err)
		return nil, err
	}

	span.AddEvent(ctx, "Temperature received",
		kv.Float64("temperature", t),
	)

	return &weatherpb.WeatherResponse{
		Condition:   l,
		Temperature: t,
	}, nil
}

func main() {
	fn := tracing.InitTraceProvider("weather")
	defer fn()
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := &server{
		locations: map[string]string{
			"dublin":   "rainy",
			"galway":   "sunny",
			"limerick": "cloudy",
		},
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor(global.Tracer("weather"))))
	weatherpb.RegisterWeatherServiceServer(s, server)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
