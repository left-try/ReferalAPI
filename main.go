package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()       // Initialize database
	server := gin.Default() // Server

	routes.Router(server) // Initialize routes

	err := server.Run("localhost:8000") // Run server
	if err != nil {

		return
	}
}
