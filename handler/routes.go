package handler

import (
	"go-e2e/controller"
	"go-e2e/middleware"

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
	router.Use(middleware.LoggerMiddleware())

	router.GET("/v1//health", h.Health)

	router.POST("/v1/user", h.UserController.SignUp)

	router.POST("/v1/login", h.UserController.Login)

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/users", h.UserController.ListUsers)
	}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "up and running",
	})
}
