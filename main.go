package main

import (
	"net/http"
	"wildfire/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	routes.SetupRoutes(router)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
