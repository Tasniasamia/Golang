package main

import (
	"fmt"
   "encoding/json"
   "net/http"
)

type userType struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Grade float64 `json:"grade"`
}
var users =[]userType{
		{Id: 1, Name: "Tasnai", Age: 24, Grade: 3.99},
		{Id: 1, Name: "Tasnai", Age: 24, Grade: 3.99},
		{Id: 1, Name: "Tasnai", Age: 24, Grade: 3.99},
	};

func main() {
	

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("POST /user", createUser)
	mux.HandleFunc("GET /users",getAllUser);
	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/user", rootHandler2)
	fmt.Println("Server is running at 5000")

	// err := http.ListenAndServe(":5000", nil)
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	// if(req.Method != "POST"){
	// 	res.WriteHeader(http.StatusMethodNotAllowed);
	// 	fmt.Fprintln(res,"Method not allowed");
	// 	return;
	// }
	fmt.Fprintln(res, "welcom to go sever")

}

func createUser(res http.ResponseWriter, req *http.Request) {
	var newUser userType;
	err :=json.NewDecoder(req.Body).Decode(&newUser);
	if err!=nil {
     fmt.Fprintln(res,err);
	 return;
	}
	newUser.Id=len(users)+1;
    users=append(users,newUser);

	data,_ :=json.Marshal(users)
	res.Header().Add("Content-Type","JSON")
	res.WriteHeader(http.StatusCreated);
	res.Write(data);
	
}

func getAllUser(res http.ResponseWriter,req*http.Request){
	//way 1
	// data,_ :=json.Marshal(users)
	// res.Header().Set("Content-Type","JSON")
	// res.Write(data);

	//way 2
	res.Header().Set("Content-Type","JSON")
	res.WriteHeader(http.StatusOK);
	json.NewEncoder(res).Encode(users);

}