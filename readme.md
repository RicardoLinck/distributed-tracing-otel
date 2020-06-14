# Distributed Tracing with OpenTelemetry

This sample application shows how to implement distributed trancing using OpenTelemetry in golang and asp.net core.

## Architecure

### client

This is the entrypoint of the sample. This is a simple golang app calling the weather-service.

### weather-service

Golang web server serving returning the weather description and temperature (obtained from the temperature-service) via gRPC. In case the location is not found, it returns an error instead.

### temperature-service

Asp.net core web server returning random temperatures via HTTP.

## Running

This sample uses Jaeger as the backend for distributed tracing. To start Jaeger run the following command:

```bash
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.18
```
