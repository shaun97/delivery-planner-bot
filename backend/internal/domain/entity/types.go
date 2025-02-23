package entity

import "time"

type Coordinates struct {
	Latitude         float64 `json:"latitude" binding:"required"`
	Longitude        float64 `json:"longitude" binding:"required"`
}
type TimeWindow struct {
	Start time.Time
	End   time.Time
}
