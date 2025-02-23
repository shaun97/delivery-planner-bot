package maps

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"delivery-planner-bot/backend/internal/domain/entity"
	"delivery-planner-bot/backend/internal/usecase/route"

	routing "cloud.google.com/go/maps/routing/apiv2"
	"cloud.google.com/go/maps/routing/apiv2/routingpb"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc/metadata"
)

type googleRoutesClient struct {
	client *routing.RoutesClient
}

func NewGoogleRoutesClient(ctx context.Context, apiKey string) (route.MapService, error) {
	client, err := routing.NewRoutesClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Maps client: %w", err)
	}

	return &googleRoutesClient{
		client: client,
	}, nil
}

func (g *googleRoutesClient) CalculateETA(ctx context.Context, origin, destination *entity.Coordinates) (time.Duration, error) {
	req := &routingpb.ComputeRoutesRequest{
		Origin: &routingpb.Waypoint{
			LocationType: &routingpb.Waypoint_Location{
				Location: &routingpb.Location{
					LatLng: &latlng.LatLng{
						Latitude:  origin.Latitude,
						Longitude: origin.Longitude,
					},
				},
			},
		},
		Destination: &routingpb.Waypoint{
			LocationType: &routingpb.Waypoint_Location{
				Location: &routingpb.Location{
					LatLng: &latlng.LatLng{
						Latitude:  destination.Latitude,
						Longitude: destination.Longitude,
					},
				},
			},
		},
		TravelMode:            routingpb.RouteTravelMode_DRIVE,
		RoutingPreference:     routingpb.RoutingPreference_TRAFFIC_AWARE,
		RouteModifiers:        &routingpb.RouteModifiers{},
		OptimizeWaypointOrder: false,
	}

	// set the field mask
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", "*")

	resp, err := g.client.ComputeRoutes(ctx, req)

	if err != nil {
		return 0, fmt.Errorf("failed to compute route: %w", err)
	}

	if len(resp.Routes) == 0 {
		return 0, fmt.Errorf("no route found")
	}

	// Duration is in seconds
	return resp.Routes[0].Duration.AsDuration(), nil
}

func (g *googleRoutesClient) OptimizeDeliverySequence(ctx context.Context, origin, destination *entity.Coordinates, deliveries []*entity.DeliveryPoint) (*entity.Route, error) {
	deliveriesWaypoints := []*routingpb.Waypoint{}
	for _, delivery := range deliveries {
		deliveriesWaypoints = append(deliveriesWaypoints, &routingpb.Waypoint{
			LocationType: &routingpb.Waypoint_Location{
				Location: &routingpb.Location{
					LatLng: &latlng.LatLng{
						Latitude:  delivery.Coordinates.Latitude,
						Longitude: delivery.Coordinates.Longitude,
					},
				},
			},
			VehicleStopover: true,
		})
	}

	req := &routingpb.ComputeRoutesRequest{
		Origin: &routingpb.Waypoint{
			LocationType: &routingpb.Waypoint_Location{
				Location: &routingpb.Location{
					LatLng: &latlng.LatLng{
						Latitude:  origin.Latitude,
						Longitude: origin.Longitude,
					},
				},
			},
		},
		Destination: &routingpb.Waypoint{
			LocationType: &routingpb.Waypoint_Location{
				Location: &routingpb.Location{
					LatLng: &latlng.LatLng{
						Latitude:  destination.Latitude,
						Longitude: destination.Longitude,
					},
				},
			},
		},
		Intermediates:         deliveriesWaypoints,
		TravelMode:            routingpb.RouteTravelMode_DRIVE,
		RoutingPreference:     routingpb.RoutingPreference_TRAFFIC_AWARE,
		RouteModifiers:        &routingpb.RouteModifiers{},
		OptimizeWaypointOrder: true,
	}

	// set the field mask
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", "*")

	resp, err := g.client.ComputeRoutes(ctx, req)

	respRoute := resp.Routes[0]

	if err != nil {
		return nil, fmt.Errorf("failed to compute route: %w", err)
	}

	if len(resp.Routes) == 0 {
		return nil, fmt.Errorf("no route found")
	}
	optimizedRoute := make([]*entity.DeliveryPoint, len(deliveries))
	for i, idx := range respRoute.OptimizedIntermediateWaypointIndex {
		if i < 0 {
			optimizedRoute[i] = deliveries[i]
		} else {
			optimizedRoute[i] = deliveries[idx]
		}

	}

	result := &entity.Route{
		EstimatedTime:   respRoute.Duration.AsDuration(),
		DeliveriesCoord: optimizedRoute,
	}

	return result, nil

}

func (g *googleRoutesClient) BuildMapURL(origin, destination string, deliveries []*entity.DeliveryPoint) string {
	res := fmt.Sprintf("https://www.google.com/maps/dir/?api=1&travelmode=driving&origin=%s&destination=%s&waypoints=", url.QueryEscape(origin), url.QueryEscape(destination))
	for idx, deliveryLoc := range deliveries {
		if idx == (len(deliveries) - 1) {
			res += fmt.Sprintf("%v", url.QueryEscape(deliveryLoc.Address))
		} else {
			res += fmt.Sprintf("%v|", url.QueryEscape(deliveryLoc.Address))
		}

	}
	return res
}
