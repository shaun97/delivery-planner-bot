package models

type Delivery struct {
	ID       string
	Address  string
	Priority int
}

type Route struct {
	ID        string
	Deliveries []Delivery
}
