package main;

import (
	"fmt"
	"encoding/json"
	"context"
	"net/http"
	"github.com/jackc/pgx/v5"
	"strconv"
)

var db*pgx.Conn;
var err error;
type user struct{
	Id int `json:"id"`;
	Name string `json:"name"`;
	Age int `json:"age"`;
	Cgpa float64 `json:"cgpa"`;
}
func main(){
func(){
	db,err =pgx.Connect(context.Background(),"postgresql://postgres:123456@localhost:5432/goDatabase");
	if err != nil {
		panic(err);
	}
	fmt.Println("Database connected successfully");
}()

mux :=http.NewServeMux();
mux.HandleFunc("GET /users",getUsers);
http.ListenAndServe(":5054",mux);

}

func getUsers(res http.ResponseWriter, req *http.Request) {
	searchQuery :=req.URL.Query();
	page :=searchQuery.Get("page");
	limit :=searchQuery.Get("limit");
	pageInt, err := strconv.Atoi(page);
	if err != nil {
		res.WriteHeader(http.StatusBadRequest);
		fmt.Fprintln(res, "Invalid page parameter");
		return;
	}
	limitInt, err := strconv.Atoi(limit);
	if err != nil {
		res.WriteHeader(http.StatusBadRequest);
		fmt.Fprintln(res, "Invalid limit parameter");
		return;
	}

	skip :=(pageInt-1)*limitInt
    query :=`SELECT * FROM users LIMIT $1 OFFSET $2`
	var users [] user;
    rows,err :=db.Query(context.Background(),query,limitInt,skip);

	for rows.Next(){
		var user user;
		err=rows.Scan(&user.Id,&user.Name,&user.Cgpa,&user.Age);
		if err != nil {
            res.WriteHeader(http.StatusNotFound);
			fmt.Fprintln(res,err);
			return;
		}
		users =append(users, user);

	}


	res.Header().Set("Content-Type", "application/json");
	json.NewEncoder(res).Encode(users);

}