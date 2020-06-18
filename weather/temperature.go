package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.opentelemetry.io/otel/plugin/httptrace"
)

// WeatherForecast represents the response from temperature service
type WeatherForecast struct {
	TemperatureC int `json:"temperatureC"`
}

func getTemperature(ctx context.Context) (float64, error) {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", "http://localhost:5000/WeatherForecast", nil)
	ctx, req = httptrace.W3C(ctx, req)
	httptrace.Inject(ctx, req)
	res, err := client.Do(req)
	if err != nil {
		return 0.0, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()
	wf, err := parseTemperatureResponse(body)
	return float64(wf.TemperatureC), err
}

func parseTemperatureResponse(body []byte) (WeatherForecast, error) {
	wf := &WeatherForecast{}
	err := json.Unmarshal(body, wf)
	return *wf, err
}
