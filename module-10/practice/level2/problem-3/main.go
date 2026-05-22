package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
)

var db *pgx.Conn;
var err error;
type user struct{
	Id int `json:"id"`;
	Name string `json:"name"`;
	Email string `json:"email"`;
	CGPA float64 `json:"cgpa"`;
}

func main() {
	func() {

		db, err = pgx.Connect(context.Background(), "postgresql://postgres:123456@localhost:5432/goDatabase")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Database is connecting here")

	}()

	mux :=http.NewServeMux();
    mux.HandleFunc("GET /users",searchQueryFunction)
	http.ListenAndServe(":5052",mux);

}

func searchQueryFunction(res http.ResponseWriter,req *http.Request){
 searchParams :=req.URL.Query();
 cgpa :=searchParams.Get("cgpa");

 query :=`SELECT * FROM users WHERE cgpa = $1`;
 
 rows,err :=db.Query(context.Background(),query,cgpa);
 if err != nil{
	res.WriteHeader(http.StatusInternalServerError);
	fmt.Fprintln(res,err);
	return;
 }
  var users []user;
 for rows.Next(){
  var findUser user;
  err =rows.Scan(&findUser.Id,&findUser.Name,&findUser.Email,&findUser.CGPA);
  if err != nil{
  	res.WriteHeader(http.StatusInternalServerError);
  	fmt.Fprintln(res,err);
  	return;
  }
  users = append(users,findUser);
 }

 res.Header().Set("Content-Type", "application/json");
 json.NewEncoder(res).Encode(users);
}