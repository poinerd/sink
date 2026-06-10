package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib" 
	"os"
	"github.com/joho/godotenv"
)

type formResponseType map[string]any // composite type

var connectionString string


// controller
// the function takes in the response from the frontend and inserts it into the submissions table

func insertFormResponse(formResponse formResponseType, token string){
    // connectionString := PROCESS.
	db, err := sql.Open("pgx", connectionString )

	if err != nil{
		fmt.Println("There was an error openign the db")
	}
	payLoadBytes, err:= json.Marshal(formResponse)
    
	if err!= nil{
		fmt.Println("There was an error converting input to json")
	}

   // you have to order the rows the way its is in the DB
	query:= `INSERT INTO submissions (form_token, payload) VALUES ($1 , $2);`
	db.Exec(query, token, payLoadBytes)

}

func submitForm(w http.ResponseWriter, r *http.Request){
	/// From the link the user used to get here, we should be able to extract some details about the exact table to write the formDetails to
	var formDetails formResponseType
	token := r.URL.Query().Get("token")
	json.NewDecoder(r.Body).Decode(&formDetails)
	insertFormResponse(formDetails, token)
	fmt.Println(formDetails)
}

// func getAllFormResponses(w http.ResponseWriter, r *http.Request){
// 	query := `SELECT * FROM SUBMISSIONS`

// 	json.Marshal()
// }

func main(){
	err := godotenv.Load()
	if err != nil{
	fmt.Println("There was an error finding the env file")
	}
	connectionString = os.Getenv("CONNECTION_STRING")

	http.HandleFunc("/submit", submitForm)

	fmt.Println("server is up and running on port 8080")
	http.ListenAndServe(":8000", nil)
}

