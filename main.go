package main

import (
	"fmt"
	"go-e2e/config"
	"go-e2e/controller"
	"go-e2e/db"
	"go-e2e/handler"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("Starting server")
	config.LoadEnv()

	router := gin.Default()
	router.Use(LoggerMiddleware())

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

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// After request
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s in %v\n",
			c.Request.Method,
			path,
			strconv.Itoa(c.Writer.Status()),
			duration,
		)
	}
}
