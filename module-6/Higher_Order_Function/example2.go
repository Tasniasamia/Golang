package main;

import "fmt";

func returnFunction(a int)(func(x int)){
	return func(y int){
		fmt.Println(a*y);
	}
}
func main(){


getFunction :=returnFunction(12);
getFunction(12);

}