package main;

import (
	"fmt"
	"time"
)

func main(){

	timeStart :=time.Now();

	fmt.Println("Hello World");

	func(){
    fmt.Println("Hello I am anonymous function");
	}()

	fmt.Println("Total time is ",time.Since(timeStart).Seconds())
}