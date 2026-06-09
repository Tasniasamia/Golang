package main;

import "fmt";
const aa=10;
func returnFunction(a int)(func()){
	return func(){
		fmt.Println(a);
	}
}
func main(){


getFunction :=returnFunction(12);
getFunction();

}