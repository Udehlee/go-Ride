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
	// TODO: Implement user registration

}

func (h *Handler) Login(c *gin.Context) {
	// TODO: Implement User login

}

func (h *Handler) PassengerRequest(c *gin.Context) {
	var rideReq models.RideRequest

	if err := c.ShouldBindJSON(&rideReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind input"})
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
