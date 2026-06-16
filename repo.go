package main

import (
	"fmt"
	"database/sql"
	_"github.com/jackc/pgx/v5/stdlib" 
	"os"
	"github.com/joho/godotenv"
)

// composite type
var connectionString string
var baseUrl string


func loadDB()(*sql.DB, error){

	err := godotenv.Load()
	if err != nil{
		return nil, nil
	}

	connectionString = os.Getenv("CONNECTION_STRING")
	baseUrl = os.Getenv("BASE_URL")
	db, err := sql.Open("pgx", connectionString )


	if err!= nil{
		fmt.Println("There was error doing stuff")
	}

	return db, nil

}