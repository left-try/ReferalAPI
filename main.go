package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"referralAPI/cron"
	"referralAPI/database"
	_ "referralAPI/docs"
	"referralAPI/routes"
	"time"
)

//	@title			Referral API
//	@version		1.0
//	@description	A referral links management service API in Go using Gin framework.

//	@host		localhost:8000
//	@BasePath	/

func main() {
	database.InitDB() // Initialize database

	router := routes.Router() // Initialize routes

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	err := router.Run("localhost:8000") // Run server
	if err != nil {

		return
	}
	defer database.DB.Close() // Close database connection when server exits

}
