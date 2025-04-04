package handler

type ResponseMeta struct {
	Timestamp string `json:"timestamp"`
	RequestID string `json:"requestId,omitempty"`
}

type ResponseWrapper struct {
	Meta    ResponseMeta `json:"meta"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message,omitempty"`
}

type RouteResponse struct {
	ID            string             `json:"id"`
	DriverID      string             `json:"driverId"`
	Origin        CoordinatesDTO     `json:"origin"`
	Destination   CoordinatesDTO     `json:"destination"`
	Deliveries    []DeliveryPointDTO `json:"deliveries"`
	Status        string             `json:"status"`
	EstimatedTime string             `json:"estimatedTime"`
}

type CoordinatesDTO struct {
	Latitude  float64 `json:"latitude" binding:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" binding:"required,gte=-180,lte=180"`
}

type DeliveryPointDTO struct {
	ID          string         `json:"id"`
	Address     string         `json:"address"`
	Coordinates CoordinatesDTO `json:"coordinates"`
	TimeWindow  TimeWindowDTO  `json:"timeWindow,omitempty"`
}

type TimeWindowDTO struct {
	Start string `json:"start" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	End   string `json:"end" binding:"required,datetime=2006-01-02T15:04:05Z07:00,gtfield=Start"`
}

type PreviewRouteResponse struct {
	Origin        string   `json:"origin"`
	Destination   string   `json:"destination"`
	Deliveries    []string `json:"deliveries"`
	EstimatedTime string   `json:"estimatedTime" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	GoogleMapsURL string   `json:"googleMapsURL"`
}
