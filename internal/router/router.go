package router

import (
	"myfirstBack/internal/auth/delivery/http"
	"myfirstBack/internal/auth/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(useCase *usecase.UserUseCase) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	http.RegisterAuthRoutes(api, useCase)

	return r
}
