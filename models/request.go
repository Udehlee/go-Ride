package models

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	Role      string  `json:"role"` // "driver" or "passenger"
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`  // Distance from a passenger
	Available bool    `json:"available"` // Whether the driver is available
}

type AddDriverRequest struct {
	DriverID  int     `json:"driver_id" binding:"required"`
	FirstName string  `json:"first_name" binding:"required"`
	Role      string  `json:"role" binding:"required"` // "driver"
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

// RideReuest represent a macthing request for a passenger
type RideRequest struct {
	PassengerID   int
	PassengerName string
	Latitude      float64
	Longitude     float64
	Result        chan User
}
