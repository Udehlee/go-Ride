package models

import (
	"time"
)

// Driver represents a driver with coordinates(latitude, longitude)
type Driver struct {
	DriverID   int     `json:"id"`
	DriverName string  `json:"driver_name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Distance   float64 `json:"distance"`
	Available  bool    `json:"available"`
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

type MatchedRide struct {
	MatchedRideID int       `json:"id"`
	DriverID      int       `json:"driverId"`
	PassengerID   int       `json:"passengerId"`
	RideStatus    string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
