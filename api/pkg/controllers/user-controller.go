package controllers

import (
	"encoding/json"
	"net/http"

	"api/pkg/models"
	"api/pkg/utils"
)

var User models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User: Alex"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
