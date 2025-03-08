package engine

import (
	"fmt"
	"log"
	"sync"

	"github.com/Udehlee/go-Ride/models"
)

// Priority defines the priority queue interface
type Priority interface {
	Insert(driver models.User) error
	Extract() models.User
	Len() int
	GetDrivers() []models.User
	UpdateDriverDistance(index int, distance float64)
}

// PriorityQueue implements a min-heap for nearest driver selection
// driverID_Index tracks DriverID to heap index
type PriorityQueue struct {
	drivers      []models.User
	mu           sync.Mutex
	userID_Index map[int]int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		drivers:      []models.User{},
		userID_Index: make(map[int]int),
	}
}

// Len returns the number of drivers in the queue
func (pq *PriorityQueue) Len() int {
	return len(pq.drivers)
}

// GetDrivers safely returns a copy of the driver list
func (pq *PriorityQueue) GetDrivers() []models.User {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	driversCopy := make([]models.User, len(pq.drivers))
	copy(driversCopy, pq.drivers)
	return driversCopy
}

// Insert adds a driver to the queue
// Check if driver already exists
func (pq *PriorityQueue) Insert(driver models.User) error {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if _, exists := pq.userID_Index[driver.ID]; exists {
		return fmt.Errorf("driver with ID %d already exists in the queue", driver.ID)
	}

	pq.drivers = append(pq.drivers, driver)
	index := pq.Len() - 1
	pq.userID_Index[driver.ID] = index
	pq.heapifyUp(index)

	return nil
}

// Extract the nearest driver
// Correctly delete the driver from index map
// If only one element was in the heap, remove it and return
// Move the last driver to the root and update index
// Reduce heap size
func (pq *PriorityQueue) Extract() models.User {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.Len() == 0 {
		return models.User{}
	}

	nearest := pq.drivers[0]
	delete(pq.userID_Index, nearest.ID)

	if pq.Len() == 1 {
		pq.drivers = nil
		return nearest
	}

	lastIdx := pq.Len() - 1
	pq.drivers[0] = pq.drivers[lastIdx]
	pq.userID_Index[pq.drivers[0].ID] = 0

	pq.drivers = pq.drivers[:lastIdx]
	pq.heapifyDown(0)

	return nearest
}

// UpdateDriverDistance updates a driver's distance and maintains heap order
func (pq *PriorityQueue) UpdateDriverDistance(driverID int, newDist float64) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	idx, exists := pq.userID_Index[driverID]
	if !exists || idx < 0 || idx >= len(pq.drivers) {
		log.Printf(" DriverID %d not found in heap index", driverID)
		return
	}

	oldDist := pq.drivers[idx].Distance
	pq.drivers[idx].Distance = newDist

	if newDist < oldDist {
		pq.heapifyUp(idx)
	} else {
		pq.heapifyDown(idx)
	}
}

// heapifyUp restores heap order from bottom to top
// Swap and update driverID_Index
// Update driverID_Index mapping
func (pq *PriorityQueue) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if pq.drivers[idx].Distance >= pq.drivers[parent].Distance {
			break
		}
		pq.drivers[idx], pq.drivers[parent] = pq.drivers[parent], pq.drivers[idx]

		pq.userID_Index[pq.drivers[idx].ID] = idx
		pq.userID_Index[pq.drivers[parent].ID] = parent

		idx = parent
	}
}

// heapifyDown restores the heap order property by moving a node down the heap
// Update driverID_Index mapping
func (pq *PriorityQueue) heapifyDown(idx int) {
	lastIdx := len(pq.drivers) - 1

	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left <= lastIdx && pq.drivers[left].Distance < pq.drivers[smallest].Distance {
			smallest = left
		}
		if right <= lastIdx && pq.drivers[right].Distance < pq.drivers[smallest].Distance {
			smallest = right
		}
		if smallest == idx {
			break
		}

		pq.drivers[idx], pq.drivers[smallest] = pq.drivers[smallest], pq.drivers[idx]

		pq.userID_Index[pq.drivers[idx].ID] = idx
		pq.userID_Index[pq.drivers[smallest].ID] = smallest

		idx = smallest
	}
}
