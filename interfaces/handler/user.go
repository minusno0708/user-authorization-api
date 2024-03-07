package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user-register-api/config"
	"user-register-api/usecase"
)

type UserHandler interface {
	HandleConnectionAPI(c *gin.Context)
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
}

func (uh userHandler) HandleConnectionAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Connection Successful",
	})
}

func (uh userHandler) HandleUserSignup(c *gin.Context) {
	var requestBody struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
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

	user, err := uh.userUseCase.InsertUser(db, requestBody.UserID, requestBody.Username, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": &responseUser{
			UserID:   user.UserID,
			Username: user.Username,
		},
	})
}

func (uh userHandler) HandleUserGet(c *gin.Context) {
	var requestBody struct {
		TokenString string `json:"token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}

	if requestBody.TokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	userID, err := uh.tokenUseCase.ValidateToken(requestBody.TokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
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

	user, err := uh.userUseCase.FindUserByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be acquired",
		"user": &responseUser{
			UserID:   user.UserID,
			Username: user.Username,
		},
	})
}

func (uh userHandler) HandleUserPut(c *gin.Context) {
	var requestBody struct {
		NewUsername string `json:"username"`
		TokenString string `json:"token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}

	if requestBody.TokenString == "" || requestBody.NewUsername == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	userID, err := uh.tokenUseCase.ValidateToken(requestBody.TokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
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

	user, err := uh.userUseCase.UpdateUsername(db, userID, requestBody.NewUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User can not be updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be updated",
		"user": &responseUser{
			UserID:   user.UserID,
			Username: user.Username,
		},
	})
}

func (uh userHandler) HandleUserDelete(c *gin.Context) {
	var requestBody struct {
		TokenString string `json:"token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}

	if requestBody.TokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	userID, err := uh.tokenUseCase.ValidateToken(requestBody.TokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to authenticate",
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

	err = uh.userUseCase.DeleteUser(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User can not be deleted",
		})
		return
	}

	err = uh.tokenUseCase.DeleteToken(requestBody.TokenString)
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
