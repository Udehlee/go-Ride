package models

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_Id"`
	Email     string    `json:"email" `
	Password  string    `json:"pass_word"`
	Role      string    `json:"user_role"` // 'passenger' or 'driver'
	CreatedAt time.Time `json:"created_at"`
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

// Driver represents a driver with location(latitude, longitude)
type Driver struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}
