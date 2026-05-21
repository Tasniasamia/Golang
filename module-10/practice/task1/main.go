package main;

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
    "github.com/jackc/pgx/v5"
	"context"

)

var db *pgx.Conn;
var err error;
type user struct{
	Id int `json:"id"`;
	Name string `json:"name"`;
	Age int `json:"age"`;
	Cgpa float64 `json:"cgpa"`;
	
}


func main(){

//connectDB
func(){
	db, err= pgx.Connect(context.Background(),"postgres://postgres:123456@localhost:5432/goDatabase")
	if err != nil {
		panic(err)
		
	}
	fmt.Println("Database Connected Successfully");
	}()


	mux :=http.NewServeMux();
    mux.HandleFunc("GET /user/{id}",getSingleUser);
	http.ListenAndServe(":5050",mux);
	
}

func getSingleUser(res http.ResponseWriter,req *http.Request){
 id :=req.PathValue("id");
 i,err :=strconv.Atoi(id);

 if err != nil {
 res.WriteHeader(http.StatusNotFound);
	 fmt.Fprintln(res, "error:", err);
	 return
 }
 
query :=`SELECT * FROM users WHERE id=$1`;
var singleUser user;
err=db.QueryRow(context.Background(),query,i).Scan(&singleUser.Id, &singleUser.Name, &singleUser.Cgpa, &singleUser.Age);
if err != nil {
 res.WriteHeader(http.StatusInternalServerError);
	 fmt.Fprintln(res, "error:", err);
	 return
 }

 res.Header().Set("Content-Type", "application/json");
 res.WriteHeader(http.StatusOK);
 json.NewEncoder(res).Encode(singleUser);

}





// go get gorm.io/gorm
// go get gorm.io/driver/postgres