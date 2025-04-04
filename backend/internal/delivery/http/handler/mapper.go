package handler

import (
	"delivery-planner-bot/backend/internal/domain/entity"
)

func (h *RouteHandler) ToPreviewRouteResponse(route *entity.Route) *PreviewRouteResponse {
	deliveries := make([]string, len(route.DeliveriesCoord))
	for idx, delivery := range route.DeliveriesCoord {
		deliveries[idx] = delivery.Address
	}
	return &PreviewRouteResponse{
		Origin:        route.Origin,
		Destination:   route.Destination,
		Deliveries:    deliveries, // TODO print the address
		EstimatedTime: route.EstimatedTime.String(),
		GoogleMapsURL: route.GoogleMapsURL,
	}
}
