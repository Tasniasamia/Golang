package main;

import (
	"fmt"
	
)

func doSomething(a int)int{
fmt.Println("Start work");
defer func(){
r :=recover();
if r != nil{
	fmt.Println("Recover from ",r);
}
}()
if(a == 0){
	panic("0 is not allowed");
}
return a;
}

func main(){

	
	fmt.Println(doSomething(0));
}