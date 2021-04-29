package router

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	// Setting up endpoints
	router.HandleFunc("/api/home", controllers.Home).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/gettoken", controllers.GetToken).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", controllers.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/register", controllers.RegisterHandler).Methods("POST", "OPTIONS")
	return router
}