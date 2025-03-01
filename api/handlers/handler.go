package handlers

import (
	"net/http"

	"github.com/Udehlee/go-Ride/models"
	"github.com/Udehlee/go-Ride/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) Index(c *gin.Context) {
	c.String(200, "Welcome Home")
}

func (h *Handler) Signup(c *gin.Context) {
	panic("Unimplement")

}

func (h *Handler) Login(c *gin.Context) {
	panic("Unimplement")

}

// PassengerRequestHandler handles rideRequest of passengers
// passengers already has details in database
func (h *Handler) PassengerRequestHandler(c *gin.Context) {
	var rideReq models.RideRequest

	if err := c.ShouldBindJSON(&rideReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind passenger request"})
		return
	}

	MatchedRide, err := h.service.RequestRide(rideReq.PassengerID, rideReq.PassengerName, rideReq.Latitude, rideReq.Longitude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully matched passenger to driver",
		"userinfo": MatchedRide,
	})
}

// AddDriverHandler handles adding driver to queue
func (h *Handler) AddDriverHandler(c *gin.Context) {
	var req models.AddDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	// Ensure user exists and update role to "driver"
	err := h.service.AddDriver(req.DriverID, req.FirstName, req.Role, req.Latitude, req.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver successfully added"})
}
