package service

import (
	"fmt"

	"github.com/Udehlee/go-Ride/engine"
	"github.com/Udehlee/go-Ride/models"
)

// MatchingService handles rider-driver matching
type Service struct {
	pq *engine.PriorityQueue
	wp *engine.WorkerPool
}

func NewService(workers int) *Service {
	pq := engine.NewPriorityQueue()
	wp := engine.NewWorkerPool(workers, pq)
	wp.Start()
	return &Service{pq: pq, wp: wp}
}

// AddDriver adds a new driver to the system
func (s *Service) AddDriver(id int, lat, lon float64) {
	driver := models.Driver{
		ID:        id,
		Latitude:  lat,
		Longitude: lon,
	}
	s.pq.Insert(driver)
}

// RequestRide handles a ride request and finds a driver
func (s *Service) RequestRide(passengerID int, passengerName string, lat, lon float64) (models.Driver, error) {
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

	if matchedDriver.ID == 0 {
		fmt.Println("No available drivers")
	} else {
		fmt.Printf("Rider %d matched with Driver %d\n", passengerID, matchedDriver.ID)
	}

	return matchedDriver, nil
}

// Close shuts down the matching service
func (s *Service) Close() {
	s.wp.Stop()
}
