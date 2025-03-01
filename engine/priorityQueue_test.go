package engine

import (
	"testing"

	"github.com/Udehlee/go-Ride/models"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		drivers     []models.User
		insert      models.User
		expectErr   bool
		expectedLen int
	}{
		{
			name: "Insert unique drivers",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 5},
				{ID: 3, Distance: 20},
			},
			insert:      models.User{ID: 4, Distance: 15},
			expectErr:   false,
			expectedLen: 4,
		},
		{
			name: "Insert duplicate driver",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 5},
				{ID: 3, Distance: 20},
			},
			insert:      models.User{ID: 1, Distance: 15},
			expectErr:   true,
			expectedLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue()

			// Insert initial drivers
			for _, d := range tt.drivers {
				_ = pq.Insert(d)
			}

			// Insert test driver
			err := pq.Insert(tt.insert)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check queue length
			assert.Equal(t, tt.expectedLen, pq.Len())
		})
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name     string
		drivers  []models.User
		expected []models.User
	}{
		{
			name: "Extract in ascending order of distance",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 5}, // Nearest
				{ID: 3, Distance: 20},
			},
			expected: []models.User{
				{ID: 2, Distance: 5}, // nearest first
				{ID: 1, Distance: 10},
				{ID: 3, Distance: 20}, // Farthest last
			},
		},
		{
			name: "Single driver",
			drivers: []models.User{
				{ID: 1, Distance: 7},
			},
			expected: []models.User{
				{ID: 1, Distance: 7},
			},
		},
		{
			name:     "Empty queue",
			drivers:  []models.User{}, // No drivers inserted
			expected: []models.User{}, // Nothing to extract
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
				assert.Equal(t, exp.ID, got.ID)
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
		drivers  []models.User
		driverID int
		newDist  float64
		expOrder []int // Expected order of DriverIDs after update
	}{
		{
			name: "Move Up - Lower Distance",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 20},
				{ID: 3, Distance: 30},
			},
			driverID: 3, // Update driver 3's distance
			newDist:  5, // Should move to the top
			expOrder: []int{3, 1, 2},
		},
		{
			name: "Move Down - Higher Distance",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 5},
				{ID: 3, Distance: 3},
			},
			driverID: 2,  // Update driver 2's distance
			newDist:  15, // Should move down
			expOrder: []int{3, 1, 2},
		},
		{
			name: "Same Distance - No Change",
			drivers: []models.User{
				{ID: 1, Distance: 10},
				{ID: 2, Distance: 20},
			},
			driverID: 1,  // Update driver 1 with the same distance
			newDist:  10, // No movement expected
			expOrder: []int{1, 2},
		},
		{
			name: "Invalid Driver ID - No Change",
			drivers: []models.User{
				{ID: 1, Distance: 10},
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
				gotOrder = append(gotOrder, pq.Extract().ID)
			}

			assert.Equal(t, tt.expOrder, gotOrder)
		})
	}
}
