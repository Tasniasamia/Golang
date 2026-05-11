package main;

import "fmt";

func print(a interface{}){
	fmt.Println(a);
	_,ok :=a.(string)
	if ok {
		//fmt.Println("value is ",c);
		fmt.Println(" String type");
		return;
	}
	_,okk :=a.(int);
	if okk{
		//fmt.Println("value is ",b);
		fmt.Println("Integer Type");
		return;
	}

	//type assertion m
}

func main(){
//any interface type
var newVar interface{};
newVar=[] int {1,2,3};
fmt.Println(newVar);
newVar="hello world";
print(newVar);
fmt.Println(newVar);
newVar=12;
fmt.Println(newVar);

//any variable type
var var2 any;
var2="Hello";
print(var2);
fmt.Println(var2);
var2=1;
print(var2);
fmt.Println(var2);



}