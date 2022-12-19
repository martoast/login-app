package controllers

import (
	"net/http"
	"time"

	"api/pkg/models"

	"github.com/golang-jwt/jwt/v4"
)

var User models.User

func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "login-app",
		"sub": "alex",
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	w.Write([]byte(jwtToken))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User: Alex"))
}
