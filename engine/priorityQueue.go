package engine

import (
	"sync"

	"github.com/Udehlee/go-Ride/models"
)

// Priority defines the priority queue interface
type Priority interface {
	Insert(driver models.Driver)
	Extract() models.Driver
	Len() int
	GetDrivers() []models.Driver
	UpdateDriverDistance(index int, distance float64)
}

// PriorityQueue implements a min-heap for nearest driver selection
//
//	driverID_Index tracks DriverID to heap index
type PriorityQueue struct {
	drivers        []models.Driver
	mu             sync.Mutex
	driverID_Index map[int]int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		drivers:        []models.Driver{},
		driverID_Index: make(map[int]int),
	}
}

// Len returns the number of drivers in the queue
func (pq *PriorityQueue) Len() int {
	return len(pq.drivers)
}

// GetDrivers safely returns a copy of the driver list
func (pq *PriorityQueue) GetDrivers() []models.Driver {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return append([]models.Driver{}, pq.drivers...)
}

// Insert adds a driver to the queue
// restores heap properties
func (pq *PriorityQueue) Insert(driver models.Driver) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.drivers = append(pq.drivers, driver)
	index := pq.Len() - 1
	pq.driverID_Index[driver.DriverID] = index // Store position in driverID_Index
	pq.heapifyUp(index)
}

// Extract removes and returns the nearest driver by distance
// Move last element to root and update index
func (pq *PriorityQueue) Extract() models.Driver {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.Len() == 0 {
		return models.Driver{}
	}

	nearest := pq.drivers[0]
	pq.driverID_Index[nearest.DriverID] = -1 // Mark as removed

	pq.drivers[0] = pq.drivers[pq.Len()-1]
	pq.driverID_Index[pq.drivers[0].DriverID] = 0
	pq.drivers = pq.drivers[:pq.Len()-1]

	pq.heapifyDown(0)
	return nearest
}

// UpdateDriverDistance updates a driver's distance and maintains heap order
func (pq *PriorityQueue) UpdateDriverDistance(driverID int, newDist float64) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	idx, exists := pq.driverID_Index[driverID]
	if !exists || idx == -1 {
		return // Driver not in heap
	}

	oldDist := pq.drivers[idx].Distance
	pq.drivers[idx].Distance = newDist

	if newDist < oldDist {
		pq.heapifyUp(idx) // move up if new distance is smaller
	} else {
		pq.heapifyDown(idx) // move down if new distance is larger
	}
}

// heapifyUp restores heap order from bottom to top
// Swap and update indexMap
func (pq *PriorityQueue) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.drivers[i].Distance >= pq.drivers[parent].Distance {
			break
		}

		pq.drivers[i], pq.drivers[parent] = pq.drivers[parent], pq.drivers[i]
		pq.driverID_Index[pq.drivers[i].DriverID] = i
		pq.driverID_Index[pq.drivers[parent].DriverID] = parent

		i = parent
	}
}

// heapifyDown restores heap order from top to bottom
// Swap and update indexMap
func (pq *PriorityQueue) heapifyDown(i int) {
	size := pq.Len()
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < size && pq.drivers[left].Distance < pq.drivers[smallest].Distance {
			smallest = left
		}
		if right < size && pq.drivers[right].Distance < pq.drivers[smallest].Distance {
			smallest = right
		}
		if smallest == i {
			break
		}

		pq.drivers[i], pq.drivers[smallest] = pq.drivers[smallest], pq.drivers[i]
		pq.driverID_Index[pq.drivers[i].DriverID] = i
		pq.driverID_Index[pq.drivers[smallest].DriverID] = smallest

		i = smallest
	}
}
