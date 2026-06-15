package main

import (
	"encoding/json"
	"net/http"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// type newFormEnpoint struct{

// 	formDetails any

// } 

type createFormResponse struct{
	EndPointURL string `json:"endpoint_url"`
}

func createFormEnpoint(w http.ResponseWriter, r *http.Request) {
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

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(responseObject)

}

