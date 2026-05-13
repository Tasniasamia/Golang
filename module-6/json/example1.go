package main;

import "fmt";

import "encoding/json";

type userInfo struct{
	Name string `json:"name"`;
	Age int `json:"age"`;
	Grade float64 `json:"grade"`;
}
func main(){

	user1 :=userInfo{
		Name:"Samia", 
		Age:23,
		Grade:4.00,
	}

	jsonData,err :=json.Marshal(user1);
	fmt.Println(string(jsonData));
	fmt.Println(err);
	// fmt.Println(string(byte(user1)));
}