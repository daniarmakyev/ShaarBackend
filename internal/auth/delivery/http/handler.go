package http

import (
	"fmt"
	"myfirstBack/internal/auth/model"
	"myfirstBack/internal/auth/usecase"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, useCase *usecase.UserUseCase) {
	router.POST("/auth/register", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")
		file, err := c.FormFile("ava")

		var savePath string

		if err == nil {
			// Читаем MIME-тип файла
			f, err := file.Open()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to open file"})
				return
			}
			defer f.Close()

			buffer := make([]byte, 512) // 512 байт хватает для определения MIME-типа
			_, err = f.Read(buffer)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to read file"})
				return
			}

			// Получаем MIME-тип
			contentType := http.DetectContentType(buffer)
			if contentType != "image/png" && contentType != "image/jpeg" {
				c.JSON(400, gin.H{"error": "Invalid file type, only .jpg and .png are allowed"})
				return
			}

			savePath = filepath.Join("uploads", file.Filename)
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				c.JSON(500, gin.H{"error": "Failed to save file"})
				return
			}
		}

		user := &model.User{
			Username: username,
			Email:    email,
			Password: password,
			Avatar:   savePath,
		}

		err = useCase.CreateUser(c, user)
		if err != nil {
			if err.Error() == fmt.Sprintf("user with email %s already exists", email) {
				c.JSON(400, gin.H{"error": "User with this email already exists"})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "User created successfully"})
	})
}
