package handlers

import (
	"github.com/Udehlee/go-Ride/client"
	"github.com/Udehlee/go-Ride/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	client  client.OSMClient
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
