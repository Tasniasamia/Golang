package main;

import "fmt";

func returnFunction(a int)(func()){
	return func(){
		fmt.Println(a);
	}
}
func main(){


getFunction :=returnFunction(12);
getFunction();

}