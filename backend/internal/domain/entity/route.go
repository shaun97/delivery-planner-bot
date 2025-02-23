package entity

import "time"

type Route struct {
	ID               string
	DriverID         string
	Origin           string
	Destination      string
	Deliveries       []string
	OriginCoord      Coordinates
	DestinationCoord Coordinates
	DeliveriesCoord  []*DeliveryPoint
	Status           RouteStatus
	EstimatedTime    time.Duration
	GoogleMapsURL    string
}

type RouteStatus string
