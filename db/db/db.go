package db

import (
	"database/sql"
	"fmt"

	"github.com/Udehlee/go-Ride/models"
)

type Repo interface {
	SaveUser(driver models.User) error
	SaveMatchedRides(driver models.MatchedRide) error
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
	query := "INSERT INTO users (email, pass_word, role) VALUES ($1,$2,$3) RETURNING user_id, email, user_role"

	row := p.Conn.QueryRow(query, user.Email, user.Password, user.Role)
	if err := row.Scan(&user.UserID, &user.Email, &user.Role); err != nil {
		return fmt.Errorf("failed to scan users row: %w", err)
	}

	return nil
}

func (p *PgConn) SaveMatchedRides(ride models.MatchedRide) error {
	query := "INSERT INTO passengers (driverId, passengerId) VALUES ($1,$2)"

	row := p.Conn.QueryRow(query, ride.DriverID, ride.PassengerID)
	if err := row.Scan(&ride.DriverID, &ride.PassengerID); err != nil {
		return fmt.Errorf("failed to scan passengers row: %w", err)
	}

	return nil
}
