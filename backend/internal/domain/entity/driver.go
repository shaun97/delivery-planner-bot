package entity

type Driver struct {
    ID        string
    Name      string
    Status    DriverStatus
    CurrentRoute *Route
}

type DriverStatus string
