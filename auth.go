package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type acessFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `db:"id"` 
	Email    string `db:"email"` 
	Password string `db:"password"`
}

// FIXED: UserID tracking type updated to match UUID string structure
type CustomClaims struct {
	UserID string `json:"user_id"` 
	jwt.RegisteredClaims
}


func HandleSignup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var signUpCredentials acessFormat
		err := json.NewDecoder(r.Body).Decode(&signUpCredentials)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpCredentials.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error secure-hashing password", http.StatusInternalServerError)
			return
		}

		userEmail := signUpCredentials.Email
		UUID := uuid.New().String()

		query := `INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3);`
		
		err = insertDataToDb(query, db, UUID, userEmail, string(hashedPassword))
		if err != nil {
			fmt.Println("Database write error:", err)
			http.Error(w, "Error writing to the DB", http.StatusInternalServerError)
			return
		}

		fmt.Printf("[SUCCESS] Created user: %s\n", userEmail)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User successfully created",
			"id":      UUID,
		})
	}
}


func handleSignIn(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var signInCredentials acessFormat
		err := json.NewDecoder(r.Body).Decode(&signInCredentials)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		var existingUser User
		query := `SELECT id, password_hash from users WHERE email = $1;`
		
		// FIXED: Included pointers (&) to write the results back to your struct
		
        err = db.QueryRow(query, signInCredentials.Email).Scan(&existingUser.ID, &existingUser.Password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(signInCredentials.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)

		claims := &CustomClaims{
			UserID: existingUser.ID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        // you sign the jwt with your secret

		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			fmt.Println("JWT signature error:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	}
}