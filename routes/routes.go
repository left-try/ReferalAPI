package routes

import (
	"github.com/gin-gonic/gin"
	"referralAPI/middleware"
)

func Router() *gin.Engine {
	server := gin.Default()
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	server.POST("/signup", SignUp)
	server.POST("/login", LogInByPass)
	server.POST("/signup_by_ref/:code", SignUpByRef)

	server.GET("/get_code_by_email/:email", getCodeByEmail)
	server.GET("/get_referrals/:referrerId", getReferralsByReferrerId)
	server.POST("/create_code", createCode)
	server.DELETE("/delete_code/:code", deleteCode)
	return server
}
