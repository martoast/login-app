package controllers

import (
	"net/http"

	"api/pkg/models"
)

var User models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User: Alex"))
}
