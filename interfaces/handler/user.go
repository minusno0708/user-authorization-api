package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"user-register-api/config"
	"user-register-api/usecase"
)

type UserHandler interface {
	HandleConnectionAPI(c *gin.Context)
	HandleUserSignin(c *gin.Context)
	HandleUserGet(c *gin.Context)
	HandleUserPut(c *gin.Context)
	HandleUserDelete(c *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleConnectionAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Connection Successful",
	})
}

func (uh userHandler) HandleUserSignin(c *gin.Context) {
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
		"user": user,
	})
}

func (uh userHandler) HandleUserGet(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 page not found",
		})
		return
	}

	var requestBody struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	
	if requestBody.Password == "" {
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

	user, err := uh.userUseCase.FindUserByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password != requestBody.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be acquired",
		"user": user,
	})
}

func (uh userHandler) HandleUserPut(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 page not found",
		})
		return
	}

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
	
	if requestBody.Password == "" || requestBody.Username == "" {
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

	user, err := uh.userUseCase.FindUserByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password != requestBody.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is incorrect",
		})
		return
	}

	user, err = uh.userUseCase.UpdateUsername(db, userID, requestBody.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User can not be updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be updated",
		"user": user,
	})
}

func (uh userHandler) HandleUserDelete(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 page not found",
		})
		return
	}

	var requestBody struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body does not exist",
		})
		return
	}
	
	if requestBody.Password == "" {
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

	user, err := uh.userUseCase.FindUserByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password != requestBody.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User can be deleted",
	})
}