package main

import (
	"context"
	"log"

	"distributed-tracing-otel/weatherpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()

	c := weatherpb.NewWeatherServiceClient(cc)
	sum(c)
}

func sum(c weatherpb.WeatherServiceClient) {
	req := &weatherpb.WeatherRequest{
		Location: "Dublin",
	}
	res, err := c.GetCurrentWeather(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum RPC: %v", err)
	}
	log.Printf("Response: %d\n", res.Condition)
}
