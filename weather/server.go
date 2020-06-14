package main

import (
	"context"
	"log"
	"net"

	"github.com/RicardoLinck/distributed-tracing-otel/weather/weatherpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	locations map[string]string
}

func (s *server) GetCurrentWeather(ctx context.Context, in *weatherpb.WeatherRequest) (*weatherpb.WeatherResponse, error) {
	log.Printf("Sum rpc invoked with req: %v\n", in)

	l, ok := s.locations[in.Location]
	if !ok {
		return nil, status.Error(codes.NotFound, "Location not found")
	}

	return &weatherpb.WeatherResponse{
		Condition: l,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := &server{
		locations: map[string]string{
			"dublin": "rainy",
			"galway": "sunny",
			"limerick":"cloudy"
		}
	}
	weatherpb.RegisterWeatherServer(s, server)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
