package entity

import "time"

type Route struct {
    ID            string
    DriverID      string
    Deliveries    []*DeliveryPoint
    Status        RouteStatus
    EstimatedTime time.Duration
}


type RouteStatus string
