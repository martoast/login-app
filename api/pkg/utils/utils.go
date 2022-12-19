package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/token"
	"github.com/shaj13/go-guardian/store"
)

var authenticator auth.Authenticator

var cache store.Cache

func SetupGoGuardian() {
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)

	basicStrategy := basic.New(ValidateUser, cache)
	tokenStrategy := token.New(VerifyToken, cache)

	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(token.CachedStrategyKey, tokenStrategy)
}

func ValidateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	// here connect to db or any other service to fetch user and validate it.
	if userName == "alex" && password == "xx" {
		return auth.NewDefaultUser("alex", "1", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		user, err := authenticator.Authenticate(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticated\n", user.UserName())
		next.ServeHTTP(w, r)
	})
}

func VerifyToken(ctx context.Context, r *http.Request, tokenString string) (auth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := auth.NewDefaultUser(claims["sub"].(string), "", nil, nil)
		return user, nil
	}

	return nil, fmt.Errorf("invaled token")
}
