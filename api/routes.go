package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"valyrian.com/internal/handlers"
)

// RegisterRoutes sets up all the API routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/v1/{key}", handlers.DataPutHandler).Methods("PUT")
	router.HandleFunc("/v1/{key}", handlers.DataGetHandler).Methods("GET")
	router.HandleFunc("/v1/{key}", handlers.DataDelHandler).Methods("DEL")
	router.NotFoundHandler = http.HandlerFunc(notAllowedHandler)
}

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
