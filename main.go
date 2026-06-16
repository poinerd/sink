package main

import (
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
	http.HandleFunc("/createForm", createFormEndpoint(db))

	fmt.Println("server is up and running on port 8080")
	http.ListenAndServe(":8000", nil)
}