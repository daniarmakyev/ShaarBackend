package main

import (
	"log"
	"myfirstBack/internal/app"
)

func main() {
	log.Println("Starting server...")
	app.Run() // Запускаем приложение из internal/app
}