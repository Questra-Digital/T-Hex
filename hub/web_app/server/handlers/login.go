package handlers

import(
	"net/http"
	"encoding/json"
	"server/utils"
	"server/models"
)

func LoginHandler(w http.ResponseWriter,r *http.Request){
	
	//Only POST Methods are allowed
	if(r.Method != http.MethodPost){
		utils.RespondError(w,"Invalid Request Method",http.StatusMethodNotAllowed)
		return
	}

	var creds models.Credentials

	//Decode received json data in credentials struct
	err := json.NewDecoder(r.Body).Decode(&creds)

	// Validate the request body
	if err!=nil || creds.Username == "" || creds.Password == ""{
		utils.RespondError(w,"Invalid Request Body",http.StatusBadRequest)
		return
	}

	password,ok := models.ValidUsers[creds.Username]

	//Ensure the key with the given username exists
	if(!ok){
		utils.RespondError(w,"Invalid username",http.StatusUnauthorized)
		return
	}

	//Ensure that password is correct
	if(password != creds.Password){
		utils.RespondError(w,"Invalid password",http.StatusUnauthorized)
		return
	}

	//Generate JWT Token
	tokenString,err := utils.GenerateJWT(creds.Username)
	if err!=nil{
		utils.RespondError(w, "Error generating token",http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w,map[string]string{"token":tokenString})

}