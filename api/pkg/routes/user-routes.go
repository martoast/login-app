package routes

import (
	"api/pkg/controllers"
	"api/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/auth/token", utils.Middleware(http.HandlerFunc(controllers.CreateToken))).Methods("GET")
	router.HandleFunc("/v1/user/{id}", utils.Middleware(http.HandlerFunc(controllers.GetUser))).Methods("GET")
}
