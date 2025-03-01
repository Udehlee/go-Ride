package db

import (
	"database/sql"
	"fmt"

	"github.com/Udehlee/go-Ride/models"
)

type Repo interface {
	SaveUser(user models.User) error
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

func (p *PgConn) SaveUser(user models.User) error {
	query := "INSERT INTO users (first_name, role, latitude, longitude) VALUES ($1, $2, $3, $4) RETURNING id, user_name, role"

	row := p.Conn.QueryRow(query, user.FirstName, user.Role, user.Latitude, user.Longitude)
	if err := row.Scan(&user.ID, &user.FirstName, &user.Role); err != nil {
		return fmt.Errorf("failed to scan user row: %w", err)
	}

	return nil
}

func (pg *PgConn) SaveMatchedRides(matchedRide models.MatchedRide) error {
	query := "INSERT INTO matched_rides (driver_id, passenger_id, ride_status, created_at)VALUES ($1, $2, $3, $4)"

	_, err := pg.Conn.Exec(query, matchedRide.DriverID, matchedRide.PassengerID, matchedRide.RideStatus, matchedRide.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save matched ride: %w", err)
	}
	return nil
}
