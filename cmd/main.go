package main

import (
	"log"
	"net/http"
	"voices/db"
	"voices/routes"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	routes.RegisterRoutes()

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
