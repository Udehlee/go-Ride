package service

import (
	"testing"

	"github.com/Udehlee/go-Ride/engine"
	"github.com/Udehlee/go-Ride/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) SaveMatchedRides(ride models.MatchedRide) error {
	args := m.Called(ride)
	return args.Error(0)
}
func (m *MockStore) SaveUser(driver models.User) error {
	args := m.Called(driver)
	return args.Error(0)
}

func TestAddDriver(t *testing.T) {
	tests := []struct {
		name       string
		driverID   int
		driverName string
		role       string
		lat, lon   float64
		expErr     bool
	}{
		{
			name:       "Successful Add",
			driverID:   1,
			driverName: "Ada Mikel",
			role:       "Driver",
			lat:        40.7128,
			lon:        -74.0060,
			expErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := engine.NewPriorityQueue()
			service := &Service{pq: pq}

			err := service.AddDriver(tt.driverID, tt.driverName, tt.role, tt.lat, tt.lon)
			if tt.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
