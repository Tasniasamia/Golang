package main;

import (
	"fmt"
	
)

func doSomething(a int)int{
fmt.Println("Start work");
defer func(){
fmt.Println("defer here");
}()
if(a == 0){
	panic("0 is not allowed");
}
return a;
}

func main(){
fmt.Println(doSomething(0));
}