package routes

import (
	"net/http"
	"voices/handlers"
	"voices/services"
)

func RegisterRoutes() {
	// Trustee
	trusteeService := services.NewTrusteeService()
	trusteeHandler := handlers.NewTrusteeHandler(trusteeService)
	http.HandleFunc("/trustees", trusteeHandler.GetAll)

	// Citizen
	citizenService := services.NewCitizenService()
	citizenHandler := handlers.NewCitizenHandler(citizenService)
	http.HandleFunc("/citizens", citizenHandler.GetAll)
	http.HandleFunc("/citizens/create", citizenHandler.Create)

	// Engagement
	engagementService := services.NewEngagementService()
	engagementHandler := handlers.NewEngagementHandler(engagementService)
	http.HandleFunc("/engagements", engagementHandler.GetAll)
	http.HandleFunc("/engagements/create", engagementHandler.Create)
}
