package routes

import (
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	server.POST("/signup", signUp)
	server.POST("/login_by_pass", logInByPass)
	server.POST("/login_by_ref", logInByRef)
}
