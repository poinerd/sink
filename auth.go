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

type User struct{
    ID int  `json:"id"`
    Email int `json:"email"`
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
        err = insertDataToDb(query, db, [...]any{UUID, userEmail, string(hashedPassword)})
        

        if err != nil{
            fmt.Println("Database write error", fmt.Errorf("There was error writing to the DB"))
            return
        }


        fmt.Printf("[SUCCESS] Created user: %s | Hash: %s\n", userEmail, string(hashedPassword))
        
        // When the server wants to send a request to the client, it is important to set the header first.
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "User successfully created",
            "id":      UUID,
        })
    }
}


func handleSignIn(db *sql.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){

        var signInCredentials acessFormat
        json.NewDecoder(r.Body).Decode(&signInCredentials)

        var existingUser User
        query := `SELECT id, password_hash from USERS WHERE email = $1;`
        err := db.QueryRow(query, signInCredentials.Email).Scan(existingUser.ID, existingUser.Password)

        if err != nil{
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(signInCredentials.Password))

        if err != nil{
            return
        }

        
    }
}