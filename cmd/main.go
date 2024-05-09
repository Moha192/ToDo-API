package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Moha192/ToDo-App/database"
	"github.com/Moha192/ToDo-App/handlers"
	"github.com/joho/godotenv"
)

func main() {
	time.Sleep(time.Second * 1) // wait for docker database connection

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	log.Println("database connected")
	defer database.DB.Close()

	handlers.InitializeRoutes()
	log.Println("handlers initialised")

	http.ListenAndServe(":8080", handlers.CorsHandler(http.DefaultServeMux))
}
