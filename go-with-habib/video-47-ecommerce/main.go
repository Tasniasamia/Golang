package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

var users []user

func handleCor(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
	w.Header().Set("Content-Type","application/json");
    
}

func handlePrefilght(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func sendResponse(w http.ResponseWriter, data interface{},status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	 handleCor(w);
	handlePrefilght(w,r);
    
	if r.Method != http.MethodGet {
		http.Error(w, "Plz give me GET request", http.StatusBadRequest)
		return
	}
	sendResponse(w, users, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    handleCor(w);
	handlePrefilght(w,r);
	if r.Method != http.MethodPost {
		http.Error(w, "Plz give me POST request", http.StatusBadRequest)
		return
	}

	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	newUser.Id = len(users) + 1
	users = append(users, newUser)
	sendResponse(w, newUser, http.StatusCreated)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getProducts)
	mux.HandleFunc("/product", createProduct)
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}

func init() {
	u1 := user{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := user{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,
		Dept: "Mathematics",
	}
	u3 := user{
		Id:   3,
		Name: "Bob Johnson",
		Roll: 103,
		Age:  24,
		Dept: "Physics",
	}
	users = append(users, u1, u2, u3)
}