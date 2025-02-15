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
