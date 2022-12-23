package main

import (
	"log"
	"net/http"

	"api/pkg/routes"
	"api/pkg/utils"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	utils.SetupGoGuardian()

	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	log.Printf("server started and listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
