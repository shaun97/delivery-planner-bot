package maps

import (
	"context"
	"fmt"
	"regexp"

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
		Components: map[maps.Component]string{
			maps.ComponentCountry: "SG", // Restrict to Singapore
		},
	}

	results, err := g.client.Geocode(ctx, r)
	if err != nil {
		return nil, "", fmt.Errorf("geocoding failed: %w", err)
	}

	if len(results) == 0 {
		return nil, "", fmt.Errorf("no coordinates found for address: %s", address)
	}
	readableAddress := results[0].FormattedAddress

	validAddressPattern := `^Singapore [0-9]{6}$`
	matched, _ := regexp.MatchString(validAddressPattern, readableAddress)
	if matched {
		for _, component := range results[0].AddressComponents {
			if component.Types[0] != "postal_code" {
				readableAddress = fmt.Sprintf("%s, %s", component.LongName, readableAddress)
				break
			}
		}
	}

	return &entity.Coordinates{
		Latitude:  results[0].Geometry.Location.Lat,
		Longitude: results[0].Geometry.Location.Lng,
	}, readableAddress, nil
}
