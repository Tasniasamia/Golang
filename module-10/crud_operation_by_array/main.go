package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
);

type userType struct{
	Id int `json:"id"`;
    Name string `json:"name"`;
	Cgpa float64 `json:"cgpa"`;
	Age int `json:"age"`;
}
	var users =[]userType{
		{
			Id:1,
			Name:"samia",
			Cgpa:3.88,
			Age:23,
		},{
			Id:2,
			Name:"ayla",
			Cgpa:3.22,
			Age:23,
		},
	};
func main(){
    userArray :=users[:];
	fmt.Println(userArray);



	route :=http.NewServeMux();
    
	route.HandleFunc("GET /",globalHandler);
    route.HandleFunc("POST /user",createUser);
	route.HandleFunc("PUT /user/{id}",updateUser);
	route.HandleFunc("DELETE /user/{id}",deleteUser);

	err :=http.ListenAndServe(":5500",route);
	if(err != nil){
		fmt.Println("Server error is ",err);
	}

}

func globalHandler(res http.ResponseWriter,req*http.Request){
res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusOK);
 json.NewEncoder(res).Encode(users);
}

func createUser(res http.ResponseWriter,req*http.Request){
var singleUser userType;
decoder := json.NewDecoder(req.Body)
decoder.DisallowUnknownFields();
err :=decoder.Decode(&singleUser);
 
 if(err !=nil){
	fmt.Fprintln(res,err);
	return;
 }

singleUser.Id=len(users)+1;
users =append(users, singleUser);
res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusCreated);
json.NewEncoder(res).Encode(users);

}

func updateUser(res http.ResponseWriter,req*http.Request){
 id :=req.PathValue("id");
 i,err :=strconv.Atoi(id);

 if(err != nil){
	fmt.Fprintln(res,"Invalid type Id");
	return;
 }


var singleUser userType;
decoder := json.NewDecoder(req.Body)
decoder.DisallowUnknownFields();
err=decoder.Decode(&singleUser);
 
 if(err !=nil){
	fmt.Fprintln(res,err);
	return;
 }


 for idx,values :=range users{
	if(values.Id == i){
	singleUser.Id=values.Id;
     users[idx]=singleUser;
	}
 }

 res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusCreated);
json.NewEncoder(res).Encode(users);

}

func deleteUser(res http.ResponseWriter,req*http.Request){
 id :=req.PathValue("id");
 i,err :=strconv.Atoi(id);

 if(err != nil){
	fmt.Fprintln(res,"Invalid type Id");
	return;
 }

 var index int;

 for idx,val :=range users{
	if(val.Id == i){
		index=idx;
		break;
	}
 }

users =append(users[:index],users[index+1:]...)
res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusCreated);
json.NewEncoder(res).Encode(users);

}

