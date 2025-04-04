package repository

import (
	"context"

	"delivery-planner-bot/backend/internal/domain/entity"
)

type RouteRepository interface {
	Create(ctx context.Context, route *entity.Route) error
	FindByID(ctx context.Context, id string) (*entity.Route, error)
	Update(ctx context.Context, route *entity.Route) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*entity.Route, error)
	FindByDriverID(ctx context.Context, driverID string) ([]*entity.Route, error)
}
