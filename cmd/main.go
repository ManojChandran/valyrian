package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"valyrian.com/api"
	"valyrian.com/config"
)

func main() {
	// Load config (e.g., env variables)
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("App Name:", cfg.AppName)
	fmt.Println("Running on port:", cfg.Port)
	fmt.Println("Database Host:", cfg.DB.Host)

	// Initialize Router
	r := mux.NewRouter()

	// Register Routes
	api.RegisterRoutes(r)

	// Paths to the SSL certificate and key
	//certFile := "server.crt"
	//keyFile := "server.key"

	port := cfg.Port
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
	//log.Fatal(http.ListenAndServeTLS(":"+port, certFile, keyFile, r))
}
