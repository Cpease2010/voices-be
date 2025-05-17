package routes

import (
	"net/http"
	"voices/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/trustees", handlers.HandleTrustees)
	http.HandleFunc("/engagements", handlers.HandleEngagements)
	http.HandleFunc("/citizens/", handlers.HandleCitizen)
}
