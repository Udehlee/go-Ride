package db

import (
	"database/sql"
	"fmt"

	"github.com/Udehlee/go-Ride/models"
)

type Repo interface {
	SaveDriver(driver models.Driver) error
	SavePassenger(passenger models.Passenger) error
	SaveMatchedRides(matchedRide models.MatchedRide) error
}
type PgConn struct {
	Conn *sql.DB
}

func NewPgConn(db *sql.DB) PgConn {
	return PgConn{
		Conn: db,
	}
}

func (p *PgConn) SaveDriver(driver models.Driver) error {
	query := "INSERT INTO drivers (driver_name, latitude, longitude) VALUES ($1,$2,$3) RETURNING driver_id, driver_name"

	row := p.Conn.QueryRow(query, driver.DriverName, driver.Latitude, driver.Longitude)
	if err := row.Scan(&driver.DriverID, &driver.DriverName); err != nil {
		return fmt.Errorf("failed to scan users row: %w", err)
	}

	return nil
}

func (p *PgConn) SavePassenger(passenger models.Passenger) error {
	query := "INSERT INTO passengers (passenger_name) VALUES ($1) RETURNING passenger_id, passenger_name"

	row := p.Conn.QueryRow(query, passenger.PassengerName)
	if err := row.Scan(&passenger.PassengerID, &passenger.PassengerName); err != nil {
		return fmt.Errorf("failed to scan users row: %w", err)
	}

	return nil
}

func (pg *PgConn) SaveMatchedRides(matchedRide models.MatchedRide) error {
	query := "INSERT INTO matched_rides (driver_id, passenger_id, ride_status, created_at)	VALUES ($1, $2, $3, $4)"

	_, err := pg.Conn.Exec(query, matchedRide.DriverID, matchedRide.PassengerID, matchedRide.RideStatus, matchedRide.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save matched ride: %w", err)
	}
	return nil
}
