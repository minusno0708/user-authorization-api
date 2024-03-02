package main

import (
	"github.com/gin-gonic/gin"

	"user-register-api/infrastructure/persistence"
	"user-register-api/interfaces/handler"
	"user-register-api/usecase"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	authHandler := handler.NewAuthHandler(userUseCase)

	r := gin.Default()
	r.GET("/", userHandler.HandleConnectionAPI)
	r.POST("/signup", userHandler.HandleUserSignup)
	r.GET("/user/:user_id", userHandler.HandleUserGet)
	r.PUT("/user/:user_id", userHandler.HandleUserPut)
	r.DELETE("/user/:user_id", userHandler.HandleUserDelete)

	r.GET("/account", userHandler.HandleUser)

	r.POST("/signin", authHandler.HandleSignin)
	r.DELETE("/signout", authHandler.HandleSignout)

	r.Run(":8080")
}
