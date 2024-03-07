package main

import (
	"github.com/gin-gonic/gin"

	"user-register-api/infrastructure/persistence"
	"user-register-api/interfaces/handler"
	"user-register-api/usecase"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	tokenPersistence := persistence.NewTokenPersistence()

	userUseCase := usecase.NewUserUseCase(userPersistence)
	tokenUseCase := usecase.NewTokenUseCase(tokenPersistence)

	userHandler := handler.NewUserHandler(userUseCase, tokenUseCase)
	authHandler := handler.NewAuthHandler(userUseCase, tokenUseCase)

	r := gin.Default()
	r.GET("/", userHandler.HandleConnectionAPI)
	r.POST("/signup", userHandler.HandleUserSignup)
	r.GET("/user", userHandler.HandleUserGet)
	r.PUT("/user", userHandler.HandleUserPut)
	r.DELETE("/user", userHandler.HandleUserDelete)

	r.POST("/signin", authHandler.HandleSignin)
	r.DELETE("/signout", authHandler.HandleSignout)

	r.Run(":8080")
}
