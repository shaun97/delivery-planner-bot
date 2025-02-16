package maps

import (
    "context"
    "fmt"

    "delivery-planner-bot/backend/internal/domain/entity"
    "delivery-planner-bot/backend/internal/usecase/route"
    "googlemaps.github.io/maps"
)

type googleMapsClient struct {
    client *maps.Client
}

func NewGoogleMapsClient(apiKey string) (route.MapService, error) {
    client, err := maps.NewClient(maps.WithAPIKey(apiKey))
    if err != nil {
        return nil, fmt.Errorf("failed to create Google Maps client: %w", err)
    }

    return &googleMapsClient{
        client: client,
    }, nil
}

func (g *googleMapsClient) CalculateETA(ctx context.Context, origin, destination *entity.Coordinates) (float64, error) {
    r := &maps.DistanceMatrixRequest{
        Origins:      []string{fmt.Sprintf("%f,%f", origin.Latitude, origin.Longitude)},
        Destinations: []string{fmt.Sprintf("%f,%f", destination.Latitude, destination.Longitude)},
    }

    resp, err := g.client.DistanceMatrix(ctx, r)
    if err != nil {
        return 0, fmt.Errorf("failed to get distance matrix: %w", err)
    }

    if len(resp.Rows) == 0 || len(resp.Rows[0].Elements) == 0 {
        return 0, fmt.Errorf("no route found")
    }

    // Duration in seconds
    return float64(resp.Rows[0].Elements[0].Duration.Seconds()), nil
}

func (g *googleMapsClient) OptimizeDeliverySequence(ctx context.Context, deliveries []*entity.DeliveryPoint) ([]*entity.DeliveryPoint, error) {
    // This is a simple implementation that could be improved with a more sophisticated algorithm
    // Currently just returns the original sequence
    // TODO: Implement proper route optimization using Google Maps Distance Matrix API
    return deliveries, nil
}
