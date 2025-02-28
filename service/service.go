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

func (s *Service) AddDriver(id int, name string, lat, lon float64) error {
	driver := models.Driver{
		DriverID:   id,
		DriverName: name,
		Latitude:   lat,
		Longitude:  lon,
		Distance:   0, // Default distance
		Available:  true,
	}

	s.pq.Insert(driver)

	err := s.store.SaveDriver(driver)
	if err != nil {
		log.Println("Error saving driver to DB:", err)
		return err
	}

	return nil
}

// RequestRide handles a ride request
// save matched driver and passsenger to db
func (s *Service) RequestRide(passengerID int, passengerName string, lat, lon float64) (models.MatchedRide, error) {
	result := make(chan models.Driver, 5)

	req := models.RideRequest{
		PassengerID:   passengerID,
		PassengerName: passengerName,
		Latitude:      lat,
		Longitude:     lon,
		Result:        result,
	}

	s.wp.Submit(req)

	select {
	case matchedDriver := <-result:
		if matchedDriver.DriverID == 0 {
			return models.MatchedRide{}, fmt.Errorf("no available drivers")
		}

		matchedRide := models.MatchedRide{
			DriverID:    matchedDriver.DriverID,
			PassengerID: passengerID,
			RideStatus:  "matched",
			CreatedAt:   time.Now(),
		}

		if err := s.store.SaveMatchedRides(matchedRide); err != nil {
			return models.MatchedRide{}, fmt.Errorf("failed to save matched ride: %v", err)
		}

		return matchedRide, nil

	case <-time.After(5 * time.Second):
		return models.MatchedRide{}, fmt.Errorf("ride request timed out")
	}
}

// Close shuts down the matching service
func (s *Service) Close() {
	s.wp.Stop()
}
