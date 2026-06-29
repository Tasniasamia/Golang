package handler;
import (
	"mains/util"
	"mains/database"
	"net/http"
	"encoding/json"
	"strconv"
)


func GetProducts(w http.ResponseWriter, r *http.Request) {
	
	util.SendResponse(w, database.Users, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
   
	var newUser database.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	newUser.Id = len(database.Users) + 1
	database.Users = append(database.Users, newUser)
	util.SendResponse(w, newUser, http.StatusCreated)
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	productId :=r.PathValue("id");
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if(productId == "") {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	for _,data :=range database.Users{
      if(data.Id == productIdInt) {
			util.SendResponse(w, data, http.StatusOK)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)	
}