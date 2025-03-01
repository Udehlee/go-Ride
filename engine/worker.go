package engine

import (
	"log"
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
		rideRequests: make(chan models.RideRequest, 10),
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

func (wp *WorkerPool) processRequest(req models.RideRequest) {
	if wp.pq.Len() == 0 {
		req.Result <- models.User{}
		return
	}

	drivers := wp.pq.GetDrivers()
	for _, driver := range drivers {
		distance := utils.CalculateDistance(
			req.Latitude, req.Longitude,
			driver.Latitude, driver.Longitude,
		)

		wp.pq.UpdateDriverDistance(driver.ID, distance)
	}

	// Extract the nearest driver after updating distances
	req.Result <- wp.pq.Extract()
}

// Submit adds a new ride request to the queue
func (wp *WorkerPool) Submit(req models.RideRequest) {
	select {
	case wp.rideRequests <- req: // Send if there's space
	default:
		log.Println("WorkerPool queue is full, dropping request")

	}
}

// Stop shuts down the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.rideRequests)
	wp.wg.Wait()
}
