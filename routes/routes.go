package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"referralAPI/middleware"
)

func Router(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	server.POST("/signup", signUp)
	server.POST("/login", logInByPass)
	server.POST("/signup_by_ref/:code", signUpByRef)

	server.GET("/get_code_by_email/:email", getCodeByEmail)
	server.GET("/get_referrals/:referrerId", getReferralsByReferrerId)
	server.POST("/create_code", createCode)
	server.DELETE("/delete_code/:code", deleteCode)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
