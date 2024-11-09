package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"referralAPI/models"
	"referralAPI/utils"
)

// signUp godoc
//
// @Summary      Register a new user
// @Description  Registers a new user by accepting a JSON object and saving it to the database. Returns the created user.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      201   {object}  map[string]interface{}  "User created"
// @Failure      400   {object}  map[string]string       "Invalid request data"
// @Failure      500   {object}  map[string]string       "Internal server error"
// @Router       /signup [post]
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

// logInByPass godoc
//
// @Summary      Log in user with email and password
// @Description  Authenticates a user based on email and password, and returns an authentication token.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User credentials"
// @Success      200   {object}  map[string]interface{}  "User logged in"
// @Failure      400   {object}  map[string]string       "Invalid request data"
// @Failure      401   {object}  map[string]string       "Unauthorized"
// @Failure      500   {object}  map[string]string       "Internal server error"
// @Router       /login [post]
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

// signUpByRef godoc
//
// @Summary      Register a new user with referral code
// @Description  Registers a new user using a referral code to link them to an existing user as the referrer.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        code  path      string      true  "Referral code"
// @Param        user  body      models.User true  "User data"
// @Success      201   {object}  map[string]interface{}  "User created"
// @Failure      400   {object}  map[string]string       "Invalid request data"
// @Failure      404   {object}  map[string]string       "Referral code not found"
// @Failure      500   {object}  map[string]string       "Internal server error"
// @Router       /signup_by_ref/{code} [post]
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
