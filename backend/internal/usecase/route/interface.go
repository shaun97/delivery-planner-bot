package route

import (
	"context"
	"time"

	"delivery-planner-bot/backend/internal/domain/entity"
)

type UseCase interface {
	CreateRoute(ctx context.Context, route *entity.Route) error
	GetRoute(ctx context.Context, id string) (*entity.Route, error)
	UpdateRoute(ctx context.Context, route *entity.Route) error
	DeleteRoute(ctx context.Context, id string) error
	ListRoutes(ctx context.Context) ([]*entity.Route, error)
	OptimizeRoute(ctx context.Context, id string) error
	GetDriverRoutes(ctx context.Context, driverID string) ([]*entity.Route, error)

	// PreviewRoute is a method to check if the route is valid before creating it
	PreviewRoute(ctx context.Context, origin, destination string, deliveries []string) (*entity.Route, error)
}

type MapService interface {
	CalculateETA(ctx context.Context, origin, destination *entity.Coordinates) (time.Duration, error)
	OptimizeDeliverySequence(ctx context.Context, origin, destination *entity.Coordinates, deliveries []*entity.DeliveryPoint) (*entity.Route, error)
	BuildMapURL(origin, destination string, deliveries []*entity.DeliveryPoint) string
}

type GeocodingService interface {
	GetCoordinates(ctx context.Context, address string) (*entity.Coordinates, string, error)
}
