package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"api/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	utils.SetupGoGuardian()
	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/token", utils.Middleware(http.HandlerFunc(createToken))).Methods("GET")
	router.HandleFunc("/v1/book/{id}", utils.Middleware(http.HandlerFunc(getBookAuthor))).Methods("GET")
	log.Printf("server started and listening on http://127.0.0.1:%s", port)
	http.ListenAndServe("127.0.0.1:"+port, router)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth-app",
		"sub": "alex",
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	w.Write([]byte(jwtToken))
}

func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	books := map[string]string{
		"1449311601": "Ryan Boyd",
		"148425094X": "Yvonne Wilson",
		"1484220498": "Prabath Siriwarden",
	}
	body := fmt.Sprintf("Author: %s \n", books[id])
	w.Write([]byte(body))
}
