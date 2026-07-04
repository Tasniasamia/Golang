package product;
import (
	"mains/util"
	"mains/database"
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"

)


func GetProducts(w http.ResponseWriter, r *http.Request) {

	users:=database.List();
	util.SendResponse(w, users, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
   
	var newUser database.User;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	insertUser:=database.Store(newUser);
	util.SendResponse(w, insertUser, http.StatusCreated)
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
	getUserById := database.Find(productIdInt);
	if(getUserById.Id == productIdInt) {
		util.SendResponse(w, getUserById, http.StatusOK)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)	
}



func UpdateProduct(w http.ResponseWriter, r *http.Request) {

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

   var findUser database.User;
	err = json.NewDecoder(r.Body).Decode(&findUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

  fmt.Println("findUser:", findUser)
   findUser.Id=productIdInt;
   getUserById := database.Update(productIdInt, findUser);
	fmt.Println("getUserById:", getUserById)
	if(getUserById.Id == productIdInt) {
		util.SendResponse(w, getUserById, http.StatusOK)
		return
	}
	http.Error(w, "Product not updated", http.StatusNotFound)	
}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {

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
	database.Delete(productIdInt);
	http.Error(w, "Product deleted", http.StatusOK)	
}
