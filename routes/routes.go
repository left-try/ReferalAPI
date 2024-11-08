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
	server.POST("/login_by_ref", signUpByRef)

	server.GET("/get_code_by_email/:email", getCodeByEmail)
	server.POST("/create_code", createCode)
	server.DELETE("/delete_code/:id", deleteCode)
}
