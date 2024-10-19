package handler

import (
	"net/http"
	"user-register-api/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	HandleLogin(c *gin.Context)
	HandleLogout(c *gin.Context)
}

type authHandler struct {
	userUseCase  usecase.UserUseCase
	tokenUseCase usecase.TokenUseCase
}

func NewAuthHandler(uu usecase.UserUseCase, tu usecase.TokenUseCase) AuthHandler {
	return &authHandler{
		userUseCase:  uu,
		tokenUseCase: tu,
	}
}

func (ah authHandler) HandleLogin(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	if requestBody.Username == "" || requestBody.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	userID, err := ah.userUseCase.ValidateUser(requestBody.Username, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID or password is incorrect",
		})
		return
	}

	user, err := ah.userUseCase.FindUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID or password is incorrect",
		})
		return
	}
	if user.Username != requestBody.Username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID or password is incorrect",
		})
		return
	}

	tokenString, err := ah.tokenUseCase.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token can not be generated",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Token can be acquired",
		"token":   tokenString,
	})
}

func (ah authHandler) HandleLogout(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	_, err := ah.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
		})
		return
	}

	err = ah.tokenUseCase.DeleteToken(tokenString)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Token can not be deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token can be deleted",
	})
}
