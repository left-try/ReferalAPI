package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	routes.Router(server)

	err := server.Run("localhost:8000")
	if err != nil {
		return
	}
}
