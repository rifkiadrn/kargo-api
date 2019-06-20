package models

// Bid struct represents bidding object of transporters
type Bid struct {
	ID            int64   `json:"id"`
	Price         float64 `json:"price"`
	Vehicle       string  `json:"vehicle"`
	Description   string  `json:"description"`
	JobID         int64   `json:"job_id"`
	TransporterID int64   `json:"transporter_id"`
}
