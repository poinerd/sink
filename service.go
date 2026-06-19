package main
import (
	"encoding/json"
	"net/http"
   "github.com/google/uuid"
	"fmt"
	"database/sql"
)

type createFormResponse struct{
	EndPointURL string `json:"endpoint_url"`
}

type newFormName struct{
   FormName string `json:"form_name"`
   TargetEmail string `json:"target_email"`
   
}

func submitForm(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){

      if r.Method != http.MethodPost{
         http.Error(w, "This method is not supoorted", http.StatusMethodNotAllowed)
         return
      }

	w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Headers","Content-Type")
   
	var formDetails formResponseType
	token := r.URL.Query().Get("token")

   var formExists bool
        checkQuery := `SELECT EXISTS(SELECT 1 FROM forms WHERE hash = $1);`
        err := db.QueryRow(checkQuery, token).Scan(&formExists)
        if err != nil || !formExists {
            http.Error(w, "Form endpoint not found or inactive", http.StatusNotFound)
            return
        }

	json.NewDecoder(r.Body).Decode(&formDetails)

   // Handle empty reponse body gracefully o
   // if the user cliks thier submit button and there is nothing in their field, it shoudl intiate a DB write even if the dev didnt set the input tags to be required
	insertFormResponseToDB(formDetails, token, db)
	fmt.Println(formDetails)
}

}

// http.HandlerFunc ia the type of the controllers you write

func createFormEndpoint(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
   if r.Method != http.MethodPost{
      http.Error(w, "This http method is not allowed", http.StatusMethodNotAllowed)
      return
   }

   var newForm newFormName
   json.NewDecoder(r.Body).Decode(&newForm)

  formHash := uuid.New().String()[:10]
  tail := "submit?token=" + formHash
  formEndpoint := fmt.Sprintf("%s%s", baseUrl, tail)

  responseObject := createFormResponse{
	EndPointURL: formEndpoint,
  }
   query := `INSERT INTO forms (hash, user_id, form_name, target_email) VALUES ($1, $2, $3, $4);`

   rawID := r.Context().Value(UserIDKey)
   userID, ok := rawID.(string)
   
   if !ok || userID == "" {

      fmt.Println("[DEBUG] Context extraction failed! Either key type mismatch or empty ID.")
      http.Error(w, "Unauthorized: Invalid user session", http.StatusUnauthorized)
      return
}
   targetEmail := newForm.TargetEmail
   newFormName := newForm.FormName

   err := insertDataToDb(query, db, formHash, userID, newFormName, targetEmail)
   if err != nil{
      http.Error(w, "Failed to save form endpoint", http.StatusInternalServerError)
      fmt.Println("error writing to the DB", err)
      return
   }
   
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(responseObject)

}

}


// The admin needs to be able to get the reponse to all of his forms
// An admin can have more than 1 form
// The stuff that connects all the differnet forms is the hash thats user 


// func getAllFormResponses(db *sql.DB) http.HandlerFunc{
//    return func(w http.ResponseWriter, r *http.Request){

//      query =`SELECT * from submissions where form_hash = $1`
     
//    }
// }







