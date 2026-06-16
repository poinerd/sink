package main

import(
	"encoding/json"
	"fmt"
	"database/sql"
)


type formResponseType map[string]any 


func insertFormResponseToDB(formResponse formResponseType, token string, db *sql.DB) error{

	payLoadBytes, err:= json.Marshal(formResponse)

	if err!= nil{
		return fmt.Errorf("There was an error converting input to json")
	}
	
	query:= `INSERT INTO submissions (form_token, payload) VALUES ($1 , $2);`
	data := [...]any{token, payLoadBytes}
    err = insertDataToDb(query, db, data)

	if err != nil{
		return fmt.Errorf("There was an error writing to the DB")
	}
	return  nil
	
}

func insertDataToDb(query string, db *sql.DB, data ...any) error{

	_, execErr := db.Exec(query, data...)

	if execErr != nil {
		return fmt.Errorf("Postgres Execution error")
	} 
    return nil
}

