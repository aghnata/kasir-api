package main

import (
	"fmt"
	"kasir-api/routes"
	"net/http"
)

func main() {
	// Setup all routes
	routes.SetupRoutes()

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
