package main;

import "fmt";



func deferFunc()(int){
	a :=10;
	defer func(){
    fmt.Println("The value of a is ",a);
	}();

	fmt.Println("After defer code");
	a+=100;
	return a;
}







func main(){
fmt.Println(deferFunc());

}