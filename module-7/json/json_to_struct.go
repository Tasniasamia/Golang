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
	
	// struct to json
	jsonData,err :=json.Marshal(user1);
     
	if(err != nil){
	fmt.Println("Error Found",err);
    } 
    convertString :=string(jsonData)
	fmt.Println("convert struct to json ",convertString);





	//json to struct

	var user userInfo;

    //err2 :=json.Unmarshal([]byte(convertString),&user);
	byteConvertion :=[]byte(convertString);
	err2 :=json.Unmarshal(byteConvertion,&user);

    if(err2 != nil){
	fmt.Println("Error Found here ",err2);
	}
    fmt.Println("convert json to struct ",user); // {Samia 23 4}
	fmt.Println("user.Name=",user.Name); // Samia
	fmt.Println("user.Age=",user.Age); // 23
	fmt.Println("user.Grade=",user.Grade); // 4

    //fmt.Println("user.name=",user.name); // not allowed
	//fmt.Println("user.age=",user.age); // not allowed
	//fmt.Println("user.grade=",user.grade); // not allowed
	


}