package models

type AddDriverRequest struct {
	DriverID   int     `json:"driver_id" `
	DriverName string  `json:"driver_name" `
	Latitude   float64 `json:"latitude" `
	Longitude  float64 `json:"longitude" `
}

// Passenger represents a driver with basic db details
type Passenger struct {
	PassengerID   int    `json:"passenger_Id"`
	PassengerName string `json:"passenger_name"`
}

// RideReuest represent a macthing request for a passenger
type RideRequest struct {
	PassengerID   int     `json:"passenger_Id"`
	PassengerName string  `json:"passenger_name"`
	Latitude      float64 `json:"lat"`
	Longitude     float64 `json:"lon"`
	Result        chan Driver
}
