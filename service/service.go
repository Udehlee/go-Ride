package service

import (
	"fmt"
	"log"
	"time"

	"github.com/Udehlee/go-Ride/db/db"
	"github.com/Udehlee/go-Ride/engine"
	"github.com/Udehlee/go-Ride/models"
)

type Service struct {
	store db.Repo
	pq    engine.Priority
	wp    engine.WorkerP
}

func NewService(db db.Repo, pq engine.Priority, wp engine.WorkerP) *Service {
	return &Service{
		store: db,
		pq:    pq,
		wp:    wp,
	}
}

func (s *Service) AddDriver(id int, name string, role string, lat, lon float64) error {
	driver := models.User{
		ID:        id,
		FirstName: name,
		Role:      role,
		Latitude:  lat,
		Longitude: lon,
		Distance:  0,
		Available: true,
	}

	if err := s.pq.Insert(driver); err != nil {
		return fmt.Errorf("error adding driver to queue%s", err)
	}

	return nil
}

// RequestRide handles a ride request
// Saves matched driver and passenger to the database
func (s *Service) RequestRide(passengerID int, passengerName string, lat, lon float64) (models.MatchedRide, error) {
	result := make(chan models.User, 1)

	req := models.RideRequest{
		PassengerID:   passengerID,
		PassengerName: passengerName,
		Latitude:      lat,
		Longitude:     lon,
		Result:        result,
	}

	log.Println("submitting  ride request")
	s.wp.Submit(req)

	select {
	case <-time.After(10 * time.Second): //
		return models.MatchedRide{}, fmt.Errorf("ride request timed out")

	case matchedUser := <-result:
		if matchedUser.ID == 0 {
			return models.MatchedRide{}, fmt.Errorf("no available drivers")
		}

		matchedRide := models.MatchedRide{
			DriverID:    matchedUser.ID,
			PassengerID: passengerID,
			RideStatus:  "matched",
			CreatedAt:   time.Now(),
		}

		if err := s.store.SaveMatchedRides(matchedRide); err != nil {
			return models.MatchedRide{}, fmt.Errorf("failed to save matched ride: %v", err)
		}

		return matchedRide, nil
	}
}

// Close shuts down the matching service
func (s *Service) Close() {
	s.wp.Stop()
}
