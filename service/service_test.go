package service

import (
	"errors"
	"testing"
	"time"

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
func (m *MockStore) SaveDriver(driver models.Driver) error {
	args := m.Called(driver)
	return args.Error(0)
}

func (m *MockStore) SavePassenger(passenger models.Passenger) error {
	args := m.Called(passenger)
	return args.Error(0)
}

func (m *MockStore) Submit(req models.RideRequest) {
	go func() {
		time.Sleep(500 * time.Millisecond)                               // Simulate processing delay
		req.Result <- models.Driver{DriverID: 1, DriverName: "John Doe"} // Modify per test case
	}()
}
func TestAddDriver(t *testing.T) {
	tests := []struct {
		name       string
		driverID   int
		driverName string
		lat, lon   float64
		mockErr    error
		expErr     bool
	}{
		{
			name:       "Successful Add",
			driverID:   1,
			driverName: "Ada Mikel",
			lat:        40.7128,
			lon:        -74.0060,
			mockErr:    nil,
			expErr:     false,
		},
		{
			name:       "DB Save Failure",
			driverID:   2,
			driverName: "John Doe",
			lat:        34.0522,
			lon:        -118.2437,
			mockErr:    errors.New("DB error"),
			expErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := new(MockStore)
			pq := engine.NewPriorityQueue()
			service := &Service{pq: pq, store: mockStore}

			mockStore.On("SaveDriver", mock.Anything).Return(tt.mockErr)

			err := service.AddDriver(tt.driverID, tt.driverName, tt.lat, tt.lon)
			if tt.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Verify driver is in priority queue
				assert.Equal(t, 1, pq.Len())

				addedDriver := pq.Extract()
				assert.Equal(t, tt.driverID, addedDriver.DriverID)
				assert.Equal(t, tt.driverName, addedDriver.DriverName)
				assert.Equal(t, tt.lat, addedDriver.Latitude)
				assert.Equal(t, tt.lon, addedDriver.Longitude)
				assert.True(t, addedDriver.Available)
			}

			// Verify store.SaveDriver was called
			mockStore.AssertExpectations(t)
		})
	}
}
