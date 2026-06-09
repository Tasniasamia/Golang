package main

import (
	"context"
	"crud_operation_by_array/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getAllUSer(res http.ResponseWriter, req *http.Request) {

	query := `SELECT * FROM users`
	rows, err := db.Db.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}

	var allUser []UserType
	for rows.Next() {
		var singleUser UserType
		err := rows.Scan(&singleUser.Id, &singleUser.Name, &singleUser.Cgpa, &singleUser.Age)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(res, err)
			return
		}
		allUser = append(allUser, singleUser)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allUser)
}

func createUser(res http.ResponseWriter, req *http.Request) {
	var singleUser UserType
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&singleUser)

	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	fmt.Println(singleUser)

	var newUser UserType

	query := `INSERT INTO users (name, age, cgpa) VALUES ($1, $2, $3) RETURNING *`
	err = db.Db.QueryRow(context.Background(), query, singleUser.Name, singleUser.Age, singleUser.Cgpa).Scan(&newUser.Id, &newUser.Name, &newUser.Cgpa, &newUser.Age)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}

	// singleUser.Id=len(users)+1;
	// users =append(users, singleUser);

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newUser)

}

func updateUser(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintln(res, "Invalid type Id")
		return
	}

	var singleUser UserType
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&singleUser)

	if err != nil {
		fmt.Println("wrong field")

		fmt.Fprintln(res, err)
		return
	}

	//findUser
	var singleUser2 UserType

	query := `SELECT*FROM users WHERE ID=$1`

	err = db.Db.QueryRow(context.Background(), query, i).Scan(&singleUser2.Id, &singleUser2.Name, &singleUser2.Cgpa, &singleUser2.Age)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println("find user wrong")
		fmt.Fprintln(res, err)
		return
	}

	//update User

	query2 := `UPDATE users SET name=$1, cgpa=$2, age=$3 WHERE id=$4 RETURNING *`
	err = db.Db.QueryRow(context.Background(), query2, singleUser.Name, singleUser.Cgpa, singleUser.Age, i).Scan(&singleUser.Id, &singleUser.Name, &singleUser.Cgpa, &singleUser.Age)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}

	fmt.Println("after update query ", err)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(singleUser)

}

func deleteUser(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintln(res, "Invalid type Id")
		return
	}

	var singleUser2 UserType

	query := `SELECT*FROM users WHERE ID=$1`

	err = db.Db.QueryRow(context.Background(), query, i).Scan(&singleUser2.Id, &singleUser2.Name, &singleUser2.Cgpa, &singleUser2.Age)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println("find user wrong")
		fmt.Fprintln(res, err)
		return
	}

	//  for idx,val :=range users{
	// 	if(val.Id == i){
	// 		index=idx;
	// 		break;
	// 	}
	//  }

	// users =append(users[:index],users[index+1:]...)

	query = `DELETE FROM users WHERE id=$1`
	_, err = db.Db.Exec(context.Background(), query, i)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "User Deleted Successfully")
	// json.NewEncoder(res).Encode(singleUser);

}
