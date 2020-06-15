package main

import (
	"context"
	"log"

	"distributed-tracing-otel/tracing"
	"distributed-tracing-otel/weatherpb"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/plugin/grpctrace"
	"google.golang.org/grpc"
)

func main() {
	fn := tracing.InitTraceProvider("client")
	defer fn()
	tracer := global.Tracer("client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor(tracer)))

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()

	c := weatherpb.NewWeatherServiceClient(cc)
	getCurrentWeather(c, tracer)
}

func getCurrentWeather(c weatherpb.WeatherServiceClient, tracer trace.Tracer) {
	req := &weatherpb.WeatherRequest{
		Location: "dublin",
	}

	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "GetCurrentWeather")
	defer span.End()
	res, err := c.GetCurrentWeather(ctx, req)
	if err != nil {
		span.RecordError(ctx, err)
		return
	}

	span.AddEvent(ctx, "Response",
		kv.String("condition", res.Condition),
		kv.Float64("temperature", res.Temperature),
	)
}
