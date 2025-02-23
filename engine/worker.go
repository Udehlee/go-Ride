package engine

import (
	"sync"

	"github.com/Udehlee/go-Ride/models"
	"github.com/Udehlee/go-Ride/utils"
)

// WorkerPool manages concurrent worker execution
type WorkerPool struct {
	workers     int
	rideRequest chan models.RideRequest
	pq          *PriorityQueue
	wg          sync.WaitGroup
}

func NewWorkerPool(workers int, pq *PriorityQueue) *WorkerPool {
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
		go wp.worker()
	}
}

// worker processes incoming ride request
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for rideReq := range wp.rideRequest {
		wp.processRequest(rideReq)
	}
}

// processTask finds the nearest driver for a passenger
func (wp *WorkerPool) processRequest(req models.RideRequest) {
	wp.pq.mu.Lock()
	defer wp.pq.mu.Unlock()

	if len(wp.pq.drivers) == 0 {
		req.Result <- models.Driver{}
		return
	}

	for i := range wp.pq.drivers {
		wp.pq.drivers[i].Distance = utils.CalculateDistance(
			req.Latitude, req.Longitude,
			wp.pq.drivers[i].Latitude, wp.pq.drivers[i].Longitude,
		)
	}
	req.Result <- wp.pq.Extract()
}

// Submit adds a new matching task
func (wp *WorkerPool) Submit(req models.RideRequest) {
	wp.rideRequest <- req
}

// Stop closes the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.rideRequest)
	wp.wg.Wait()
}
