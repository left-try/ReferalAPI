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
