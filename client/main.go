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
	getCurrentWeather(c)
}

func getCurrentWeather(c weatherpb.WeatherServiceClient) {
	req := &weatherpb.WeatherRequest{
		Location: "dublin",
	}
	res, err := c.GetCurrentWeather(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum RPC: %v", err)
	}
	log.Printf("Response: %s\n", res.Condition)
}
