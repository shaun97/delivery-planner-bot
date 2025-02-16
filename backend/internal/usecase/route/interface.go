package route

import (
    "context"

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
}

type MapService interface {
    CalculateETA(ctx context.Context, origin, destination *entity.Coordinates) (float64, error)
    OptimizeDeliverySequence(ctx context.Context, deliveries []*entity.DeliveryPoint) ([]*entity.DeliveryPoint, error)
}
