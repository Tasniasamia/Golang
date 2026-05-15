package main;

import (
	"fmt"
	"log"
	
	
)

func doSomething(a int)int{
fmt.Println("Start work");
defer func(){
fmt.Println("defer here");
}()
log.Fatal("System crash here");
return a;
}

func main(){
fmt.Println(doSomething(0));
}