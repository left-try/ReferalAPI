package routes

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}
	user.Id = 0
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func logIn(context *gin.Context) {
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
