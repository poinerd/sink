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
var baseUrl string

// controller
// the function takes in the response from the frontend and inserts it into the submissions table

func insertFormResponse(formResponse formResponseType, token string, db *sql.DB){
    // connectionString := PROCESS.
	
	payLoadBytes, err:= json.Marshal(formResponse)

	if err!= nil{
		fmt.Println("There was an error converting input to json")
	}

   // you have to order the rows the way its is in the DB

	query:= `INSERT INTO submissions (form_token, payload) VALUES ($1 , $2);`
	_, execErr := db.Exec(query, token, payLoadBytes)

	if execErr != nil {
    fmt.Println("PostGres Error: ", execErr)
	} else {
    fmt.Println("Stuff written to database sucessfully")
}

}


func submitForm(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	/// From the link the user used to get here, we should be able to extract some details about the exact table to write the formDetails to
	
	var formDetails formResponseType
	token := r.URL.Query().Get("token")
	json.NewDecoder(r.Body).Decode(&formDetails)
	insertFormResponse(formDetails, token, db)
	fmt.Println(formDetails)
}

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
	baseUrl = os.Getenv("BASE_URL")

	db, err := sql.Open("pgx", connectionString )


	http.HandleFunc("/submit", submitForm(db))
	http.HandleFunc("/signup", HandleSignup(db))
	http.HandleFunc("/createForm", createFormEnpoint)

	fmt.Println("server is up and running on port 8080")
	http.ListenAndServe(":8000", nil)
}