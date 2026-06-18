package main

import (
	"fmt"
	"net/http"
	_ "github.com/jackc/pgx/v5/stdlib" 

)
func main(){
	db, err := loadDB()
	if err != nil{
		fmt.Println("There was an error loading the Database")
	}
	http.HandleFunc("/submit", submitForm(db))
	http.HandleFunc("/signup", HandleSignup(db))
    http.HandleFunc("/signin", handleSignIn(db) )
	http.HandleFunc("/create", authMiddleWare(createFormEndpoint(db)))
	
	fmt.Println("server is up and running on port 8000")
	http.ListenAndServe(":8000", nil)
}