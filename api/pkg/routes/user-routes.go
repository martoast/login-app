package routes

import (
	"api/pkg/controllers"
	"api/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/auth/token", utils.Middleware(http.HandlerFunc(controllers.CreateToken))).Methods("GET")
	router.HandleFunc("/api/users/{id}", utils.Middleware(http.HandlerFunc(controllers.GetUserById))).Methods("GET")
}
