package handlers

import (
	"net/http"

	"github.com/Udehlee/go-Ride/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Location(c *gin.Context) {
	var locationReq models.Location

	if err := c.ShouldBindJSON(&locationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to read location details"})
		return
	}

	Addr, err := h.client.CurrentAddr(c, locationReq.Latitude, locationReq.Longitude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to get current address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "successful gotten the current address",
		"userinfo": Addr,
	})

}
