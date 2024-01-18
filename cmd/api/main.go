package main

import (
	"github.com/gin-gonic/gin"

	"user-register-api/interfaces/handler"
	"user-register-api/usecase"
	"user-register-api/infrastructure/persistence"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/", userHandler.HandleConnectionAPI)
	r.POST("/signin", userHandler.HandleUserSignin)
	r.GET("/user/:user_id", userHandler.HandleUserGet)

	r.Run(":8080")
}