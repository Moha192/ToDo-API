package main

import (
	"log"
	"net/http"

	"github.com/Moha192/ToDo-App/database"
	"github.com/Moha192/ToDo-App/handlers"
)

func main() {
	database.InitDB()
	log.Println("database connected")
	defer database.DB.Close()

	handlers.InitializeRoutes()

	log.Println("handlers initialised")
	http.ListenAndServe(":8080", handlers.CorsHandler(http.DefaultServeMux))
}
