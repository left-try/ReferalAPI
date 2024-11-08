package routes

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createCode(context *gin.Context) {
	var code models.Code
	err := context.ShouldBindJSON(&code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	code.UserId = context.GetInt64("userId")
	err = code.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create code"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"code": code.Code})
}

func deleteCode(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	code := &models.Code{Id: id}

	err = code.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete code"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Code deleted successfully"})
}

func getCodeByEmail(context *gin.Context) {
	email := context.Query("email")
	code, err := models.GetCodeByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get code"})
		return
	}

	if code == "" {
		context.JSON(http.StatusNotFound, gin.H{"message": "Code not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": code})
}

func getReferralsByReferrerId(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	referrals, err := models.GetReferrals(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get referrals"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"referrals": referrals})
}
