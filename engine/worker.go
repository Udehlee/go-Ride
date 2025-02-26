package engine

import (
	"sync"

	"github.com/Udehlee/go-Ride/models"
	"github.com/Udehlee/go-Ride/utils"
)

// WorkerP defines the worker pool interface
type WorkerP interface {
	Start()
	Submit(req models.RideRequest)
	Stop()
}

// WorkerPool manages concurrent ride-matching workers
type WorkerPool struct {
	workers      int
	rideRequests chan models.RideRequest
	pq           Priority
	wg           sync.WaitGroup
}

func NewWorkerPool(workers int, pq Priority) *WorkerPool {
	return &WorkerPool{
		workers:      workers,
		rideRequests: make(chan models.RideRequest),
		pq:           pq,
	}
}

// Start launches worker goroutines
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker processes incoming ride requests
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for rideReq := range wp.rideRequests {
		wp.processRequest(rideReq)
	}
}

// processRequest finds the nearest driver for a passenger
// Update driver distances
// Extract nearest driver
func (wp *WorkerPool) processRequest(req models.RideRequest) {
	if wp.pq.Len() == 0 {
		req.Result <- models.Driver{}
		return
	}

	drivers := wp.pq.GetDrivers() // Safe copy
	for i := range drivers {
		distance := utils.CalculateDistance(
			req.Latitude, req.Longitude,
			drivers[i].Latitude, drivers[i].Longitude,
		)
		wp.pq.UpdateDriverDistance(i, distance) // Maintain heap correctness
	}

	req.Result <- wp.pq.Extract()
}

// Submit adds a new ride request to the queue
func (wp *WorkerPool) Submit(req models.RideRequest) {
	wp.rideRequests <- req
}

// Stop shuts down the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.rideRequests)
	wp.wg.Wait()
}
