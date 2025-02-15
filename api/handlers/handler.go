package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {

}

func (h *Handler) Index(c *gin.Context) {
	c.String(200, "Welcome Home")
}

func (h *Handler) Signup(c *gin.Context) {
	// TODO: Implement user registration

}

func (h *Handler) Login(c *gin.Context) {
	// TODO: Implement login

}
