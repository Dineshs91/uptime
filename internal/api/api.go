package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func authRoutes(router *mux.Router) {
	s := router.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", LoginHandler).Methods("POST")
	s.HandleFunc("/logout", LogoutHandler).Methods("DELETE")
	s.HandleFunc("/register", RegisterHandler).Methods("POST")
	s.HandleFunc("/forgot-password", ForgotPasswordHandler).Methods("POST")
	s.HandleFunc("/reset-password", ResetPasswordHandler).Methods("POST")
}

func monitoringDetailsRoutes(router *mux.Router) {
	router.HandleFunc("/monitoring-urls", AddMonitoringURLHandler).Methods("POST")
	router.HandleFunc("/monitoring-urls", GetMonitoringURLsHandler).Methods("GET")
	router.HandleFunc("/monitoring-urls", UpdateMonitoringURLHandler).Methods("PUT")
	router.HandleFunc("/monitoring-urls/{monitoringURLID}", DeleteMonitoringURLHandler).Methods("DELETE")
}

func integrationRoutes(router *mux.Router) {
	router.HandleFunc("/integrations", AddIntegrationHandler).Methods("POST")
	router.HandleFunc("/integrations", GetIntegrationsHandler).Methods("GET")
	router.HandleFunc("/integrations/{integrationID}", GetIntegrationHandler).Methods("GET")
	router.HandleFunc("/integrations/{integrationID}", DeleteIntegrationHandler).Methods("DELETE")
}

// StartServer Start the server.
func StartServer() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.HandleFunc("/", HomeHandler)

	monitoringDetailsRoutes(router)
	integrationRoutes(router)
	authRoutes(router)

	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(router),
	)
}
