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

	fmt.Println(user1.Name); // allowed
	fmt.Println(user1.Age); // allowed
	fmt.Println(user1.Grade); // allowed
	
	jsonData,err :=json.Marshal(user1);
     
	if(err != nil){
	fmt.Println("Error Found",err);
    } 
    convertString :=string(jsonData)
	fmt.Println(convertString);

	//fmt.Println(convertString.Name); //not allowed
	//fmt.Println(convertString.Age); //not allowed
	//fmt.Println(convertString.Grade); //not allowed

}