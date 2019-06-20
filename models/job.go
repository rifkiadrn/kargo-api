package models

import (
	"time"
)

// Job struct represents job object of shippers
type Job struct {
	ID           int64     `json:"id"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	ShipmentDate time.Time `json:"shipment_date"`
	Description  string    `json:"description"`
	Budget       float64   `json:"budget"`
	ShipperID    int64     `json:"shipper_id"`
}
