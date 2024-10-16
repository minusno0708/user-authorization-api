package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user-register-api/config"
	"user-register-api/infrastructure/persistence"
	"user-register-api/interfaces/handler"
	"user-register-api/usecase"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		panic(err)
	}
	defer cdb.Close()

	userPersistence := persistence.NewUserPersistence(db)
	tokenPersistence := persistence.NewTokenPersistence(cdb)

	userUseCase := usecase.NewUserUseCase(userPersistence)
	tokenUseCase := usecase.NewTokenUseCase(tokenPersistence)

	userHandler := handler.NewUserHandler(userUseCase, tokenUseCase)
	authHandler := handler.NewAuthHandler(userUseCase, tokenUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection Successful",
		})
	})

	r.POST("/signup", userHandler.HandleUserSignup)
	r.GET("/user", userHandler.HandleUserGet)
	r.PUT("/user", userHandler.HandleUserPut)
	r.DELETE("/user", userHandler.HandleUserDelete)

	r.POST("/login", authHandler.HandleLogin)
	r.DELETE("/logout", authHandler.HandleLogout)

	r.Run(":8080")
}
