package maps

import (
	"context"
	"fmt"

	"delivery-planner-bot/backend/internal/domain/entity"
	"delivery-planner-bot/backend/internal/usecase/route"

	"googlemaps.github.io/maps"
)

type googleGeocodingClient struct {
	client *maps.Client
}

func NewGoogleGeocodingClient(ctx context.Context, apiKey string) (route.GeocodingService, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Maps client: %w", err)
	}

	return &googleGeocodingClient{
		client: client,
	}, nil
}

func (g *googleGeocodingClient) GetCoordinates(ctx context.Context, address string) (*entity.Coordinates, string, error) {
	if address == "" {
		return nil, "", fmt.Errorf("address cannot be empty")
	}

	r := &maps.GeocodingRequest{
		Address: address,
	}

	results, err := g.client.Geocode(ctx, r)
	if err != nil {
		return nil, "", fmt.Errorf("geocoding failed: %w", err)
	}

	if len(results) == 0 {
		return nil, "", fmt.Errorf("no coordinates found for address: %s", address)
	}

	return &entity.Coordinates{
		Latitude:  results[0].Geometry.Location.Lat,
		Longitude: results[0].Geometry.Location.Lng,
	}, results[0].FormattedAddress, nil
}
