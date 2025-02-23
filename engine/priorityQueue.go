package engine

import (
	"sync"

	"github.com/Udehlee/go-Ride/models"
)

// PriorityQueue implements a min-heap for nearest driver selection
// driverIndx assigns driver id to the index in the []Driver
type PriorityQueue struct {
	drivers []models.Driver
	mu      sync.Mutex
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		drivers: []models.Driver{},
	}
	return pq
}

// Insert adds a driver to the queue
// allows for quick lookup of drivers position in the []Driver
func (pq *PriorityQueue) Insert(driver models.Driver) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.drivers = append(pq.drivers, driver)
}

// Extract removes and returns the nearest driver
func (pq *PriorityQueue) Extract() models.Driver {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if len(pq.drivers) == 0 {
		return models.Driver{}
	}

	minIdx := 0
	for i, driver := range pq.drivers {
		if driver.Distance < pq.drivers[minIdx].Distance {
			minIdx = i
		}
	}

	driver := pq.drivers[minIdx]
	pq.drivers = append(pq.drivers[:minIdx], pq.drivers[minIdx+1:]...)
	return driver
}
