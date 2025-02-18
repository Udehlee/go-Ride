package db

import (
	"database/sql"
	"fmt"

	"github.com/Udehlee/go-Ride/models"
)

type Repo interface {
	SaveDriver(driver models.Driver) error
	SavePassenger(driver models.Passenger) error
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
	query := "INSERT INTO drivers (email, pass_word) VALUES ($1,$2) RETURNING driver_id, email, ratings"

	row := p.Conn.QueryRow(query, driver.Email, driver.Password)
	if err := row.Scan(&driver.DriverID, &driver.Email, &driver.Ratings); err != nil {
		return fmt.Errorf("failed to scan drivers row: %w", err)
	}

	return nil
}

func (p *PgConn) SavePassenger(passenger models.Passenger) error {
	query := "INSERT INTO passengers (email, pass_word) VALUES ($1,$2) RETURNING passenger_id, email"

	row := p.Conn.QueryRow(query, passenger.Email, passenger.Password)
	if err := row.Scan(&passenger.PassengerID, &passenger.Email); err != nil {
		return fmt.Errorf("failed to scan passengers row: %w", err)
	}

	return nil
}
