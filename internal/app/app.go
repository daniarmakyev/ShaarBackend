package app

import (
	"log"
	"myfirstBack/config"
	"myfirstBack/internal/auth/repository/postgresUser"
	"myfirstBack/internal/auth/usecase"
	"myfirstBack/internal/postgres"
	"myfirstBack/internal/router"
)

func Run() {

	cfg := config.LoadConfig()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	userRepo := postgresUser.NewPostgresUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	r := router.NewRouter(userUseCase)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer postgres.CloseDB(db)
}
