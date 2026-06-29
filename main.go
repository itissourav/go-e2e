package main

import (
	"go-e2e/config"
	"go-e2e/controller"
	"go-e2e/db"
	"go-e2e/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("Starting server")
	config.LoadEnv()

	router := gin.Default()

	// Initialize DB
	dbConn := db.ConnectDB()
	defer dbConn.Close()

	// Initialize Controllers
	userController := controller.NewUserController(dbConn)

	// Initialize Handler
	h := handler.NewHandler(userController)

	// Register all routes
	h.RegisterRoutes(router)

	// Start server
	router.Run(":8080")
}
