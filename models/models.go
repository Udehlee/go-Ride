package models

type Driver struct {
	DriverID int    `json:"driver_id"`
	Email    string `json:"email"`
	Password string `json:"pass_word"`
	Ratings  int    `json:"ratings"`
}

type Passenger struct {
	PassengerID int    `json:"passenger_id"`
	Email       string `json:"email"`
	Password    string `json:"pass_word"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OSMResponse struct {
	Address struct {
		HouseNumber string `json:"house_number"`
		Road        string `json:"road"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
	} `json:"address"`
}
