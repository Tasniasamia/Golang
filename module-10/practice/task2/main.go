package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn
var err error

type user struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Age  int     `json:"age"`
	Cgpa float32 `json:"cgpa"`
}

func main() {

	func() {
		db, err = pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/goDatabase")
		if err != nil {
			panic(err)
		}
		fmt.Println("Server connected successfully")
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", createUser)
	http.ListenAndServe(":5051", mux)

}

func createUser(res http.ResponseWriter, req *http.Request) {
	var userdata user
	decode := json.NewDecoder(req.Body)
	decode.DisallowUnknownFields()
	err := decode.Decode(&userdata);
	fmt.Println(userdata);
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, "Invalid JSON data")
		return
	}

	if userdata.Age <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, "Invalid age")
		return
	}
	if userdata.Cgpa < 0 || userdata.Cgpa > 4 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, "Invalid CGPA")
		return
	}
	if userdata.Name == "" {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, "Invalid name")
		return
	}

	query := `INSERT INTO users (name, cgpa, age) VALUES($1,$2,$3) RETURNING *`

	err = db.QueryRow(context.Background(),query,userdata.Name,userdata.Cgpa,userdata.Age).Scan(&userdata.Id, &userdata.Name, &userdata.Cgpa, &userdata.Age)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}
    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(userdata)
}
