package router

// @title MyFirstBack API
// @version 1.0
// @description API documentation for MyFirstBack.
// @host localhost:8080
// @BasePath /api

import (
	"myfirstBack/docs"
	"myfirstBack/internal/auth/delivery/http"
	"myfirstBack/internal/auth/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(useCase *usecase.UserUseCase) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	http.RegisterAuthRoutes(api, useCase)

	return r
}
