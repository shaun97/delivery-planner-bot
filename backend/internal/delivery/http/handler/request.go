package handler

type CreateRouteRequest struct {
	DriverID    string             `json:"driverId" binding:"required"`
	Origin      CoordinatesDTO     `json:"origin" binding:"required"`
	Destination CoordinatesDTO     `json:"destination" binding:"required"`
	Deliveries  []DeliveryPointDTO `json:"deliveries" binding:"required"`
}

type UpdateRouteRequest struct {
	DriverID    string             `json:"driverId,omitempty"`
	Origin      *CoordinatesDTO    `json:"origin,omitempty"`
	Destination *CoordinatesDTO    `json:"destination,omitempty"`
	Deliveries  []DeliveryPointDTO `json:"deliveries,omitempty"`
	Status      *string            `json:"status,omitempty"`
}

// Add this type at the top of the file with other type declarations
type PreviewRouteRequest struct {
	Origin      string   `json:"origin" binding:"required"`
	Destination string   `json:"destination" binding:"required"`
	Deliveries  []string `json:"deliveries" binding:"required"`
}
