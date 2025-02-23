package route

import (
	"context"
	"fmt"

	"delivery-planner-bot/backend/internal/domain/entity"
	"delivery-planner-bot/backend/internal/domain/repository"
)

type service struct {
	routeRepo        repository.RouteRepository
	mapService       MapService
	geocodingService GeocodingService
}

func NewService(routeRepo repository.RouteRepository, mapService MapService, geocodingService GeocodingService) UseCase {
	return &service{
		routeRepo:        routeRepo,
		mapService:       mapService,
		geocodingService: geocodingService,
	}
}

func (s *service) CreateRoute(ctx context.Context, route *entity.Route) error {
	return s.routeRepo.Create(ctx, route)
}

func (s *service) GetRoute(ctx context.Context, id string) (*entity.Route, error) {
	return s.routeRepo.FindByID(ctx, id)
}

func (s *service) UpdateRoute(ctx context.Context, route *entity.Route) error {
	return s.routeRepo.Update(ctx, route)
}

func (s *service) DeleteRoute(ctx context.Context, id string) error {
	return s.routeRepo.Delete(ctx, id)
}

func (s *service) ListRoutes(ctx context.Context) ([]*entity.Route, error) {
	return s.routeRepo.List(ctx)
}

func (s *service) GetDriverRoutes(ctx context.Context, driverID string) ([]*entity.Route, error) {
	return s.routeRepo.FindByDriverID(ctx, driverID)
}

func (s *service) PreCheckRoute(ctx context.Context, origin, destination string, deliveries []string) (*entity.Route, error) {
	route, err := s.createRouteWithGeocodedAddresses(ctx, origin, destination, deliveries)
	if err != nil {
		return nil, err
	}

	err = s.optimizeAndEnrichRoute(ctx, route)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (s *service) createRouteWithGeocodedAddresses(ctx context.Context, origin, destination string, deliveries []string) (*entity.Route, error) {
	originCoord, originAddr, err := s.geocodingService.GetCoordinates(ctx, origin)
	if err != nil {
		return nil, fmt.Errorf("failed to geocode origin address: %w", err)
	}

	destCoord, destAddr, err := s.geocodingService.GetCoordinates(ctx, destination)
	if err != nil {
		return nil, fmt.Errorf("failed to geocode destination address: %w", err)
	}

	deliveriesCoord, err := s.geocodeDeliveries(ctx, deliveries)
	if err != nil {
		return nil, err
	}

	return &entity.Route{
		OriginCoord:      *originCoord,
		DestinationCoord: *destCoord,
		DeliveriesCoord:  deliveriesCoord,
		Origin:           originAddr,
		Destination:      destAddr,
	}, nil
}

func (s *service) geocodeDeliveries(ctx context.Context, deliveries []string) ([]*entity.DeliveryPoint, error) {
	deliveriesCoord := make([]*entity.DeliveryPoint, len(deliveries))
	for idx, delivery := range deliveries {
		coord, addr, err := s.geocodingService.GetCoordinates(ctx, delivery)
		if err != nil {
			return nil, fmt.Errorf("failed to geocode delivery address: %w", err)
		}
		deliveriesCoord[idx] = &entity.DeliveryPoint{
			Coordinates: *coord,
			Address:     addr,
		}
	}
	return deliveriesCoord, nil
}

func (s *service) optimizeAndEnrichRoute(ctx context.Context, route *entity.Route) error {
	resp, err := s.mapService.OptimizeDeliverySequence(ctx, &route.OriginCoord, &route.DestinationCoord, route.DeliveriesCoord)
	if err != nil {
		return fmt.Errorf("failed to optimize delivery sequence: %w", err)
	}

	route.DeliveriesCoord = resp.DeliveriesCoord
	route.EstimatedTime = resp.EstimatedTime
	route.GoogleMapsURL = s.mapService.BuildMapURL(route.Origin, route.Destination, route.DeliveriesCoord)

	return nil
}

// Factor in time windows, vehicle capacity, and other constraints
func (s *service) OptimizeRoute(ctx context.Context, id string) error {
	route, err := s.routeRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find route: %w", err)
	}

	// optimizedDeliveries, err := s.mapService.OptimizeDeliverySequence(ctx, route.DeliveriesCoord)
	// if err != nil {
	// 	return fmt.Errorf("failed to optimize delivery sequence: %w", err)
	// }

	// route.DeliveriesCoord = optimizedDeliveries

	// Calculate ETA between each point
	// var totalETA float64
	// for i := 0; i < len(route.Deliveries)-1; i++ {
	// 	eta, err := s.mapService.CalculateETA(ctx,
	// 		&route.Deliveries[i].Coordinates,
	// 		&route.Deliveries[i+1].Coordinates)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to calculate ETA: %w", err)
	// 	}
	// 	totalETA += eta
	// }

	return s.routeRepo.Update(ctx, route)
}
