package routes

import (
	"github.com/Udehlee/go-Ride/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {

	r.POST("/signup", h.Signup)
	r.POST("login", h.Login)

	r.GET("/", h.Index)
	r.POST("/request-a-ride", h.PassengerRequestHandler)
	r.POST("/add-driver", h.AddDriverHandler)

}
