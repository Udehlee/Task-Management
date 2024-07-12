package main

import (
	"log"
	"net/http"

	"github.com/Udehlee/Task-Management/api/handler"
	"github.com/Udehlee/Task-Management/api/routes"
	"github.com/Udehlee/Task-Management/db"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	config, err := db.LoadConfig()
	if err != nil {
		log.Fatal("error loading config")
	}

	conn, err := db.InitDB(config)
	if err != nil {
		log.Fatal("error connecting to db")
	}

	// mux := http.NewServeMux()
	h := handler.NewHandler(conn)

	r := routes.SetupRoutes(h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}

}
