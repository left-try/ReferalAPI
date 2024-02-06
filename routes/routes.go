package routes

import (
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)
	server.POST("/signup", signUp)
	server.POST("/login", logIn)
	authenticated.POST("/events/;id/register", registerForEvent)

	authenticated.PUT("/events/:eventId", updateEvent)

	authenticated.DELETE("/events/:eventId", deleteEvent)
	authenticated.DELETE("/events/;id/register", unregisterForEvent)
}
