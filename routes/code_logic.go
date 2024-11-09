package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag"
	"net/http"
	"referralAPI/models"
	"strconv"
)

// createCode godoc
//
// @Summary      Create a new code
// @Description  Takes a code JSON and stores it in the database. Returns the saved JSON object.
// @Tags         code
// @Accept       json
// @Produce      json
// @Param        code  body      models.Code  true  "Code data"
// @Success      201   {object}  models.Code  "Code created"
// @Failure      400   {object}  map[string]string  "Invalid request"
// @Failure      500   {object}  map[string]string  "Failed to create code"
// @Router       /create_code [post]
func createCode(context *gin.Context) {
	var code models.Code
	err := context.ShouldBindJSON(&code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err = code.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create code"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"code": code.Code})
}

// deleteCode godoc
//
// @Summary      Delete a code by ID
// @Description  Deletes a code specified by its code ID.
// @Tags         code
// @Produce      json
// @Param        code  path      string  true  "Code ID"
// @Success      200   {object}  map[string]string  "Code deleted successfully"
// @Failure      500   {object}  map[string]string  "Failed to delete code"
// @Router       /delete_code/{code} [delete]
func deleteCode(context *gin.Context) {
	code := context.Param("code")
	var Code models.Code
	Code.Code = code

	err := Code.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete code"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Code deleted successfully"})
}

// getCodeByEmail godoc
//
// @Summary      Get code by email
// @Description  Retrieves a code based on the provided email address.
// @Tags         code
// @Produce      json
// @Param        email  path      string  true  "Email address"
// @Success      200    {object}  map[string]string  "Code found"
// @Failure      404    {object}  map[string]string  "Code not found"
// @Failure      500    {object}  map[string]string  "Failed to get code"
// @Router       /get_code_by_email/{email} [get]
func getCodeByEmail(context *gin.Context) {
	email := context.Param("email")
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

// getReferralsByReferrerId godoc
//
// @Summary      Get referrals by referrer ID
// @Description  Retrieves a list of referrals based on the referrer's user ID.
// @Tags         code
// @Produce      json
// @Param        referrerId  path      int  true  "Referrer ID"
// @Success      200         {object}  map[string][]integer  "List of referrals"
// @Failure      400         {object}  map[string]string             "Invalid user ID"
// @Failure      500         {object}  map[string]string             "Failed to get referrals"
// @Router       /get_referrals/{referrerId} [get]
func getReferralsByReferrerId(context *gin.Context) {
	referrerId, err := strconv.ParseInt(context.Param("referrerId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	referrals, err := models.GetReferrals(referrerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get referrals"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"referrals": referrals})
}
