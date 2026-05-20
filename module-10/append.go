package main;

import (
	"fmt"
	
)

func main(){
	newArr :=[] int{1,2,3,4,5};
    //same slice needed
	newArr=append(newArr[:],newArr[:]...);
	
fmt.Println(newArr);
}