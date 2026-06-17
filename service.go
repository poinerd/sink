package main

import (
	"encoding/json"
	"net/http"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"database/sql"
)

type createFormResponse struct{
	EndPointURL string `json:"endpoint_url"`
}


func submitForm(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers","Content-Type")

	/// From the link the user used to get here, we should be able to extract some details about the exact table to write the formDetails to

	var formDetails formResponseType
	token := r.URL.Query().Get("token")
	json.NewDecoder(r.Body).Decode(&formDetails)
	insertFormResponseToDB(formDetails, token, db)
	fmt.Println(formDetails)
}

}

// http.HandlerFunc ia the type of the controllers you write

func createFormEndpoint(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
   /// Create a unique hash to asoociate with this particular form
   /// attach the hash to the back of the base url
   /// sav the complete enpoint to the DB
   /// Return that API endpoint to the user
   
   bytes := make([]byte, 8)  // This is a slice of 8 bytes
   _ , err := rand.Read(bytes)

   if err != nil{
	fmt.Println("There was an error generating the random Number")
   }
  formHash := hex.EncodeToString(bytes)
  formEndpoint := fmt.Sprintf("%s%s", baseUrl, formHash )

  responseObject := createFormResponse{
	EndPointURL: formEndpoint,
  }

  
   query := `INSERT INTO forms (hash, user_id) VALUES ($1, $2);`
   userID, _ := r.Context().Value(UserIDKey).(string)
   err = insertDataToDb(query, db, formHash, userID)

   if err != nil{
      fmt.Println("error writing to the DB", err)
      return
   }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(responseObject)

}

}