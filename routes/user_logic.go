package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"referralAPI/models"
	"referralAPI/utils"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	user.Id = 0
	user.ReferrerId = -1
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func logInByPass(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not log in user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in", "token": token})
}

func signUpByRef(context *gin.Context) {
	code := context.Param("code")
	referrerId, err := models.GetUserIdByCode(code)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Invalid code"})
		return
	}
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	user.Id = 0
	user.ReferrerId = referrerId
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}
