package main;
import (
	"net/http"
	"encoding/json"
	"goproject/config"
	"fmt"
	"goproject/rest/middleware"
)


type User struct{
	Id int `json:"id"`;
	Name string `json:"name"`;
	Email string `json:"email"`;
	Password string `json:"password"`;
}

var users []User;

func main(){

	mux:=http.NewServeMux();
   
	globalRoutes:=middleware.NewManager();

	globalRoutes.Use(middleware.Logger,middleware.Preflight,middleware.Cors)
	




	mux.HandleFunc("POST /resister",globalRoutes.With(createUser,middleware.Check));

	

	address:=":"+config.GetConfig().HTTP_PORT
    fmt.Println("Server is running",address);
	err:=http.ListenAndServe(address,mux);
	if(err != nil){
		panic(err);
	}

}


func createUser(w http.ResponseWriter,r *http.Request){

	var newUser User;
	err :=json.NewDecoder(r.Body).Decode(&newUser);
	if(err != nil){
		http.Error(w,err.Error(),http.StatusBadRequest);
	}
	newUser.Id=len(users)+1;
	users =append(users, newUser);
	w.Header().Add("Content-Type","application/json");

	w.WriteHeader(http.StatusCreated);

	 json.NewEncoder(w).Encode(users);
	
}

func init(){
  config.LoadConfig();
}