package entity

type DeliveryPoint struct {
	ID          string
	Address     string
	Coordinates Coordinates
	TimeWindow  TimeWindow
}
