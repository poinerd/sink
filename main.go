package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type formResponseType map[string]any // composite type

// controller

func submitForm(w http.ResponseWriter, r *http.Request){
	var formDetails formResponseType
	json.NewDecoder(r.Body).Decode(&formDetails)
	fmt.Println(formDetails)
}

func main(){
	http.HandleFunc("/submit", submitForm)
	fmt.Println("server is up and running on port 8080")
	http.ListenAndServe(":8000", nil)
}

