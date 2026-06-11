package main

import (
	"encoding/json"
	"net/http"
)

type acessFormat struct {
	email string
	password string
}
// controllers/// r stands for request as in request from the server to the client

func signUp(w http.ResponseWriter, r *http.Request){
	var signUpCredentials acessFormat
    json.NewDecoder(r.Body).Decode(&signUpCredentials)
	json.NewEncoder(w).Encode(signUpCredentials)
}

func signIn(w http.ResponseWriter, r *http.Request){
	var signInCredentials acessFormat
	json.NewDecoder(r.Body).Decode(&signInCredentials)
	json.NewEncoder(w).Encode(signInCredentials)
}


