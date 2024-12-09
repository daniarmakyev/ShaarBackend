package app

import (
	"log"
	"myfirstBack/config"
	"myfirstBack/internal/auth/repository/postgresUser"
	"myfirstBack/internal/auth/usecase"
	"myfirstBack/internal/postgres"
	"myfirstBack/internal/router"

	"github.com/gin-contrib/cors"
)

func Run() {
	cfg := config.LoadConfig()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer postgres.CloseDB(db)

	userRepo := postgresUser.NewPostgresUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	r := router.NewRouter(userUseCase)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
