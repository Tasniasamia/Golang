package main

import (
    "crud_operation_by_array/db"
    "fmt"
	"net/http"
	"github.com/joho/godotenv"
);



func main(){
    err :=godotenv.Load();
	if err != nil {
		panic(err);
	}
	db.ConnectDB();

    



	route :=http.NewServeMux();
    
	route.HandleFunc("GET /",getAllUSer);
    route.HandleFunc("POST /user",createUser);
	route.HandleFunc("PUT /user/{id}",updateUser);
	route.HandleFunc("DELETE /user/{id}",deleteUser);

	err=http.ListenAndServe(":5500",route);
	if(err != nil){
		fmt.Println("Server error is ",err);
	}

}

