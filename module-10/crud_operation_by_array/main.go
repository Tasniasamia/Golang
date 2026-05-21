package main

import (
	"encoding/json"
	"context"
	"fmt"
	"os"
	"net/http"
	"strconv"
	"github.com/jackc/pgx/v5"
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
var db *pgx.Conn;
var err error;
func connectDB(){
	 urlString := "postgres://postgres:123456@localhost:5432/goDatabase"
	db, err= pgx.Connect(context.Background(), urlString)
	if err != nil {
		panic(err);
		os.Exit(1)
	}
	fmt.Println("Server is connecting");
}
func main(){
	connectDB();

    userArray :=users[:];
	fmt.Println(userArray);



	route :=http.NewServeMux();
    
	route.HandleFunc("GET /",getAllUSer);
    route.HandleFunc("POST /user",createUser);
	route.HandleFunc("PUT /user/{id}",updateUser);
	route.HandleFunc("DELETE /user/{id}",deleteUser);

	err :=http.ListenAndServe(":5500",route);
	if(err != nil){
		fmt.Println("Server error is ",err);
	}

}

func getAllUSer(res http.ResponseWriter,req*http.Request){

query :=`SELECT * FROM users`;
rows,err :=db.Query(context.Background(),query);
defer rows.Close()
if err != nil{
	res.WriteHeader(http.StatusInternalServerError);
	fmt.Fprintln(res,err);
	return;
}

var allUser[] userType;
for rows.Next(){
	var singleUser userType;
	err:=rows.Scan(&singleUser.Id,&singleUser.Name,&singleUser.Age,&singleUser.Cgpa);
	if err != nil{
		res.WriteHeader(http.StatusInternalServerError);
        fmt.Fprintln(res,err);
		return;
	}
	allUser=append(allUser, singleUser);
}

res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusOK);
 json.NewEncoder(res).Encode(allUser);
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

var newUser userType;
query := `INSERT INTO users (name, cgpa, age) VALUES ($1, $2, $3) RETURNING *`;
err=db.QueryRow(context.Background(),query,singleUser.Name,singleUser.Age,singleUser.Cgpa).Scan(&newUser.Id,&newUser.Name,&newUser.Age,&newUser.Cgpa);

if err != nil{
	res.WriteHeader(http.StatusInternalServerError);
	fmt.Fprintln(res,err);
	return;
}

// singleUser.Id=len(users)+1;
// users =append(users, singleUser);


res.Header().Set("Content-Type","application/json");
res.WriteHeader(http.StatusCreated);
json.NewEncoder(res).Encode(newUser);

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

