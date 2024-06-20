package main

import (
	"express-ku/database"
	"express-ku/routes"
	"log"
	"net/http"
)

func main() {
	db := database.Init()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	log.Println("Server started at :9000")
	http.ListenAndServe(":9000", server)
}
