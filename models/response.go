package models

import "time"

type MatchedRide struct {
	MatchedRideID int       `json:"id"`
	DriverID      int       `json:"driverId"`
	PassengerID   int       `json:"passengerId"`
	RideStatus    string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
