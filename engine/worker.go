package engine

import (
	"sync"

	"github.com/Udehlee/go-Ride/models"
	"github.com/Udehlee/go-Ride/utils"
)

// WorkerP interface for worker pool operations
type WorkerP interface {
	Start()
	Submit(req models.RideRequest)
	Stop()
}

// WorkerPool manages concurrent worker execution
type WorkerPool struct {
	workers     int
	rideRequest chan models.RideRequest
	pq          Priority
	wg          sync.WaitGroup
}

func NewWorkerPool(workers int, pq Priority) *WorkerPool {
	return &WorkerPool{
		workers:     workers,
		rideRequest: make(chan models.RideRequest),
		pq:          pq,
	}
}

// Start begins processing ride requests with workers
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.Worker()
	}
}

// worker processes incoming ride request
func (wp *WorkerPool) Worker() {
	defer wp.wg.Done()
	for rideReq := range wp.rideRequest {
		wp.processRequest(rideReq)
	}
}

// processRequest finds the nearest driver for a passenger
// Updating driver distances
func (wp *WorkerPool) processRequest(req models.RideRequest) {
	if wp.pq.Len() == 0 {
		req.Result <- models.Driver{}
		return
	}

	for i := range wp.pq.(*PriorityQueue).drivers {
		wp.pq.(*PriorityQueue).drivers[i].Distance = utils.CalculateDistance(
			req.Latitude, req.Longitude,
			wp.pq.(*PriorityQueue).drivers[i].Latitude, wp.pq.(*PriorityQueue).drivers[i].Longitude,
		)
	}

	req.Result <- wp.pq.Extract()
}

// Submit adds a new matching request
func (wp *WorkerPool) Submit(req models.RideRequest) {
	wp.rideRequest <- req
}

// Stop closes the worker pool
// and waits for them to finish
func (wp *WorkerPool) Stop() {
	close(wp.rideRequest)
	wp.wg.Wait()
}
