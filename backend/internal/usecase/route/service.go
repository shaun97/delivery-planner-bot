package route

import (
	"context"
	"fmt"

	"delivery-planner-bot/backend/internal/domain/entity"
	"delivery-planner-bot/backend/internal/domain/repository"
)

type service struct {
	routeRepo  repository.RouteRepository
	mapService MapService
}

func NewService(routeRepo repository.RouteRepository, mapService MapService) UseCase {
	return &service{
		routeRepo:  routeRepo,
		mapService: mapService,
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

func (s *service) OptimizeRoute(ctx context.Context, id string) error {
	route, err := s.routeRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find route: %w", err)
	}

	optimizedDeliveries, err := s.mapService.OptimizeDeliverySequence(ctx, route.Deliveries)
	if err != nil {
		return fmt.Errorf("failed to optimize delivery sequence: %w", err)
	}

	route.Deliveries = optimizedDeliveries

	// Calculate ETA between each point
	var totalETA float64
	for i := 0; i < len(route.Deliveries)-1; i++ {
		eta, err := s.mapService.CalculateETA(ctx,
			&route.Deliveries[i].Coordinates,
			&route.Deliveries[i+1].Coordinates)
		if err != nil {
			return fmt.Errorf("failed to calculate ETA: %w", err)
		}
		totalETA += eta
	}

	return s.routeRepo.Update(ctx, route)
}
