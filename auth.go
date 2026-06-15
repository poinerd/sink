package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)


type acessFormat struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


// controllers/// r stands for request as in request from the server to the client

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
        _, execErr := db.Exec(query, UUID, userEmail, string(hashedPassword))
        
        if execErr != nil {
            fmt.Println("Database Execution Error:", execErr)
            http.Error(w, "Database write failure", http.StatusInternalServerError)
            return
        } 

        fmt.Printf("[SUCCESS] Created user: %s | Hash: %s\n", userEmail, string(hashedPassword))

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "User successfully created",
            "id":      UUID,
        })
    }
}