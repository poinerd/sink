package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type formResponseType map[string]any

func insertFormResponseToDB(formResponse formResponseType, token string, db *sql.DB) error {
	payLoadBytes, err := json.Marshal(formResponse)
	if err != nil {
		return fmt.Errorf("there was an error converting input to json")
	}

	query := `INSERT INTO submissions (payload, form_hash) VALUES ($1 , $2);`

	// i need to put hash here too
	
	err = insertDataToDb(query, db, payLoadBytes, token)
	if err != nil {
		return fmt.Errorf("there was an error writing to the DB: %w", err)
	}
	return nil
}

func insertDataToDb(query string, db *sql.DB, data ...any) error {
	_, execErr := db.Exec(query, data...)
	if execErr != nil {
		return fmt.Errorf("postgres execution error: %v", execErr)
	}
	return nil
}