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
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
		return
	}
	if requestBody.Username == "" || requestBody.Email == "" || requestBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
		return
	}

	user, err := uh.userUseCase.InsertUser(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Create Successfully",
		"user": &responseUser{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (uh userHandler) HandleUserGet(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token validation error",
		})
		return
	}

	user, err := uh.userUseCase.FindUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get successfully",
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
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
		return
	}

	if requestBody.UpdateUsername == "" && requestBody.UpdateEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
		return
	}

	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token validation error",
		})
		return
	}

	err = uh.userUseCase.UpdateUser(userID, requestBody.UpdateUsername, requestBody.UpdateEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update user",
		})
		return
	}

	user, err := uh.userUseCase.FindUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Update successfully",
		"user": &responseUser{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (uh userHandler) HandleUserDelete(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	userID, err := uh.tokenUseCase.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token validation error",
		})
		return
	}

	err = uh.userUseCase.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete user",
		})
		return
	}

	err = uh.tokenUseCase.DeleteToken(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Delete successfully",
	})
}
