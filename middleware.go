package main

import (
	"strings"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"context"
	"fmt"
)

type contextKey string
const UserIDKey contextKey = "UserID"

func authMiddleWare(next http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		authBody := r.Header.Get("Authorization")

		if authBody == "" || !strings.HasPrefix(authBody, "Bearer ") {
			http.Error(w, "Unauthorized: Missing or malformed token", http.StatusUnauthorized)
			return
		}
		
		authToken := strings.TrimPrefix(authBody, "Bearer ")

  	claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil // Using your global 'secret' byte slice
		})

		if err != nil || !token.Valid {
			fmt.Println("[JWT DEBUG ERROR]:", err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			
			return
		}

	ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
	next(w, r.WithContext(ctx))
	}

}