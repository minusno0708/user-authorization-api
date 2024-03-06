package handler

import (
	"net/http"
	"user-register-api/config"
	"user-register-api/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	HandleSignin(c *gin.Context)
	HandleSignout(c *gin.Context)
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

func (ah authHandler) HandleSignin(c *gin.Context) {
	var requestBody struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	if requestBody.UserID == "" || requestBody.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database connection error",
		})
		return
	}
	defer db.Close()

	/* JWT実装後にコメントアウトを外す
	user, err := ah.userUseCase.FindUserByUserID(db, requestBody.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID or password is incorrect",
		})
		return
	}

	if user.Password != requestBody.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID or password is incorrect",
		})
		return
	}
	*/

	token, err := ah.tokenUseCase.GenerateToken(requestBody.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token can not be generated",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Token can be acquired",
		"token":   token,
	})
}

func (ah authHandler) HandleSignout(c *gin.Context) {
	var requestBody struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	if requestBody.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	err := ah.tokenUseCase.DeleteToken(requestBody.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Token can be deleted",
	})
}
