package service

import (
	"fmt"
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

// AddDriver adds a new driver to the system
func (s *Service) AddDriver(id int, name string, lat, lon float64) error {
	driver := models.Driver{
		DriverID:   id,
		DriverName: name,
		Latitude:   lat,
		Longitude:  lon,
	}
	s.pq.Insert(driver)

	err := s.store.SaveDriver(driver)
	if err != nil {
		return err
	}
	return nil
}

// RequestRide handles a ride request
// and finds a driver
// save matched driver and passsenger to db
func (s *Service) RequestRide(passengerID int, passengerName string, lat, lon float64) (models.MatchedRide, error) {
	result := make(chan models.Driver)

	req := models.RideRequest{
		PassengerID:   passengerID,
		PassengerName: passengerName,
		Latitude:      lat,
		Longitude:     lon,
		Result:        result,
	}

	s.wp.Submit(req)
	matchedDriver := <-result

	if matchedDriver.DriverID == 0 {
		fmt.Println("No available drivers")
		return models.MatchedRide{}, fmt.Errorf("no available drivers")
	}

	return s.SaveMatchedRideRecord(matchedDriver.DriverID, passengerID)
}

func (s *Service) SaveMatchedRideRecord(driverID, passengerID int) (models.MatchedRide, error) {
	matchedRide := models.MatchedRide{
		DriverID:    driverID,
		PassengerID: passengerID,
		RideStatus:  "matched",
		CreatedAt:   time.Now(),
	}

	err := s.store.SaveMatchedRides(matchedRide)
	if err != nil {
		return models.MatchedRide{}, fmt.Errorf("failed to save matched ride: %w", err)
	}

	fmt.Printf("Rider %d matched with Driver %d\n", passengerID, driverID)
	return matchedRide, nil
}

// Close shuts down the matching service
func (s *Service) Close() {
	s.wp.Stop()
}
