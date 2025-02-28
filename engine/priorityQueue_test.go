package engine

import (
	"testing"

	"github.com/Udehlee/go-Ride/models"
	"github.com/go-playground/assert/v2"
)

func TestInsert(t *testing.T) {
	pq := NewPriorityQueue()

	drivers := []models.Driver{
		{DriverID: 1, Distance: 10},
		{DriverID: 2, Distance: 5},
		{DriverID: 3, Distance: 20},
	}

	for _, d := range drivers {
		pq.Insert(d)
	}
	t.Log("Expected:", len(drivers), "Got:", pq.Len())
	assert.Equal(t, len(drivers), pq.Len())

}

func TestExtract(t *testing.T) {
	tests := []struct {
		name     string
		drivers  []models.Driver
		expected []models.Driver
	}{
		{
			name: "Extract in ascending order of distance",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 10},
				{DriverID: 2, Distance: 5}, // Nearest
				{DriverID: 3, Distance: 20},
			},
			expected: []models.Driver{
				{DriverID: 2, Distance: 5}, // nearest first
				{DriverID: 1, Distance: 10},
				{DriverID: 3, Distance: 20}, // Farthest last
			},
		},
		{
			name: "Single driver",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 7},
			},
			expected: []models.Driver{
				{DriverID: 1, Distance: 7},
			},
		},
		{
			name:     "Empty queue",
			drivers:  []models.Driver{}, // No drivers inserted
			expected: []models.Driver{}, // Nothing to extract
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue()

			for _, driver := range tt.drivers {
				pq.Insert(driver)
			}

			for _, exp := range tt.expected {
				got := pq.Extract()
				assert.Equal(t, exp.DriverID, got.DriverID)
				assert.Equal(t, exp.Distance, got.Distance)
			}

			// Ensure queue is empty at the end if all drivers are extracted
			assert.Equal(t, 0, pq.Len())
		})
	}
}

func TestUpdateDriverDistance(t *testing.T) {
	tests := []struct {
		name     string
		drivers  []models.Driver
		driverID int
		newDist  float64
		expOrder []int // Expected order of DriverIDs after update
	}{
		{
			name: "Move Up - Lower Distance",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 10},
				{DriverID: 2, Distance: 20},
				{DriverID: 3, Distance: 30},
			},
			driverID: 3, // Update driver 3's distance
			newDist:  5, // Should move to the top
			expOrder: []int{3, 1, 2},
		},
		{
			name: "Move Down - Higher Distance",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 10},
				{DriverID: 2, Distance: 5},
				{DriverID: 3, Distance: 3},
			},
			driverID: 2,  // Update driver 2's distance
			newDist:  15, // Should move down
			expOrder: []int{3, 1, 2},
		},
		{
			name: "Same Distance - No Change",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 10},
				{DriverID: 2, Distance: 20},
			},
			driverID: 1,  // Update driver 1 with the same distance
			newDist:  10, // No movement expected
			expOrder: []int{1, 2},
		},
		{
			name: "Invalid Driver ID - No Change",
			drivers: []models.Driver{
				{DriverID: 1, Distance: 10},
			},
			driverID: 99, // Driver does not exist
			newDist:  5,
			expOrder: []int{1}, // No change
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue()
			for _, d := range tt.drivers {
				pq.Insert(d)
			}

			pq.UpdateDriverDistance(tt.driverID, tt.newDist)

			gotOrder := []int{}
			for pq.Len() > 0 {
				gotOrder = append(gotOrder, pq.Extract().DriverID)
			}

			assert.Equal(t, tt.expOrder, gotOrder)
		})
	}
}
