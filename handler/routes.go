package handler

import (
	"go-e2e/controller"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserController *controller.UserController
}

func NewHandler(
	userController *controller.UserController,
) *Handler {

	return &Handler{
		UserController: userController,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {

	router.GET("/health", h.Health)

	router.POST("/user", h.UserController.SignUp)

	router.GET("/users", h.UserController.ListUsers)
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "up and running",
	})
}
