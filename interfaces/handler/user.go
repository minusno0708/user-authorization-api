package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user-register-api/usecase"
)

type UserHandler interface {
	HandleUserSignup(c *gin.Context)
	HandleUserGet(c *gin.Context)
	HandleUserPut(c *gin.Context)
	HandleUserDelete(c *gin.Context)
}

type userHandler struct {
	userUseCase  usecase.UserUseCase
	tokenUseCase usecase.TokenUseCase
}

func NewUserHandler(uu usecase.UserUseCase, tu usecase.TokenUseCase) UserHandler {
	return &userHandler{
		userUseCase:  uu,
		tokenUseCase: tu,
	}
}

type responseUser struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (uh userHandler) HandleUserSignup(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	if requestBody.Username == "" || requestBody.Email == "" || requestBody.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	err := uh.userUseCase.InsertUser(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (uh userHandler) HandleUserGet(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
		})
		return
	}

	user, err := uh.userUseCase.FindUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be acquired",
		"user": &responseUser{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (uh userHandler) HandleUserPut(c *gin.Context) {
	var requestBody struct {
		UpdateUsername string `json:"username"`
		UpdateEmail    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}

	if requestBody.UpdateUsername == "" && requestBody.UpdateEmail == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
		})
		return
	}

	err = uh.userUseCase.UpdateUsername(userID, requestBody.UpdateUsername, requestBody.UpdateEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User can not be updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be updated",
	})
}

func (uh userHandler) HandleUserDelete(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
		})
		return
	}

	err = uh.userUseCase.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User can not be deleted",
		})
		return
	}

	err = uh.tokenUseCase.DeleteToken(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token can not be deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be deleted",
	})
}
