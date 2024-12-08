package http

import (
	"myfirstBack/internal/auth/model"
	"myfirstBack/internal/auth/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, useCase *usecase.UserUseCase) {
	router.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}
		err := useCase.CreateUser(c, &user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User created successfully"})
	})
}
