package api

import (
	"github.com/gorilla/mux"
	"valyrian.com/internal/handlers"
)

// RegisterRoutes sets up all the API routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/v1/{key}", handlers.DataPutHandler).Methods("PUT")
	router.HandleFunc("/v1/{key}", handlers.DataGetHandler).Methods("GET")
	router.HandleFunc("/v1/{key}", handlers.DataDelHandler).Methods("GET")
}
