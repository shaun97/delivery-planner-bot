package entity

import "time"

type Coordinates struct {
    Latitude  float64
    Longitude float64
}

type TimeWindow struct {
    Start time.Time
    End   time.Time
}
