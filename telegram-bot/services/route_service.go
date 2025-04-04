package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RouteService struct {
	client  *http.Client
	baseURL string
}

type PreviewRouteRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Deliveries  []string `json:"deliveries"`
}

type PreviewRouteResponse struct {
	Origin        string   `json:"origin"`
	Destination   string   `json:"destination"`
	Deliveries    []string `json:"deliveries"`
	EstimatedTime string   `json:"estimatedTime"`
	GoogleMapsURL string   `json:"googleMapsURL"`
}

func NewRouteService(baseURL string) *RouteService {
	return &RouteService{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

func (s *RouteService) PreviewRoute(origin, destination string, deliveries []string) (*PreviewRouteResponse, error) {
	url := fmt.Sprintf("%s/api/v1/routes/preview", s.baseURL)

	reqBody := PreviewRouteRequest{
		Origin:      origin,
		Destination: destination,
		Deliveries:  deliveries,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}

	var previewResp PreviewRouteResponse
	if err := json.NewDecoder(resp.Body).Decode(&previewResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &previewResp, nil
}
