package user

import (
	"mains/util"
	"encoding/json"
	"mains/database"
	"net/http"
)
func CreateUser(w http.ResponseWriter, r *http.Request){

	var newUser database.AuthUser;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	insertUser:=newUser.Resister();
	util.SendResponse(w, insertUser, http.StatusCreated)

}

func Login(w http.ResponseWriter, r *http.Request){

	

	var newUser database.AuthUser;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	insertUser:=database.Login(newUser.Email, newUser.Password);
	util.SendResponse(w, insertUser, http.StatusCreated)

}