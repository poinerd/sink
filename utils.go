package main

import(
	"encoding/json"
	"fmt"
	"database/sql"
)

func insertFormResponse(formResponse formResponseType, token string, db *sql.DB){
    // connectionString := PROCESS.
	
	payLoadBytes, err:= json.Marshal(formResponse)

	if err!= nil{
		fmt.Println("There was an error converting input to json")
	}

   // you have to order the rows the way it is in the DB

	query:= `INSERT INTO submissions (form_token, payload) VALUES ($1 , $2);`
	_, execErr := db.Exec(query, token, payLoadBytes)

	if execErr != nil {
    fmt.Println("PostGres Error: ", execErr)
	} else {
    fmt.Println("Stuff written to database sucessfully")
}

}