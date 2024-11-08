package main

import (
	"github.com/gin-gonic/gin"
	"referralAPI/cron"
	"referralAPI/database"
	"referralAPI/routes"
	"time"
)

func main() {
	database.InitDB()       // Initialize database
	server := gin.Default() // Server

	routes.Router(server) // Initialize routes

	ticker := time.NewTicker(5 * time.Minute) // Try to delete expired codes every 5 minutes
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				cron.DeleteInTime()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	err := server.Run("localhost:8000") // Run server
	if err != nil {

		return
	}
	defer database.DB.Close() // Close database connection when server exits

}
