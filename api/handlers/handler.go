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

func (h *Handler) PassengerRequestHandler(c *gin.Context) {
	var rideReq models.RideRequest

	if err := c.ShouldBindJSON(&rideReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind passenger  request"})
		return
	}

	MatchedRide, err := h.service.RequestRide(rideReq.PassengerID, rideReq.PassengerName, rideReq.Latitude, rideReq.Longitude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to request a  ride"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully matched passenger to driver",
		"userinfo": MatchedRide,
	})

}

func (h *Handler) AddDriverHandler(c *gin.Context) {
	var addDriverReq models.Driver

	if err := c.ShouldBindJSON(&addDriverReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind driver details"})
		return
	}

	err := h.service.AddDriver(addDriverReq.DriverID, addDriverReq.DriverName, addDriverReq.Latitude, addDriverReq.Longitude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to add driver"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully added driver to queue",
	})

}
